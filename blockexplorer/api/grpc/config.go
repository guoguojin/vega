// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package grpc

import (
	"code.vegaprotocol.io/vega/libs/config/encoding"
	"code.vegaprotocol.io/vega/logging"
)

var namedLogger = "api.grpc"

type Config struct {
	Reflection         encoding.Bool     `description:"Enable GRPC reflection, required for grpc-ui"                              long:"reflection"`
	Level              encoding.LogLevel `choice:"debug"                                                                          choice:"info"            choice:"warning" long:"log-level"`
	MaxPageSizeDefault uint32            `description:"How many results to return per page if client does not specify explicitly" long:"default-page-size"`
}

func NewDefaultConfig() Config {
	return Config{
		Reflection:         encoding.Bool(true),
		Level:              encoding.LogLevel{Level: logging.InfoLevel},
		MaxPageSizeDefault: 50,
	}
}
