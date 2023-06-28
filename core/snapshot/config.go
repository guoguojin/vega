// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package snapshot

import (
	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/libs/config/encoding"
	"code.vegaprotocol.io/vega/logging"
)

const (
	LevelDB    = "GOLevelDB"
	InMemoryDB = "memory"
)

type Config struct {
	Level       encoding.LogLevel `choice:"debug"                                                                                                                                                                    choice:"info"                 choice:"warning"                  choice:"error" choice:"panic" choice:"fatal" description:"Logging level (default: info)" long:"log-level"`
	KeepRecent  int               `description:"Number of historic snapshots to keep on disk. Limited to the 10 most recent ones"                                                                                    long:"snapshot-keep-recent"`
	RetryLimit  int               `description:"Maximum number of times to try and apply snapshot chunk"                                                                                                             long:"max-retries"`
	Storage     string            `choice:"GOLevelDB"                                                                                                                                                                choice:"memory"               description:"Storage type to use" long:"storage"`
	StartHeight int64             `description:"Load from a snapshot at the given block-height. Setting to -1 will load from the latest snapshot available, 0 will force the chain to replay if not using statesync" long:"load-from-block-height"`
}

// NewDefaultConfig creates an instance of the package specific configuration, given a
// pointer to a logger instance to be used for logging within the package.
func NewDefaultConfig() Config {
	return Config{
		Level:       encoding.LogLevel{Level: logging.InfoLevel},
		KeepRecent:  10,
		RetryLimit:  5,
		Storage:     LevelDB,
		StartHeight: -1,
	}
}

// validate checks the values in the config file are sensible, and returns the path
// which is create/load the snapshots from.
func (c *Config) validate() error {
	switch c.Storage {
	case InMemoryDB, LevelDB:
		return nil
	default:
		return types.ErrInvalidSnapshotStorageMethod
	}
}
