package subscribers

import (
	"context"

	"code.vegaprotocol.io/vega/events"
	types "code.vegaprotocol.io/vega/proto"
)

// NME - NewMarketEvent
type NME interface {
	Market() types.Market
}

type MarketStore interface {
	SaveBatch(markets []types.Market) error
}

type Market struct {
	*Base
	store MarketStore
}

func NewMarketSub(ctx context.Context, store MarketStore, ack bool) *Market {
	m := &Market{
		Base:  NewBase(ctx, 1, ack),
		store: store,
	}
	if m.isRunning() {
		go m.loop(m.ctx)
	}
	return m
}

func (m *Market) loop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			m.Halt()
			return
		case e := <-m.ch:
			if m.isRunning() {
				m.Push(e)
			}
		}
	}
}

func (m *Market) Push(e events.Event) {
	if te, ok := e.(NME); ok {
		m.store.SaveBatch([]types.Market{te.Market()})
	}
}

func (m *Market) Types() []events.Type {
	return []events.Type{
		events.MarketCreatedEvent,
	}
}
