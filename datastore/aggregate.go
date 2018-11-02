package datastore

import (
	"vega/msg"
	"fmt"
	"github.com/dgraph-io/badger"
)


type PriceHistory []*tradeInfo

type tradeInfo struct {
	timestamp uint64
	price     uint64
	size      uint64
}

func newCandle() *msg.Candle {
	return &msg.Candle{}
}

func (ts *tradeStore) GetCandle(market string, sinceBlock, currentBlock uint64) (*msg.Candle, error) {
	var candle *msg.Candle
	txn := ts.persistentStore.NewTransaction(false)
	candleKey := []byte(fmt.Sprintf("M:%s_C:%dB_T:%d", market, 1, sinceBlock))
	candleItem, err := txn.Get(candleKey)
	if err != nil {
		return nil, err
	}
	candleBuf, err := candleItem.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	candle.XXX_Unmarshal(candleBuf)
	return candle, nil
}
//
//	candle := &msg.Candle{
//		CloseBlockNumber: currentBlock,
//		OpenBlockNumber: sinceBlock,
//	}
//
//	for idx, t := range ts.store.markets[market].tradesByTimestamp {
//		// iterate trades until reached ones of interest
//		if t.trade.Timestamp < sinceBlock {
//			if t.trade.Price != 0 {
//				// keep updating empty candle with latest price so that in case there are no trades of interest,
//				// open close high and low values are set to the correct level of most recent trade
//				candle.Open = t.trade.Price
//				candle.Close = candle.Open
//				candle.High = candle.Open
//				candle.Low = candle.Open
//				candle.Volume = 0
//			}
//			continue
//		}
//
//		if candle.Open == 0 {
//			candle.Open = t.trade.Price
//		}
//
//		if candle.Volume == 0 {
//			candle.Open = t.trade.Price
//		}
//		candle.Volume += t.trade.Size
//		if candle.High < t.trade.Price {
//			candle.High = t.trade.Price
//		}
//		if candle.Low > t.trade.Price || candle.Low == 0 {
//			candle.Low = t.trade.Price
//		}
//		if idx == len(ts.store.markets[market].tradesByTimestamp)-1 {
//			candle.Close = t.trade.Price
//		}
//	}
//
//	return candle, nil
//}

// TODO: move into pregeneration

func (ts *tradeStore) GetCandles(market string, sinceBlock, currentBlock, interval uint64) ([]*msg.Candle, error) {
	var candles []*msg.Candle
	txn := ts.persistentStore.NewTransaction(false)
	it := txn.NewIterator(badger.DefaultIteratorOptions)
	defer it.Close()
	candlePrefix := []byte(fmt.Sprintf("M:%s_C:%dB_T:%d", market, 1, sinceBlock))
	for it.Seek(candlePrefix); it.ValidForPrefix(candlePrefix); it.Next() {
		candleItem := it.Item()
		candleBuf, _ := candleItem.ValueCopy(nil)

		var candle msg.Candle
		candle.XXX_Unmarshal(candleBuf)
		candles = append(candles, &candle)
	}
	return candles, nil
}

//	if err := ts.marketExists(market); err != nil {
//		return msg.Candles{}, err
//	}
//
//	nOfCandles := uint64(math.Ceil(float64((currentBlock-sinceBlock)/interval)))+1
//	var candles = make([]*msg.Candle, nOfCandles, nOfCandles)
//
//	for idx := range candles {
//		candles[idx] = &msg.Candle{}
//		candles[idx].OpenBlockNumber = sinceBlock + uint64(idx) * interval
//		candles[idx].CloseBlockNumber = candles[idx].OpenBlockNumber + interval - 1
//	}
//
//	found := false
//	idx := 0
//
//	for tidx, t := range ts.store.markets[market].tradesByTimestamp {
//		// iterate trades until reached ones of interest
//		if t.trade.Timestamp < sinceBlock {
//			continue
//		}
//
//		// OK I have now only trades I need
//		if candles[idx].OpenBlockNumber <= t.trade.Timestamp && t.trade.Timestamp <= candles[idx].CloseBlockNumber {
//			updateCandle(candles, idx, &t.trade)
//		} else {
//			// if current trade is not fit for current candle, close the candle with previous trade if non-empty candle
//			if candles[idx].Volume != 0 {
//				candles[idx].Close = ts.store.markets[market].tradesByTimestamp[tidx-1].trade.Price
//			}
//
//			// if we start from a candle that is empty, and there are no previous candles to copy close price
//			// its values should be populated with values of the previous trade that is outside of the sinceBlock scope
//			if idx == 0 && tidx > 0 && candles[idx].Volume == 0 {
//				candles[idx].Volume = 0
//				candles[idx].Open = ts.store.markets[market].tradesByTimestamp[tidx-1].trade.Price
//				candles[idx].Close = ts.store.markets[market].tradesByTimestamp[tidx-1].trade.Price
//				candles[idx].High = ts.store.markets[market].tradesByTimestamp[tidx-1].trade.Price
//				candles[idx].Low = ts.store.markets[market].tradesByTimestamp[tidx-1].trade.Price
//			}
//
//			// proceed to next candle
//			idx++
//			// otherwise look for next candle that fits to the current trade and add update candle with new trade
//			found = false
//			for !found {
//				// if reached the end of candles break
//				if idx > int(nOfCandles)-1 {
//					break
//				}
//				if candles[idx].OpenBlockNumber <= t.trade.Timestamp && t.trade.Timestamp <= candles[idx].CloseBlockNumber {
//					updateCandle(candles, idx, &t.trade)
//					found = true
//				} else {
//					// if candle is empty apply values from previous one
//					candles[idx].Volume = 0
//					if idx >= 1 {
//						candles[idx].Open = candles[idx-1].Close
//						candles[idx].Close = candles[idx-1].Close
//						candles[idx].High = candles[idx-1].Close
//						candles[idx].Low = candles[idx-1].Close
//					}
//					idx++
//				}
//			}
//			// if reached the end of candles break
//			if idx > int(nOfCandles)-1 {
//				break
//			}
//		}
//		candles[idx].Close = t.trade.Price
//	}
//
//	var output = msg.Candles{}
//	output.Candles = candles
//	return output, nil
//	return nil, nil
//}

//func updateCandle(candles []*msg.Candle, idx int, trade *msg.Trade) {
//	if candles[idx].Volume == 0 {
//		candles[idx].Open = trade.Price
//	}
//	candles[idx].Volume += trade.Size
//	if candles[idx].High < trade.Price {
//		candles[idx].High = trade.Price
//	}
//	if candles[idx].Low > trade.Price || candles[idx].Low == 0 {
//		candles[idx].Low = trade.Price
//	}
//}

func (m *orderStore) GetMarketDepth(market string) (*msg.MarketDepth, error) {
	//if err := m.marketExists(market); err != nil {
	//	return &msg.MarketDepth{}, err
	//}

	// get from store, recalculate accumulated volume and respond
	buy := m.orderBookDepth.getBuySide()
	sell := m.orderBookDepth.getSellSide()

	// recalculate accumulated volume
	for idx := range buy {
		if idx == 0 {
			buy[idx].CumulativeVolume = buy[idx].Volume
			continue
		}
		buy[idx].CumulativeVolume = buy[idx-1].CumulativeVolume + buy[idx].Volume
	}

	for idx := range m.orderBookDepth.getSellSide() {
		if idx == 0 {
			sell[idx].CumulativeVolume = sell[idx].Volume
			continue
		}
		sell[idx].CumulativeVolume = sell[idx-1].CumulativeVolume + sell[idx].Volume
	}

	orderBookDepth := msg.MarketDepth{Name: market, Buy: buy, Sell: sell}

	return &orderBookDepth, nil
}
