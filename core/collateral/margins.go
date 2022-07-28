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

package collateral

import (
	"code.vegaprotocol.io/vega/core/events"
	"code.vegaprotocol.io/vega/core/types/num"

	"code.vegaprotocol.io/vega/core/types"
)

type marginUpdate struct {
	events.MarketPosition
	margin          *types.Account
	general         *types.Account
	lock            *types.Account
	bond            *types.Account
	asset           string
	marketID        string
	marginShortFall *num.Uint
}

func (n marginUpdate) Transfer() *types.Transfer {
	return nil
}

func (n marginUpdate) Asset() string {
	return n.asset
}

func (n marginUpdate) MarketID() string {
	return n.marketID
}

func (n marginUpdate) MarginBalance() *num.Uint {
	if n.margin == nil {
		return num.Zero()
	}
	return n.margin.Balance.Clone()
}

// GeneralBalance here we cumulate both the general
// account and bon account so other package do not have
// to worry about how much funds are available in both
// if a bond account exists
// TODO(): maybe rename this method into AvailableBalance
// at some point if it makes senses overall the codebase.
func (n marginUpdate) GeneralBalance() *num.Uint {
	gen, bond := num.Zero(), num.Zero()
	if n.general != nil && n.general.Balance != nil {
		gen = n.general.Balance
	}
	if n.bond != nil && n.bond.Balance != nil {
		bond = n.bond.Balance
	}
	return num.Sum(bond, gen)
}

func (n marginUpdate) MarginShortFall() *num.Uint {
	return n.marginShortFall.Clone()
}

// BondBalance - returns nil if no bond account is present, *num.Uint otherwise.
func (n marginUpdate) BondBalance() *num.Uint {
	if n.bond == nil {
		return num.Zero()
	}
	return n.bond.Balance.Clone()
}
