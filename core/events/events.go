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

package events

import (
	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/libs/num"
)

// MarketPosition is an event with a change to a position.
type MarketPosition interface {
	Party() string
	Size() int64
	Buy() int64
	Sell() int64
	Price() *num.Uint
	BuySumProduct() *num.Uint
	SellSumProduct() *num.Uint
	VWBuy() *num.Uint
	VWSell() *num.Uint
}

// TradeSettlement Part of the SettlePosition interface -> traces trades as they happened.
type TradeSettlement interface {
	Size() int64
	Price() *num.Uint
	MarketPrice() *num.Uint
}

// LossSocialization ...
type LossSocialization interface {
	MarketID() string
	PartyID() string
	AmountLost() int64
}

// SettlePosition is an event that the settlement buffer will propagate through the system
// used by the plugins (currently only the positions API).
type SettlePosition interface {
	MarketID() string
	Trades() []TradeSettlement
	Margin() (uint64, bool)
	Party() string
	Price() uint64
}

// FeeTransfer is a transfer initiated after trade occurs.
type FeesTransfer interface {
	// The list of transfers to be made by the collateral
	Transfers() []*types.Transfer
	// The total amount of fees to be paid (all cumulated)
	// per party if all the  transfers are to be executed
	// map is party id -> total amount of fees to be transferred
	TotalFeesAmountPerParty() map[string]*num.Uint
}

// Transfer is an event passed on by settlement engine, contains position
// and the resulting transfer for the collateral engine to use. We need MarketPosition
// because we can't loose the long/short status of the open positions.
type Transfer interface {
	MarketPosition
	Transfer() *types.Transfer
}

// Margin is an event with a change to balances after settling e.g. MTM.
type Margin interface {
	MarketPosition
	Asset() string
	MarginBalance() *num.Uint
	GeneralBalance() *num.Uint
	BondBalance() *num.Uint
	MarketID() string
	MarginShortFall() *num.Uint
}

// Risk is an event that summarizes everything and an eventual update to margin account.
type Risk interface {
	Margin
	Amount() *num.Uint
	Transfer() *types.Transfer // I know, it's included in the Transfer interface, but this is to make it clear that this particular func is masked at this level
	MarginLevels() *types.MarginLevels
}
