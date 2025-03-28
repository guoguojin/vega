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

package genesis

import (
	"encoding/json"

	tmtypes "github.com/tendermint/tendermint/types"
)

func FromJSON(rawGenesisDoc []byte) (*tmtypes.GenesisDoc, *State, error) {
	genesisDoc, err := tmtypes.GenesisDocFromJSON(rawGenesisDoc)
	if err != nil {
		return nil, nil, err
	}

	genesisState := &State{}
	err = json.Unmarshal(genesisDoc.AppState, genesisState)
	if err != nil {
		return nil, nil, err
	}
	return genesisDoc, genesisState, nil
}
