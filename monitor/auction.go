package monitor

import (
	"context"
	"time"

	"code.vegaprotocol.io/vega/events"
	types "code.vegaprotocol.io/vega/proto"
)

// Trigger - placeholder type for proto enum
type Trigger int

const (
	Trigger_None Trigger = iota // none -> continuous trading
	Trigger_OpeningAuction
	Trigger_FBA
	Trigger_PriceMonitoring
	Trigger_LiquidityMonitoring
)

type AuctionState struct {
	mode        types.MarketState      // current trading mode
	defMode     types.MarketState      // default trading mode for market
	trigger     Trigger                // Set to the value indicating what started the auction
	begin       *time.Time             // optional setting auction start time (will be set if start flag is true)
	end         *types.AuctionDuration // will be set when in auction, defines parameters that end an auction period
	start, stop bool                   // flags to clarify whether we're entering or leaving auction
	m           *types.Market          // keep market definition handy, useful to end auctions when default is FBA
}

func NewAuctionState(mkt *types.Market, now time.Time) *AuctionState {
	s := AuctionState{
		mode:    types.MarketState_MARKET_STATE_AUCTION,
		defMode: types.MarketState_MARKET_STATE_CONTINUOUS,
		trigger: Trigger_OpeningAuction,
		begin:   &now,
		end:     mkt.OpeningAuction,
		start:   true,
	}
	if mkt.GetContinuous() == nil {
		s.defMode = types.MarketState_MARKET_STATE_AUCTION
	}
	// no opening auction
	if mkt.OpeningAuction == nil {
		s.mode = s.defMode
		s.begin = nil
		s.start = false
		s.trigger = Trigger_None
	}
	return &s
}

// StartLiquidityAuction - set the state to start a liquidity triggered auction
// @TODO these functions will be removed once the types are in proto
func (a *AuctionState) StartLiquidityAuction(t time.Time, d *types.AuctionDuration) {
	a.mode = types.MarketState_MARKET_STATE_AUCTION // auction mode
	a.trigger = Trigger_LiquidityMonitoring
	a.start = true
	a.stop = false
	a.begin = &t
	a.end = d
}

// StartPriceAuction - set the state to start a price triggered auction
// @TODO these functions will be removed once the types are in proto
func (a *AuctionState) StartPriceAuction(t time.Time, d *types.AuctionDuration) {
	a.mode = types.MarketState_MARKET_STATE_AUCTION // auction mode
	a.trigger = Trigger_PriceMonitoring
	a.start = true
	a.stop = false
	a.begin = &t
	a.end = d
}

// ExtendDuration - extend current auction, leaving trigger etc... in tact
func (a *AuctionState) ExtendAuction(delta types.AuctionDuration) {
	a.end.Duration += delta.Duration
	a.end.Volume += delta.Volume
	a.stop = false // the auction was supposed to stop, but we've extended it
}

// EndAuction is called by monitoring engines to mark if an auction period has expired
func (a *AuctionState) EndAuction() {
	a.stop = true
}

// Duration returns a copy of the current auction duration object
func (a AuctionState) Duration() types.AuctionDuration {
	if a.end == nil {
		return types.AuctionDuration{}
	}
	return *a.end
}

// Start - returns time pointer of the start of the auction (nil if not in auction)
func (a AuctionState) Start() time.Time {
	if a.begin == nil {
		return time.Time{} // zero time
	}
	return *a.begin
}

// ExpiresAt returns end as time -> if nil, the auction duration either isn't determined by time
// or we're simply not in an auction
func (a AuctionState) ExpiresAt() *time.Time {
	if a.begin == nil { // no start time == no end time
		return nil
	}
	if a.end == nil || a.end.Duration == 0 { // not time limited
		return nil
	}
	// add duration to start time, return
	t := a.begin.Add(time.Duration(a.end.Duration))
	return &t
}

// Mode returns current trading mode
func (a AuctionState) Mode() types.MarketState {
	return a.mode
}

// InAuction returns bool if the market is in auction for any reason
// Returns false if auction is triggered, but not yet started by market (execution)
func (a AuctionState) InAuction() bool {
	return (!a.start && a.trigger != Trigger_None)
}

func (a AuctionState) IsOpeningAuction() bool {
	return (a.trigger == Trigger_OpeningAuction)
}

func (a AuctionState) IsLiquidityAuction() bool {
	return (a.trigger == Trigger_LiquidityMonitoring)
}

func (a AuctionState) IsPriceAuction() bool {
	return (a.trigger == Trigger_PriceMonitoring)
}

func (a AuctionState) IsFBA() bool {
	return (a.trigger == Trigger_FBA)
}

// AuctionEnd bool indicating whether auction should be closed or not, if true, we can still extend the auction
// but when the market takes over (after monitoring engines), the auction will be closed
func (a AuctionState) AuctionEnd() bool {
	return a.stop
}

// AuctionStart bool indicates something has already triggered an auction to start, we can skip other monitoring potentially
// and we know to create an auction event
func (a AuctionState) AuctionStart() bool {
	return a.start
}

// AuctionStarted is called by the execution package to set flags indicating the market has started the auction
func (a *AuctionState) AuctionStarted(ctx context.Context) *events.Auction {
	a.start = false
	end := int64(0)
	if a.end != nil && a.end.Duration > 0 {
		end = a.begin.Add(time.Duration(a.end.Duration)).UnixNano()
	}
	return events.NewAuctionEvent(ctx, a.m.Id, false, a.begin.UnixNano(), end, a.IsOpeningAuction())
}

// AuctionEnded is called by execution to update internal state indicating this auction was closed
func (a *AuctionState) AuctionEnded(ctx context.Context, now time.Time) *events.Auction {
	// the end-of-auction event
	evt := events.NewAuctionEvent(ctx, a.m.Id, true, a.begin.UnixNano(), now.UnixNano(), a.IsOpeningAuction())
	a.start, a.stop = false, false
	a.begin, a.end = nil, nil
	a.trigger = Trigger_None
	a.mode = a.defMode
	// default mode is auction, this is an FBA market
	if a.mode == types.MarketState_MARKET_STATE_AUCTION {
		a.trigger = Trigger_FBA
	}
	return evt
}
