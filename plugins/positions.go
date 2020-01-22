package plugins

import (
	"context"
	"fmt"
	"sync"

	"code.vegaprotocol.io/vega/events"
	types "code.vegaprotocol.io/vega/proto"

	"github.com/pkg/errors"
)

var (
	ErrMarketNotFound = errors.New("could not find market")
	ErrPartyNotFound  = errors.New("party not found")
)

// PosBuffer ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/pos_buffer_mock.go -package mocks code.vegaprotocol.io/vega/plugins PosBuffer
type PosBuffer interface {
	Subscribe() (<-chan []events.SettlePosition, int)
	Unsubscribe(int)
}

// Positions plugin taking settlement data to build positions API data
type Positions struct {
	mu   *sync.RWMutex
	buf  PosBuffer
	ref  int
	ch   <-chan []events.SettlePosition
	data map[string]map[string]types.Position
}

func NewPositions(buf PosBuffer) *Positions {
	return &Positions{
		mu:   &sync.RWMutex{},
		data: map[string]map[string]types.Position{},
		buf:  buf,
	}
}

func (p *Positions) Start(ctx context.Context) {
	p.mu.Lock()
	if p.ch == nil {
		// get the channel and the reference
		p.ch, p.ref = p.buf.Subscribe()
		// start consuming the data
		go p.consume(ctx)
	}
	p.mu.Unlock()
}

func (p *Positions) Stop() {
	p.mu.Lock()
	if p.ch != nil {
		// only unsubscribe if ch was set, otherwise we might end up unregistering ref 0, which
		// could (in theory at least) be used by another component
		p.buf.Unsubscribe(p.ref)
		p.ch = nil
		p.ref = 0
	}
	// we don't need to reassign ch here, because the channel is closed, the consume routine
	// will pick up on the fact that we don't have to consume data anylonger, and the ch/ref fields
	// will be unset there
	p.mu.Unlock()
}

// consume keep reading the channel for as long as we need to
func (p *Positions) consume(ctx context.Context) {
	defer func() {
		p.mu.Lock()
		p.buf.Unsubscribe(p.ref)
		p.ref = 0
		p.ch = nil
		p.mu.Unlock()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case update, open := <-p.ch:
			if !open {
				return
			}
			p.mu.Lock()
			p.updateData(update)
			p.mu.Unlock()
		}
	}
}

func (p *Positions) updateData(raw []events.SettlePosition) {
	for _, sp := range raw {
		mID, tID := sp.MarketID(), sp.Party()
		if _, ok := p.data[mID]; !ok {
			p.data[mID] = map[string]types.Position{}
		}
		calc, ok := p.data[mID][tID]
		if !ok {
			calc = evtToProto(sp)
		}
		updatePosition(&calc, sp)
		p.data[mID][tID] = calc
	}
}

// GetPositionsByMarketAndParty get the position of a single trader in a given market
func (p *Positions) GetPositionsByMarketAndParty(market, party string) (*types.Position, error) {
	p.mu.RLock()
	mp, ok := p.data[market]
	if !ok {
		p.mu.RUnlock()
		return nil, nil
	}
	pos, ok := mp[party]
	if !ok {
		p.mu.RUnlock()
		pos = types.Position{
			PartyID:  party,
			MarketID: market,
		}
		return nil, nil
	}
	p.mu.RUnlock()
	return &pos, nil
}

// GetPositionsByParty get all positions for a given trader
func (p *Positions) GetPositionsByParty(party string) ([]*types.Position, error) {
	p.mu.RLock()
	// at most, trader is active in all markets
	positions := make([]*types.Position, 0, len(p.data))
	for _, traders := range p.data {
		if pos, ok := traders[party]; ok {
			positions = append(positions, &pos)
		}
	}
	p.mu.RUnlock()
	if len(positions) == 0 {
		return nil, nil
		// return nil, ErrPartyNotFound
	}
	return positions, nil
}

// GetPositionsByMarket get all trader positions in a given market
func (p *Positions) GetPositionsByMarket(market string) ([]*types.Position, error) {
	p.mu.RLock()
	mp, ok := p.data[market]
	if !ok {
		p.mu.RUnlock()
		return nil, ErrMarketNotFound
	}
	s := make([]*types.Position, 0, len(mp))
	for _, tp := range mp {
		s = append(s, &tp)
	}
	p.mu.RUnlock()
	return s, nil
}

/*
def calculate_open_closed_volume(self, traded_volume: int) -> (int, int):
    if self.open_volume != 0 and ((self.open_volume > 0) != (traded_volume > 0)):  # If sign of open volume and traded volume differ then some volume has been closed
      closed_volume = self.open_volume if abs(traded_volume) > abs(self.open_volume) else -traded_volume
      return traded_volume + closed_volume, closed_volume
    return traded_volume, 0
*/
func calculateOpenClosedVolume(currentOpenVolume, tradedVolume int64) (int64, int64) {
	fmt.Printf("tradedvolume(%v) openvolume(%v)\n", tradedVolume, currentOpenVolume)
	if currentOpenVolume != 0 && ((currentOpenVolume > 0) != (tradedVolume > 0)) {
		var closedVolume int64
		fmt.Printf("tradedvolume(%v) openvolume(%v)\n", tradedVolume, currentOpenVolume)
		if absUint64(tradedVolume) > absUint64(currentOpenVolume) {
			closedVolume = currentOpenVolume
		} else {
			fmt.Printf("traded volume: %v\n", -tradedVolume)
			closedVolume = -tradedVolume
		}
		return tradedVolume + closedVolume, closedVolume
	}
	return tradedVolume, 0
}

func closeV(p *types.Position, closedVolume int64, tradedPrice uint64) int64 {
	if closedVolume == 0 {
		return 0
	}
	realisedPnlDelta := closedVolume * int64(tradedPrice-p.AverageEntryPrice)
	p.RealisedPNL += int64(realisedPnlDelta)
	p.OpenVolume -= closedVolume
	return realisedPnlDelta
}

func updateVWAP(vwap uint64, volume int64, addVolume int64, addPrice uint64) uint64 {
	if volume+addVolume == 0 {
		return 0
	}
	return uint64((((int64(vwap) * volume) + (int64(addPrice) * addVolume)) / (volume + addVolume)))
}

func openV(p *types.Position, openedVolume int64, tradedPrice uint64) {
	p.AverageEntryPrice = updateVWAP(p.AverageEntryPrice, p.OpenVolume, openedVolume, tradedPrice)
	p.OpenVolume += openedVolume

}

func mtm(p *types.Position, markPrice uint64) {
	if p.OpenVolume == 0 {
		p.UnrealisedPNL = 0
		return
	}
	p.UnrealisedPNL = p.OpenVolume * int64(markPrice-p.AverageEntryPrice)
}

func updatePosition(p *types.Position, e events.SettlePosition) {
	// if this settlePosition event has a margin event embedded, that means we're dealing
	// with a trader who was closed out...
	if margin, ok := e.Margin(); ok {
		p.OpenVolume = 0
		p.UnrealisedPNL = 0
		p.AverageEntryPrice = 0
		// realised P&L includes whatever we had in margin account at this point
		p.RealisedPNL -= int64(margin.MarginBalance())
		// @TODO average entry price shouldn't be affected(?)
		// the volume now is zero, though, so we'll end up moving this position to storage
		return
	}
	for _, t := range e.Trades() {
		openedVolume, closedVolume := calculateOpenClosedVolume(p.OpenVolume, t.Size())
		_ = closeV(p, closedVolume, t.Price())
		openV(p, openedVolume, t.Price())
		mtm(p, t.Price())
	}
}

func evtToProto(e events.SettlePosition) types.Position {
	p := types.Position{
		MarketID: e.MarketID(),
		PartyID:  e.Party(),
	}
	// NOTE: We don't call this here because the call is made in updateEvt for all positions
	// we don't want to add the same data twice!
	// updatePosition(&p, e)
	return p
}

func absUint64(v int64) uint64 {
	if v < 0 {
		v *= -1
	}
	return uint64(v)
}
