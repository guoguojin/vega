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

package future

import (
	"context"
	"time"

	"code.vegaprotocol.io/vega/core/execution/common"
	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/logging"
)

func (m *Market) checkAuction(ctx context.Context, now time.Time, idgen common.IDGenerator) {
	if !m.as.InAuction() {
		// new block, check liquidity, start auction if needed
		m.checkLiquidity(ctx, nil, true)
		if m.as.AuctionStart() {
			m.enterAuction(ctx)
		}
		return
	}

	if m.mkt.State == types.MarketStateSuspendedViaGovernance {
		if endTS := m.as.ExpiresAt(); endTS != nil && endTS.Before(now) {
			m.as.ExtendAuctionSuspension(types.AuctionDuration{Duration: int64(m.minDuration)})
		}
	}

	// here we are in auction, we'll want to check
	// the triggers if we are leaving
	defer func() {
		m.triggerStopOrders(ctx, idgen)
	}()

	// as soon as we have an indicative uncrossing price in opening auction it needs to be passed into the price monitoring engine so statevar calculation can start
	isOpening := m.as.IsOpeningAuction()
	if isOpening && !m.pMonitor.Initialised() {
		trades, err := m.matching.GetIndicativeTrades()
		if err != nil {
			m.log.Panic("Can't get indicative trades")
		}
		if len(trades) > 0 {
			// pass the first uncrossing trades to price engine so state variables depending on it can be initialised
			m.pMonitor.CheckPrice(ctx, m.as, trades, true)
			m.OnOpeningAuctionFirstUncrossingPrice()
		}
	}

	if endTS := m.as.ExpiresAt(); endTS == nil || !endTS.Before(now) {
		return
	}
	trades, err := m.matching.GetIndicativeTrades()
	if err != nil {
		m.log.Panic("Can't get indicative trades")
	}

	// opening auction
	if isOpening {
		if len(trades) == 0 {
			return
		}

		// first check liquidity - before we mark auction as ready to leave
		m.checkLiquidity(ctx, trades, true)
		if !m.as.CanLeave() {
			if e := m.as.AuctionExtended(ctx, now); e != nil {
				m.broker.Send(e)
			}
			return
		}
		// opening auction requirements satisfied at this point, other requirements still need to be checked downstream though
		m.as.SetReadyToLeave()
		m.pMonitor.CheckPrice(ctx, m.as, trades, true)
		if m.as.ExtensionTrigger() == types.AuctionTriggerPrice {
			// this should never, ever happen
			m.log.Panic("Leaving opening auction somehow triggered price monitoring to extend the auction")
		}

		// if we don't have yet consensus for the floating point parameters, stay in the opening auction
		if !m.CanLeaveOpeningAuction() {
			m.log.Info("cannot leave opening auction - waiting for floating point to complete the first round")
			return
		}
		m.log.Info("leaving opening auction for market", logging.String("market-id", m.mkt.ID))
		m.leaveAuction(ctx, now)

		m.equityShares.OpeningAuctionEnded()
		// start the market fee window
		m.feeSplitter.TimeWindowStart(now)
		return
	}
	// price and liquidity auctions
	if endTS := m.as.ExpiresAt(); endTS == nil || !endTS.Before(now) {
		return
	}
	isPrice := m.as.IsPriceAuction() || m.as.IsPriceExtension()
	if !isPrice {
		m.checkLiquidity(ctx, trades, true)
	}
	if isPrice || m.as.CanLeave() {
		m.pMonitor.CheckPrice(ctx, m.as, trades, true)
	}
	end := m.as.CanLeave()
	if isPrice && end {
		m.checkLiquidity(ctx, trades, true)
	}
	if evt := m.as.AuctionExtended(ctx, m.timeService.GetTimeNow()); evt != nil {
		m.broker.Send(evt)
		end = false
	}
	// price monitoring engine and liquidity monitoring engine both indicated auction can end
	if end {
		// can we leave based on the book state?
		m.leaveAuction(ctx, now)
	}

	// This is where FBA handling will go
}
