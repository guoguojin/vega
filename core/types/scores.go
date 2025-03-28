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

package types

import "code.vegaprotocol.io/vega/libs/num"

type ScoreData struct {
	RawValScores      map[string]num.Decimal
	PerformanceScores map[string]num.Decimal
	MultisigScores    map[string]num.Decimal
	ValScores         map[string]num.Decimal
	NormalisedScores  map[string]num.Decimal
	NodeIDSlice       []string
}
