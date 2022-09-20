// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.DATANODE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package sqlstore

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"

	"code.vegaprotocol.io/vega/datanode/entities"
	"code.vegaprotocol.io/vega/datanode/metrics"
	v2 "code.vegaprotocol.io/vega/protos/data-node/api/v2"
	"code.vegaprotocol.io/vega/protos/vega"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

var accountOrdering = TableOrdering{
	ColumnOrdering{Name: "account_id", Sorting: ASC},
}

type Accounts struct {
	*ConnectionSource
	cache     map[accountCacheKey]entities.Account
	cacheLock sync.Mutex
}

func NewAccounts(connectionSource *ConnectionSource) *Accounts {
	a := &Accounts{
		ConnectionSource: connectionSource,
		cache:            make(map[accountCacheKey]entities.Account),
	}
	return a
}

type accountCacheKey struct {
	PartyID  entities.PartyID
	AssetID  entities.AssetID
	MarketID entities.MarketID
	Type     vega.AccountType
}

func (as *Accounts) checkCache(
	partyID entities.PartyID,
	assetID entities.AssetID,
	marketID entities.MarketID,
	accType vega.AccountType,
) (entities.Account, bool) {
	key := accountCacheKey{partyID, assetID, marketID, accType}
	account, ok := as.cache[key]
	return account, ok
}

func (as *Accounts) updateCache(a entities.Account) {
	key := accountCacheKey{a.PartyID, a.AssetID, a.MarketID, a.Type}
	as.cache[key] = a
}

// Add inserts a row and updates supplied struct with autogenerated ID.
func (as *Accounts) Add(ctx context.Context, a *entities.Account) error {
	defer metrics.StartSQLQuery("Accounts", "Add")()

	err := as.Connection.QueryRow(ctx,
		`INSERT INTO accounts(id, party_id, asset_id, market_id, type, tx_hash, vega_time)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)
		 RETURNING id`,
		deterministicAccountID(a),
		a.PartyID,
		a.AssetID,
		a.MarketID,
		a.Type,
		a.TxHash,
		a.VegaTime).Scan(&a.ID)
	return err
}

func (as *Accounts) GetByID(id entities.AccountID) (entities.Account, error) {
	a := entities.Account{}
	ctx := context.Background()
	defer metrics.StartSQLQuery("Accounts", "GetByID")()
	err := pgxscan.Get(ctx, as.Connection, &a,
		`SELECT id, party_id, asset_id, market_id, type, tx_hash, vega_time
		 FROM accounts WHERE id=$1`,
		id)
	return a, err
}

func (as *Accounts) GetAll() ([]entities.Account, error) {
	ctx := context.Background()
	accounts := []entities.Account{}
	defer metrics.StartSQLQuery("Accounts", "GetAll")()
	err := pgxscan.Select(ctx, as.Connection, &accounts, `
		SELECT id, party_id, asset_id, market_id, type, tx_hash, vega_time
		FROM accounts`)
	return accounts, err
}

// Obtain will either fetch or create an account in the database.
// If an account with matching party/asset/market/type does not exist in the database, create one.
// If an account already exists, fetch that one.
// In either case, update the entities.Account object passed with an ID from the database.
func (as *Accounts) Obtain(ctx context.Context, a *entities.Account) error {
	as.cacheLock.Lock()
	defer as.cacheLock.Unlock()
	if account, ok := as.checkCache(a.PartyID, a.AssetID, a.MarketID, a.Type); ok {
		a.ID = account.ID
		return nil
	}
	insertQuery := `INSERT INTO accounts(id, party_id, asset_id, market_id, type, tx_hash, vega_time)
                           VALUES ($1, $2, $3, $4, $5, $6, $7)
                           ON CONFLICT (party_id, asset_id, market_id, type) DO NOTHING`

	selectQuery := `SELECT id, party_id, asset_id, market_id, type, tx_hash, vega_time
	                FROM accounts
	                WHERE party_id=$1 AND asset_id=$2 AND market_id=$3 AND type=$4`

	batch := pgx.Batch{}

	accountID := deterministicAccountID(a)
	batch.Queue(insertQuery, accountID, a.PartyID, a.AssetID, a.MarketID, a.Type, a.TxHash, a.VegaTime)
	batch.Queue(selectQuery, a.PartyID, a.AssetID, a.MarketID, a.Type)
	defer metrics.StartSQLQuery("Accounts", "Obtain")()
	results := as.Connection.SendBatch(ctx, &batch)
	defer results.Close()

	if _, err := results.Exec(); err != nil {
		return fmt.Errorf("inserting account: %w", err)
	}

	rows, err := results.Query()
	if err != nil {
		return fmt.Errorf("querying accounts: %w", err)
	}

	if err = pgxscan.ScanOne(a, rows); err != nil {
		return fmt.Errorf("scanning account: %w", err)
	}

	as.updateCache(*a)
	return nil
}

func deterministicAccountID(a *entities.Account) string {
	idAsBytes := sha256.Sum256([]byte(a.AssetID.String() + a.PartyID.String() + a.MarketID.String() + a.Type.String()))
	accountID := hex.EncodeToString(idAsBytes[:])
	return accountID
}

func (as *Accounts) Query(filter entities.AccountFilter) ([]entities.Account, error) {
	query, args, err := filterAccountsQuery(filter)
	if err != nil {
		return nil, err
	}
	accs := []entities.Account{}

	defer metrics.StartSQLQuery("Accounts", "Query")()
	rows, err := as.Connection.Query(context.Background(), query, args...)
	if err != nil {
		return accs, fmt.Errorf("querying accounts: %w", err)
	}
	defer rows.Close()

	if err = pgxscan.ScanAll(&accs, rows); err != nil {
		return accs, fmt.Errorf("scanning account: %w", err)
	}

	return accs, nil
}

func (as *Accounts) QueryBalancesV1(ctx context.Context, filter entities.AccountFilter, pagination entities.OffsetPagination) ([]entities.AccountBalance, error) {
	query, args, err := filterAccountBalancesQuery(filter)
	if err != nil {
		return nil, fmt.Errorf("querying account balances: %w", err)
	}

	query, args = orderAndPaginateQuery(query, nil, pagination, args...)

	accountBalances := make([]entities.AccountBalance, 0)

	defer metrics.StartSQLQuery("Accounts", "QueryBalancesV1")()
	rows, err := as.Connection.Query(ctx, query, args...)
	if err != nil {
		return accountBalances, fmt.Errorf("querying account balances: %w", err)
	}
	defer rows.Close()

	if err = pgxscan.ScanAll(&accountBalances, rows); err != nil {
		return accountBalances, fmt.Errorf("parsing account balances: %w", err)
	}

	return accountBalances, nil
}

func (as *Accounts) QueryBalances(ctx context.Context,
	filter entities.AccountFilter,
	pagination entities.CursorPagination,
) ([]entities.AccountBalance, entities.PageInfo, error) {
	query, args, err := filterAccountBalancesQuery(filter)
	if err != nil {
		return nil, entities.PageInfo{}, fmt.Errorf("querying account balances: %w", err)
	}

	query, args, err = PaginateQuery[entities.AccountCursor](query, args, accountOrdering, pagination)
	if err != nil {
		return nil, entities.PageInfo{}, fmt.Errorf("querying account balances: %w", err)
	}

	defer metrics.StartSQLQuery("Accounts", "QueryBalances")()

	accountBalances := make([]entities.AccountBalance, 0)
	rows, err := as.Connection.Query(ctx, query, args...)
	if err != nil {
		return accountBalances, entities.PageInfo{}, fmt.Errorf("querying account balances: %w", err)
	}
	defer rows.Close()

	if err = pgxscan.ScanAll(&accountBalances, rows); err != nil {
		return accountBalances, entities.PageInfo{}, fmt.Errorf("parsing account balances: %w", err)
	}

	pagedAccountBalances, pageInfo := entities.PageEntities[*v2.AccountEdge](accountBalances, pagination)
	return pagedAccountBalances, pageInfo, nil
}
