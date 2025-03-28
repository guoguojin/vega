package snapshot

import (
	"context"
	"fmt"
	"io"
	"path"
	"sort"
	"time"

	"code.vegaprotocol.io/vega/logging"

	"go.uber.org/zap"

	"github.com/jackc/pgtype"
	"github.com/klauspost/compress/zip"

	"code.vegaprotocol.io/vega/datanode/networkhistory/segment"
	"code.vegaprotocol.io/vega/datanode/sqlstore"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type LoadResult struct {
	LoadedFromHeight int64
	LoadedToHeight   int64
	RowsLoaded       int64
}

type LoadLog interface {
	Infof(s string, args ...interface{})
	Info(msg string, fields ...zap.Field)
}

type loadableSegment interface {
	GetToHeight() int64
	GetDatabaseVersion() int64
	ZipFileName() string
	ZipFilePath() string
}

func (b *Service) RollbackToSegment(ctx context.Context, log LoadLog, rollbackToSegment loadableSegment) error {
	dbMeta, err := NewDatabaseMetaData(ctx, b.connPool)
	if err != nil {
		return fmt.Errorf("failed to get database meta data: %w", err)
	}

	if rollbackToSegment.GetDatabaseVersion() < dbMeta.DatabaseVersion {
		log.Infof("rolling back database to version %d from version %d", rollbackToSegment.GetDatabaseVersion(), dbMeta.DatabaseVersion)

		err = b.migrateSchemaDownToVersion(rollbackToSegment.GetDatabaseVersion())
		if err != nil {
			return fmt.Errorf("failed to migrate down database from version %d to %d: %w",
				dbMeta.DatabaseVersion, rollbackToSegment.GetDatabaseVersion(), err)
		}

		// Update the meta data after schema migration
		dbMeta, err = NewDatabaseMetaData(ctx, b.connPool)
		if err != nil {
			return fmt.Errorf("failed to get database meta data after migration: %w", err)
		}
	}

	tx, err := b.connPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	// Rollback on a committed transaction has no effect
	defer func() { _ = tx.Rollback(ctx) }()

	rollbackToBlock, err := sqlstore.GetAtHeightUsingConnection(ctx, tx, rollbackToSegment.GetToHeight())
	if err != nil {
		return fmt.Errorf("failed to get block at height: %w", err)
	}

	for _, meta := range dbMeta.TableNameToMetaData {
		if meta.Hypertable {
			result, err := tx.Exec(ctx, fmt.Sprintf("delete from %s where vega_time > $1", meta.Name), rollbackToBlock.VegaTime)
			if err != nil {
				return fmt.Errorf("failed to delete rows from %s: %w", meta.Name, err)
			}
			log.Infof("deleted %d rows from %s", result.RowsAffected(), meta.Name)
		}
	}

	segments := []loadableSegment{rollbackToSegment}
	rowsCopied, err := b.loadSegmentsWithTransaction(ctx, log, tx.Conn(), segments, true, true, dbMeta, map[string]time.Time{})
	if err != nil {
		return fmt.Errorf("failed to load current state: %w", err)
	}

	log.Infof("Restored current state from snapshot, %d rows loaded", rowsCopied)

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Infof("updating continuous aggregate data")
	err = updateContinuousAggregateDataFromHeight(ctx, b.connPool, rollbackToSegment.GetToHeight())
	if err != nil {
		return fmt.Errorf("failed to update continuous aggregate data: %w", err)
	}

	return nil
}

func (b *Service) LoadSnapshotData(ctx context.Context, log LoadLog, ch segment.ContiguousHistory[segment.Staged], connConfig sqlstore.ConnectionConfig, optimiseForAppend, verbose bool,
) (LoadResult, error) {
	if len(ch.Segments) == 0 {
		return LoadResult{}, fmt.Errorf("no segments to load")
	}

	datanodeBlockSpan, err := sqlstore.GetDatanodeBlockSpan(ctx, b.connPool)
	if err != nil {
		return LoadResult{}, fmt.Errorf("failed to check if datanode has data: %w", err)
	}

	if err = validateSpanOfHistoryToLoad(datanodeBlockSpan, ch.HeightFrom, ch.HeightTo); err != nil {
		return LoadResult{}, fmt.Errorf("failed to validate span of history to load: %w", err)
	}

	loadedFromHeight := int64(0)
	if datanodeBlockSpan.HasData {
		loadedFromHeight = datanodeBlockSpan.ToHeight + 1
	} else {
		err = sqlstore.RevertToSchemaVersionZero(b.log, connConfig, sqlstore.EmbedMigrations, verbose)
		if err != nil {
			return LoadResult{}, fmt.Errorf("failed to revert scheam to version zero: %w", err)
		}
		loadedFromHeight = ch.HeightFrom
	}

	_, err = b.connPool.Exec(ctx, "SET TIME ZONE 0")
	if err != nil {
		return LoadResult{}, fmt.Errorf("failed to set timezone to UTC: %w", err)
	}

	dbMetaData, err := NewDatabaseMetaData(ctx, b.connPool)
	if err != nil {
		return LoadResult{}, fmt.Errorf("failed to get database meta data: %w", err)
	}

	log.Info("copying data into database", logging.Int64("database version", dbMetaData.DatabaseVersion))

	var totalRowsCopied int64

	historyTableLastTimestampMap, err := b.getLastHistoryTimestampMap(ctx, dbMetaData)
	if err != nil {
		return LoadResult{}, fmt.Errorf("failed to get last timestamp for history tables: %w", err)
	}

	dbVersionToSegments := map[int64][]loadableSegment{}

	for _, segment := range ch.Segments {
		dbVersion := segment.GetDatabaseVersion()
		dbVersionToSegments[dbVersion] = append(dbVersionToSegments[dbVersion], segment)
	}

	dbVersionsAsc := make([]int64, 0, len(dbVersionToSegments))
	for k := range dbVersionToSegments {
		dbVersionsAsc = append(dbVersionsAsc, k)
	}
	sort.Slice(dbVersionsAsc, func(i, j int) bool {
		return dbVersionsAsc[i] < dbVersionsAsc[j]
	})

	for _, targetDatabaseVersion := range dbVersionsAsc {
		if dbMetaData.DatabaseVersion != targetDatabaseVersion {
			currentDatabaseVersion := dbMetaData.DatabaseVersion
			log.Info("migrating database", logging.Int64("current database version", currentDatabaseVersion), logging.Int64("target database version", targetDatabaseVersion))

			err := b.migrateSchemaUpToVersion(targetDatabaseVersion)
			if err != nil {
				return LoadResult{}, fmt.Errorf("failed to migrate schema to version %d: %w", targetDatabaseVersion, err)
			}

			// After migration update the database meta-data
			dbMetaData, err = NewDatabaseMetaData(ctx, b.connPool)
			if err != nil {
				return LoadResult{}, fmt.Errorf("failed to get database meta data after database migration: %w", err)
			}

			log.Infof("finished migrating database from version %d to version %d", currentDatabaseVersion, targetDatabaseVersion)
		}

		log.Infof("loading all segments with database version: %d", targetDatabaseVersion)
		tx, err := b.connPool.Begin(ctx)
		if err != nil {
			return LoadResult{}, fmt.Errorf("failed to begin transaction: %w", err)
		}
		// Rollback on a committed transaction has no effect
		defer func() { _ = tx.Rollback(ctx) }()

		rowsCopied, err := b.loadSegmentsWithTransaction(ctx, log, tx.Conn(), dbVersionToSegments[targetDatabaseVersion], optimiseForAppend,
			false, dbMetaData, historyTableLastTimestampMap)
		if err != nil {
			return LoadResult{}, fmt.Errorf("failed to load segments for database version %d: %w", targetDatabaseVersion, err)
		}
		err = tx.Commit(ctx)
		if err != nil {
			return LoadResult{}, fmt.Errorf("failed to commit transaction: %w", err)
		}

		totalRowsCopied += rowsCopied
	}

	log.Infof("recreating continuous aggregate data")
	err = UpdateContinuousAggregateDataFromHighWaterMark(ctx, b.connPool, ch.HeightTo)
	if err != nil {
		return LoadResult{}, fmt.Errorf("failed to recreate continuous aggregate data: %w", err)
	}

	return LoadResult{
		LoadedFromHeight: loadedFromHeight,
		LoadedToHeight:   ch.HeightTo,
		RowsLoaded:       totalRowsCopied,
	}, nil
}

func (b *Service) loadSegmentsWithTransaction(ctx context.Context, log LoadLog, conn *pgx.Conn, segments []loadableSegment,
	optimiseForAppend, currentStateOnly bool, dbMetaData DatabaseMetadata,
	historyTableLastTimestampMap map[string]time.Time,
) (int64, error) {
	err := executeAllSql(ctx, conn, log, dbMetaData.CurrentStateTablesDropConstraintsSql)
	if err != nil {
		return 0, fmt.Errorf("failed to executed current state table drop constraints sql: %w", err)
	}

	var droppedIndexes []IndexInfo
	if !optimiseForAppend {
		log.Infof("dropping history table constraints")
		err = executeAllSql(ctx, conn, log, dbMetaData.HistoryStateTablesDropConstraintsSql)
		if err != nil {
			return 0, fmt.Errorf("failed to executed history table drop constraints sql: %w", err)
		}

		log.Infof("dropping history table indexes")
		droppedIndexes, err = dropHistoryTableIndexes(ctx, conn, log, dbMetaData)
		if err != nil {
			return 0, fmt.Errorf("failed to drop history table indexes: %w", err)
		}
	}

	var historyRowsCopied int64
	if !currentStateOnly {
		historyRowsCopied, err = b.loadHistorySegments(ctx, log, conn, segments, dbMetaData, historyTableLastTimestampMap)
		if err != nil {
			return 0, fmt.Errorf("failed to load history segments: %w", err)
		}
	}

	currentStateRowsCopied, err := b.loadCurrentState(ctx, log, conn, segments[len(segments)-1], dbMetaData,
		historyTableLastTimestampMap)
	if err != nil {
		return 0, fmt.Errorf("failed to load current state: %w", err)
	}

	if !optimiseForAppend {
		log.Infof("restoring history table indexes")
		err = createIndexes(ctx, conn, log, droppedIndexes)
		if err != nil {
			return 0, fmt.Errorf("failed to create indexes: %w", err)
		}

		log.Infof("restoring history table constraints")
		err = executeAllSql(ctx, conn, log, dbMetaData.HistoryStateTablesCreateConstraintsSql)
		if err != nil {
			return 0, fmt.Errorf("failed to executed history table create constraints sql: %w", err)
		}
	}

	log.Infof("restoring current state table constraints")
	err = executeAllSql(ctx, conn, log, dbMetaData.CurrentStateTablesCreateConstraintsSql)
	if err != nil {
		return 0, fmt.Errorf("failed to executed current state table create constraints sql: %w", err)
	}

	return historyRowsCopied + currentStateRowsCopied, nil
}

func (b *Service) loadCurrentState(ctx context.Context, log LoadLog, conn *pgx.Conn, segment loadableSegment,
	dbMetaData DatabaseMetadata, historyTableLastTimestampMap map[string]time.Time,
) (int64, error) {
	rowsCopied, err := b.loadSegment(ctx, conn, log, segment, "currentstate/", dbMetaData, historyTableLastTimestampMap)
	if err != nil {
		return 0, fmt.Errorf("failed to load current state snapshot %s: %w", segment, err)
	}

	return rowsCopied, nil
}

func (b *Service) loadHistorySegments(ctx context.Context, log LoadLog, conn *pgx.Conn, segments []loadableSegment,
	dbMetaData DatabaseMetadata, historyTableLastTimestampMap map[string]time.Time,
) (int64, error) {
	var totalRowsCopied int64
	for _, segment := range segments {
		rowsCopied, err := b.loadSegment(ctx, conn, log, segment, "history/", dbMetaData, historyTableLastTimestampMap)
		if err != nil {
			return 0, fmt.Errorf("failed to load history segment %s: %w", segment, err)
		}
		totalRowsCopied += rowsCopied
	}

	return totalRowsCopied, nil
}

func validateSpanOfHistoryToLoad(existingDatanodeSpan sqlstore.DatanodeBlockSpan, historyFromHeight int64, historyToHeight int64) error {
	if !existingDatanodeSpan.HasData {
		return nil
	}

	if historyFromHeight < existingDatanodeSpan.FromHeight {
		return fmt.Errorf("loading history from height %d is not possible as it is before the datanodes oldest block height %d, to load this history first empty the datanode",
			historyFromHeight, existingDatanodeSpan.FromHeight)
	}

	if historyFromHeight > existingDatanodeSpan.ToHeight+1 {
		return fmt.Errorf("the from height of the history to load, %d, must fall within or be one greater than the datanodes current span of %d to %d", historyFromHeight,
			existingDatanodeSpan.FromHeight, existingDatanodeSpan.ToHeight)
	}

	if historyFromHeight >= existingDatanodeSpan.FromHeight && historyToHeight <= existingDatanodeSpan.ToHeight {
		return fmt.Errorf("the span of history requested to load, %d to %d, is within the datanodes current span of %d to %d", historyFromHeight,
			historyToHeight, existingDatanodeSpan.FromHeight, existingDatanodeSpan.ToHeight)
	}

	return nil
}

func (b *Service) loadSegment(ctx context.Context, vegaDbConn *pgx.Conn, loadLog LoadLog, segment loadableSegment, zipDir string,
	dbMetaData DatabaseMetadata, historyTableLastTimestamps map[string]time.Time,
) (int64, error) {
	reader, err := zip.OpenReader(segment.ZipFilePath())
	if err != nil {
		return 0, fmt.Errorf("failed to open zip reader: %w", err)
	}

	loadLog.Infof("copying %s into database", segment.ZipFileName())
	startTime := time.Now()
	rowsCopied, err := copyDataIntoDatabase(ctx, vegaDbConn, reader, zipDir, dbMetaData, historyTableLastTimestamps)
	if err != nil {
		return 0, fmt.Errorf("failed to copy data into the database %s : %w", segment.ZipFileName(), err)
	}
	elapsed := time.Since(startTime)
	loadLog.Infof("copied %d rows from %s into database in %s", rowsCopied, segment.ZipFileName(), elapsed.String())

	return rowsCopied, nil
}

func getLastPartitionColumnEntryForHistoryTable(ctx context.Context, vegaDbConn *pgxpool.Conn, historyTableMetaData TableMetadata) (time.Time, error) {
	timeSelect := fmt.Sprintf(`SELECT %s FROM %s order by %s desc limit 1`, historyTableMetaData.PartitionColumn, historyTableMetaData.Name,
		historyTableMetaData.PartitionColumn)

	rows, err := vegaDbConn.Query(ctx, timeSelect)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to query last partition column time: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return time.Time{}, fmt.Errorf("failed to get values for row: %w", err)
		}
		if len(values) != 1 {
			return time.Time{}, fmt.Errorf("expected just 1 value got %d", len(values))
		}
		partitionTime, ok := values[0].(time.Time)
		if !ok {
			return time.Time{}, fmt.Errorf("expected value to be of type time, got %v", values[0])
		}

		return partitionTime, nil
	}

	return time.Time{}, nil
}

func (b *Service) getLastHistoryTimestampMap(ctx context.Context, dbMetadata DatabaseMetadata) (map[string]time.Time, error) {
	lastHistoryTimestampMap := map[string]time.Time{}

	conn, err := b.connPool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	for table, metadata := range dbMetadata.TableNameToMetaData {
		if metadata.PartitionColumn != "" {
			lastPartitionTime, err := getLastPartitionColumnEntryForHistoryTable(ctx, conn, metadata)
			if err != nil {
				return nil, fmt.Errorf("failed to get last partition column entry for history table %s: %w", table, err)
			}
			lastHistoryTimestampMap[table] = lastPartitionTime
		}
	}

	return lastHistoryTimestampMap, nil
}

func executeAllSql(ctx context.Context, vegaDbConn sqlstore.Connection, loadLog LoadLog, allSql []string) error {
	for _, sql := range allSql {
		loadLog.Infof("executing sql: %s", sql)
		_, err := vegaDbConn.Exec(ctx, sql)
		if err != nil {
			return fmt.Errorf("failed to execute sql %s: %w", sql, err)
		}
	}
	return nil
}

func dropHistoryTableIndexes(ctx context.Context, vegaDbConn sqlstore.Connection, loadLog LoadLog,
	dbMetadata DatabaseMetadata,
) ([]IndexInfo, error) {
	var indexes []IndexInfo

	rows, err := vegaDbConn.Query(ctx, `select tablename, Indexname, Indexdef from pg_indexes where schemaname ='public' order by tablename`)
	if err != nil {
		return nil, fmt.Errorf("failed to get table indexes: %w", err)
	}

	var allIndexes []IndexInfo
	if err = pgxscan.ScanAll(&allIndexes, rows); err != nil {
		return nil, fmt.Errorf("scanning table indexes: %w", err)
	}

	for _, index := range allIndexes {
		if dbMetadata.TableNameToMetaData[index.Tablename].Hypertable {
			indexes = append(indexes, index)
		}
	}

	loadLog.Infof("dropping history table indexes")
	for _, index := range indexes {
		_, err = vegaDbConn.Exec(ctx, fmt.Sprintf("DROP INDEX %s", index.Indexname))

		if err != nil {
			return nil, fmt.Errorf("failed to drop index %s: %w", index.Indexname, err)
		}
	}

	return indexes, nil
}

func copyDataIntoDatabase(ctx context.Context, conn *pgx.Conn, zipReader *zip.ReadCloser, zipDir string, dbMetaData DatabaseMetadata,
	historyTableLastTimestamps map[string]time.Time,
) (rowsCopied int64, err error) {
	// Disable all triggers
	_, err = conn.Exec(ctx, "SET session_replication_role = replica;")
	if err != nil {
		return 0, fmt.Errorf("failed to disable triggers, setting session replication role to replica failed: %w", err)
	}
	defer func() {
		_, triggersErr := conn.Exec(ctx, "SET session_replication_role = DEFAULT;")
		if err == nil && triggersErr != nil {
			err = fmt.Errorf("failed to re-enable triggers, setting session replication role to default failed: %w", err)
			rowsCopied = 0
		}
	}()

	var totalRowsCopied int64

	for _, file := range zipReader.File {
		dir, tableName := path.Split(file.Name)
		if dir != zipDir {
			continue
		}

		tableReader, err := file.Open()
		if err != nil {
			return 0, fmt.Errorf("failed to open file inside segment archive %w", err)
		}

		rowsCopied, err := copyTableDataIntoDatabase(ctx, dbMetaData.TableNameToMetaData[tableName], tableReader, conn, historyTableLastTimestamps)
		if err != nil {
			return 0, fmt.Errorf("failed to copy data into table %s: %w", tableName, err)
		}

		totalRowsCopied += rowsCopied
	}

	return totalRowsCopied, nil
}

func copyTableDataIntoDatabase(ctx context.Context, tableMetaData TableMetadata, reader io.ReadCloser,
	conn *pgx.Conn, historyTableLastTimestamps map[string]time.Time,
) (int64, error) {
	var err error
	var rowsCopied int64

	if tableMetaData.Hypertable {
		rowsCopied, err = copyHistoryTableDataIntoDatabase(ctx, tableMetaData, reader, conn, historyTableLastTimestamps)
		if err != nil {
			return 0, fmt.Errorf("failed to copy history table data into database: %w", err)
		}
	} else {
		tableTruncateSQL := fmt.Sprintf("truncate table %s", tableMetaData.Name)
		_, err := conn.Exec(ctx, tableTruncateSQL)
		if err != nil {
			return 0, fmt.Errorf("failed to truncate table %s: %w", tableMetaData.Name, err)
		}

		rowsCopied, err = copyCurrentStateTableDataIntoDatabase(ctx, tableMetaData, conn, reader)
		if err != nil {
			return 0, fmt.Errorf("failed to copy current state table data into database: %w", err)
		}
	}
	return rowsCopied, nil
}

func copyCurrentStateTableDataIntoDatabase(ctx context.Context, tableMetaData TableMetadata, conn *pgx.Conn, reader io.Reader) (int64, error) {
	query := fmt.Sprintf(`copy %s from STDIN (FORMAT csv, HEADER)`, tableMetaData.Name)
	tag, err := conn.PgConn().CopyFrom(ctx, reader, query)
	if err != nil {
		return 0, fmt.Errorf("failed to copy data into current state table: %w", err)
	}

	return tag.RowsAffected(), nil
}

func copyHistoryTableDataIntoDatabase(ctx context.Context, tableMetaData TableMetadata, reader io.Reader, conn *pgx.Conn,
	historyTableLastTimestamps map[string]time.Time,
) (int64, error) {
	partitionColumn := tableMetaData.PartitionColumn
	timestampString, err := encodeTimestampToString(historyTableLastTimestamps[tableMetaData.Name])
	if err != nil {
		return 0, fmt.Errorf("failed to encode timestamp into string: %w", err)
	}

	copyQuery := fmt.Sprintf(`copy %s from STDIN (FORMAT csv, HEADER) where %s > timestamp '%s'`, tableMetaData.Name,
		partitionColumn, timestampString)

	tag, err := conn.PgConn().CopyFrom(ctx, reader, copyQuery)
	if err != nil {
		return 0, fmt.Errorf("failed to copy data into hyper-table %s: %w", tableMetaData.Name, err)
	}

	return tag.RowsAffected(), nil
}

// encodeTimestampToString is required as pgx does not support parameter interpolation on copy statements.
func encodeTimestampToString(lastPartitionColumnEntry time.Time) ([]byte, error) {
	lastPartitionColumnEntry = lastPartitionColumnEntry.UTC()

	ts := pgtype.Timestamp{
		Time:   lastPartitionColumnEntry,
		Status: pgtype.Present,
	}

	var err error
	var timeText []byte
	timeText, err = ts.EncodeText(nil, timeText)
	if err != nil {
		return nil, fmt.Errorf("failed to encode timestamp: %w", err)
	}
	return timeText, nil
}

func createIndexes(ctx context.Context, vegaDbConn sqlstore.Connection,
	loadLog LoadLog, indexes []IndexInfo,
) error {
	for _, index := range indexes {
		loadLog.Infof("creating index %s", index.Indexname)
		_, err := vegaDbConn.Exec(ctx, index.Indexdef)
		if err != nil {
			return fmt.Errorf("failed to create index %s: %w", index.Indexname, err)
		}
	}
	return nil
}

func updateContinuousAggregateDataFromHeight(ctx context.Context, conn *pgxpool.Pool, height int64) error {
	fromBlock, err := sqlstore.GetAtHeightUsingConnection(ctx, conn, height)
	if err != nil {
		return fmt.Errorf("failed to get to block: %w", err)
	}

	dbMetaData, err := NewDatabaseMetaData(ctx, conn)
	if err != nil {
		return fmt.Errorf("failed to get database meta data: %w", err)
	}

	for _, cagg := range dbMetaData.ContinuousAggregatesMetaData {
		var err error
		highWatermark, err := getHighwaterMarkForCagg(ctx, conn, cagg)
		if err != nil {
			return fmt.Errorf("failed to get high watermark for cagg %s: %w", cagg.Name, err)
		}

		if highWatermark.Before(time.UnixMilli(0)) {
			// No cagg has been calculated yet, skip
			continue
		}

		toString, err := toStoredProcTimestampArg(highWatermark)
		if err != nil {
			return fmt.Errorf("failed to convert from timestamp %s to postgres string: %w", highWatermark, err)
		}

		// Truncate the from Time down to nearest complete cagg boundary
		fromTime := fromBlock.VegaTime.Truncate(cagg.BucketInterval)

		// When calling `refresh_continuous_aggregate` the refresh interval needs to be at least 2 times the bucket interval
		if fromTime.Sub(highWatermark) < 2*cagg.BucketInterval {
			fromTime = highWatermark.Add(-2 * cagg.BucketInterval)
		}
		fromString, err := toStoredProcTimestampArg(fromTime)
		if err != nil {
			return fmt.Errorf("failed to convert from timestamp %s to postgres string: %w", fromTime, err)
		}

		_, err = conn.Exec(ctx, fmt.Sprintf("CALL refresh_continuous_aggregate('%s', %s, %s);;", cagg.Name, fromString, toString))
		if err != nil {
			return fmt.Errorf("failed to refresh continuous aggregate %s: %w", cagg.Name, err)
		}
	}

	return nil
}

func getHighwaterMarkForCagg(ctx context.Context, conn *pgxpool.Pool, cagg ContinuousAggregateMetaData) (time.Time, error) {
	query := fmt.Sprintf(`SELECT COALESCE(
    _timescaledb_internal.to_timestamp(_timescaledb_internal.cagg_watermark(%d)),
    '-infinity'::timestamp with time zone);`, cagg.ID)
	row := conn.QueryRow(ctx, query)

	var highWatermark time.Time
	err := row.Scan(&highWatermark)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get high water mark: %w", err)
	}
	return highWatermark, nil
}

func UpdateContinuousAggregateDataFromHighWaterMark(ctx context.Context, conn *pgxpool.Pool, toHeight int64) error {
	toBlock, err := sqlstore.GetAtHeightUsingConnection(ctx, conn, toHeight)
	if err != nil {
		return fmt.Errorf("failed to get to block: %w", err)
	}

	dbMetaData, err := NewDatabaseMetaData(ctx, conn)
	if err != nil {
		return fmt.Errorf("failed to get database meta data: %w", err)
	}

	for _, cagg := range dbMetaData.ContinuousAggregatesMetaData {
		var err error
		highWatermark, err := getHighwaterMarkForCagg(ctx, conn, cagg)
		if err != nil {
			return fmt.Errorf("failed to get high watermark for cagg %s: %w", cagg.Name, err)
		}

		// Truncate the toTime down to latest complete cagg boundary
		toTime := toBlock.VegaTime.Truncate(cagg.BucketInterval)

		// When calling `refresh_continuous_aggregate` the refresh interval needs to be at least 2 times the bucket interval
		if toTime.Sub(highWatermark) < 2*cagg.BucketInterval {
			continue
		}

		fromString := "NULL"
		if highWatermark.After(time.UnixMilli(0)) {
			fromString, err = toStoredProcTimestampArg(highWatermark)
			if err != nil {
				return fmt.Errorf("failed to convert from timestamp %s to postgres string: %w", highWatermark, err)
			}
		}

		toString, err := toStoredProcTimestampArg(toTime)
		if err != nil {
			return fmt.Errorf("failed to convert to timestamp %s to postgres string: %w", toTime, err)
		}

		_, err = conn.Exec(ctx, fmt.Sprintf("CALL refresh_continuous_aggregate('%s', %s, %s);;", cagg.Name, fromString, toString))
		if err != nil {
			return fmt.Errorf("failed to refresh continuous aggregate %s: %w", cagg.Name, err)
		}
	}

	return nil
}

func toStoredProcTimestampArg(from time.Time) (string, error) {
	t := pgtype.Timestamp{}
	t.Set(from)
	fromBytes, err := t.EncodeText(nil, []byte{})
	if err != nil {
		if err != nil {
			return "", fmt.Errorf("failed to encode time: %w", err)
		}
	}

	fromString := string(fromBytes)
	return "'" + fromString + "'", nil
}
