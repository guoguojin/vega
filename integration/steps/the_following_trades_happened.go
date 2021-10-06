package steps

import (
	"fmt"
	"time"

	"github.com/cucumber/godog"

	"code.vegaprotocol.io/vega/integration/stubs"
)

func TheFollowingTradesShouldBeExecuted(
	broker *stubs.BrokerStub,
	table *godog.Table,
) error {
	var err error
	for _, row := range parseExecutedTradesTable(table) {
		buyer := row.MustStr("buyer")
		seller := row.MustStr("seller")
		price := row.MustU64("price")
		size := row.MustU64("size")
		aggressorRaw := row.Str("aggressor side")
		aggressor, aerr := Side(aggressorRaw)
		if aggressorRaw != "" && aerr != nil {
			return aerr
		}

		data := broker.GetTrades()
		var found bool
		for _, v := range data {
			if v.Buyer == buyer && v.Seller == seller && stringToU64(v.Price) == price && v.Size == size && (aggressorRaw == "" || aggressor == v.GetAggressor()) {
				found = true
			}
		}

		if !found {
			return errMissingTrade(buyer, seller, price, size)
		}
	}

	return err
}

func parseExecutedTradesTable(table *godog.Table) []RowWrapper {
	return StrictParseTable(table, []string{
		"buyer",
		"seller",
		"price",
		"size",
	}, []string{
		"aggressor side",
	})
}

// TheAuctionTradedVolumeAndPriceShouldBe pass in time at which the trades should happen in case there are previous trades in the broker stub.
func TheAuctionTradedVolumeAndPriceShouldBe(broker *stubs.BrokerStub, volume, price string, now time.Time) error {
	v, err := U64(volume)
	if err != nil {
		return err
	}
	p, err := U64(price)
	if err != nil {
		return err
	}
	// get all trades from stub
	trades := broker.GetTrades()
	sawV := uint64(0)
	for _, t := range trades {
		// no trades after the given time
		if t.Timestamp > now.UnixNano() {
			continue
		}
		if stringToU64(t.Price) != p {
			return fmt.Errorf(
				"expected trades to happen at price %d, instead saw a trade of size %d at price %s (%#v)",
				p, t.Size, t.Price, t,
			)
		}
		sawV += t.Size
	}
	if sawV != v {
		return fmt.Errorf(
			"expected a total traded volume of %d, instead saw a traded volume of %d len(%d): (%#v)",
			v, sawV, len(trades), trades,
		)
	}
	return nil
}

func errMissingTrade(buyer string, seller string, price uint64, volume uint64) error {
	return fmt.Errorf(
		"expecting trade was missing: buyer(%v), seller(%v), price(%v), volume(%v)",
		buyer, seller, price, volume,
	)
}
