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

//lint:file-ignore ST1003 Ignore underscores in names, this is straigh copied from the proto package to ease introducing the domain types

package types

import (
	"errors"
	"fmt"
	"strings"

	"code.vegaprotocol.io/vega/libs/num"
	proto "code.vegaprotocol.io/vega/protos/vega"
)

type LiquidityProviderFeeShares []*LiquidityProviderFeeShare

func (ls LiquidityProviderFeeShares) String() string {
	if ls == nil {
		return "[]"
	}
	strs := make([]string, 0, len(ls))
	for _, l := range ls {
		strs = append(strs, l.String())
	}
	return "[" + strings.Join(strs, ", ") + "]"
}

type LiquidityProviderFeeShare = proto.LiquidityProviderFeeShare

var (
	ErrNilTradableInstrument = errors.New("nil tradable instrument")
	ErrNilInstrument         = errors.New("nil instrument")
	ErrNilProduct            = errors.New("nil product")
	ErrUnknownAsset          = errors.New("unknown asset")
)

type MarketTimestamps struct {
	Proposed int64
	Pending  int64
	Open     int64
	Close    int64
}

func MarketTimestampsFromProto(p *proto.MarketTimestamps) *MarketTimestamps {
	var ts MarketTimestamps
	if p != nil {
		ts = MarketTimestamps{
			Proposed: p.Proposed,
			Pending:  p.Pending,
			Open:     p.Open,
			Close:    p.Close,
		}
	}
	return &ts
}

func (m MarketTimestamps) IntoProto() *proto.MarketTimestamps {
	return &proto.MarketTimestamps{
		Proposed: m.Proposed,
		Pending:  m.Pending,
		Open:     m.Open,
		Close:    m.Close,
	}
}

func (m MarketTimestamps) DeepClone() *MarketTimestamps {
	return &MarketTimestamps{
		Proposed: m.Proposed,
		Pending:  m.Pending,
		Open:     m.Open,
		Close:    m.Close,
	}
}

func (m MarketTimestamps) String() string {
	return fmt.Sprintf(
		"proposed(%v) open(%v) pending(%v) close(%v)",
		m.Proposed,
		m.Open,
		m.Pending,
		m.Close,
	)
}

type MarketTradingMode = proto.Market_TradingMode

const (
	// Default value, this is invalid.
	MarketTradingModeUnspecified MarketTradingMode = proto.Market_TRADING_MODE_UNSPECIFIED
	// Normal trading.
	MarketTradingModeContinuous MarketTradingMode = proto.Market_TRADING_MODE_CONTINUOUS
	// Auction trading (FBA).
	MarketTradingModeBatchAuction MarketTradingMode = proto.Market_TRADING_MODE_BATCH_AUCTION
	// Opening auction.
	MarketTradingModeOpeningAuction MarketTradingMode = proto.Market_TRADING_MODE_OPENING_AUCTION
	// Auction triggered by monitoring.
	MarketTradingModeMonitoringAuction MarketTradingMode = proto.Market_TRADING_MODE_MONITORING_AUCTION
	// No trading allowed.
	MarketTradingModeNoTrading MarketTradingMode = proto.Market_TRADING_MODE_NO_TRADING
	// Special auction mode for market suspended via governance.
	MarketTradingModeSuspendedViaGovernance MarketTradingMode = proto.Market_TRADING_MODE_SUSPENDED_VIA_GOVERNANCE
)

type MarketState = proto.Market_State

const (
	// Default value, invalid.
	MarketStateUnspecified MarketState = proto.Market_STATE_UNSPECIFIED
	// The Governance proposal valid and accepted.
	MarketStateProposed MarketState = proto.Market_STATE_PROPOSED
	// Outcome of governance votes is to reject the market.
	MarketStateRejected MarketState = proto.Market_STATE_REJECTED
	// Governance vote passes/wins.
	MarketStatePending MarketState = proto.Market_STATE_PENDING
	// Market triggers cancellation condition or governance
	// votes to close before market becomes Active.
	MarketStateCancelled MarketState = proto.Market_STATE_CANCELLED
	// Enactment date reached and usual auction exit checks pass.
	MarketStateActive MarketState = proto.Market_STATE_ACTIVE
	// Price monitoring or liquidity monitoring trigger.
	MarketStateSuspended MarketState = proto.Market_STATE_SUSPENDED
	// Governance vote (to close).
	MarketStateClosed MarketState = proto.Market_STATE_CLOSED
	// Defined by the product (i.e. from a product parameter,
	// specified in market definition, giving close date/time).
	MarketStateTradingTerminated MarketState = proto.Market_STATE_TRADING_TERMINATED
	// Settlement triggered and completed as defined by product.
	MarketStateSettled MarketState = proto.Market_STATE_SETTLED
	// Market has been suspended via a governance proposal.
	MarketStateSuspendedViaGovernance MarketState = proto.Market_STATE_SUSPENDED_VIA_GOVERNANCE
)

type AuctionTrigger = proto.AuctionTrigger

const (
	// Default value for AuctionTrigger, no auction triggered.
	AuctionTriggerUnspecified AuctionTrigger = proto.AuctionTrigger_AUCTION_TRIGGER_UNSPECIFIED
	// Batch auction.
	AuctionTriggerBatch AuctionTrigger = proto.AuctionTrigger_AUCTION_TRIGGER_BATCH
	// Opening auction.
	AuctionTriggerOpening AuctionTrigger = proto.AuctionTrigger_AUCTION_TRIGGER_OPENING
	// Price monitoring trigger.
	AuctionTriggerPrice AuctionTrigger = proto.AuctionTrigger_AUCTION_TRIGGER_PRICE
	// Liquidity monitoring due to unmet target trigger.
	AuctionTriggerLiquidityTargetNotMet AuctionTrigger = proto.AuctionTrigger_AUCTION_TRIGGER_LIQUIDITY_TARGET_NOT_MET
	// Liquidity monitoring due to being unable to deploy LP orders due to missing best bid or ask.
	AuctionTriggerUnableToDeployLPOrders AuctionTrigger = proto.AuctionTrigger_AUCTION_TRIGGER_UNABLE_TO_DEPLOY_LP_ORDERS
	// Governance triggered auction.
	AuctionTriggerGovernanceSuspension AuctionTrigger = proto.AuctionTrigger_AUCTION_TRIGGER_GOVERNANCE_SUSPENSION
)

type InstrumentMetadata struct {
	Tags []string
}

func InstrumentMetadataFromProto(m *proto.InstrumentMetadata) *InstrumentMetadata {
	return &InstrumentMetadata{
		Tags: append([]string{}, m.Tags...),
	}
}

func (i InstrumentMetadata) IntoProto() *proto.InstrumentMetadata {
	tags := make([]string, 0, len(i.Tags))
	return &proto.InstrumentMetadata{
		Tags: append(tags, i.Tags...),
	}
}

func (i InstrumentMetadata) String() string {
	return fmt.Sprintf(
		"tags(%v)",
		Tags(i.Tags).String(),
	)
}

func (i InstrumentMetadata) DeepClone() *InstrumentMetadata {
	ret := &InstrumentMetadata{
		Tags: make([]string, len(i.Tags)),
	}
	copy(ret.Tags, i.Tags)
	return ret
}

type AuctionDuration struct {
	Duration int64
	Volume   uint64
}

func AuctionDurationFromProto(ad *proto.AuctionDuration) *AuctionDuration {
	if ad == nil {
		return nil
	}
	return &AuctionDuration{
		Duration: ad.Duration,
		Volume:   ad.Volume,
	}
}

func (a AuctionDuration) IntoProto() *proto.AuctionDuration {
	return &proto.AuctionDuration{
		Duration: a.Duration,
		Volume:   a.Volume,
	}
}

func (a AuctionDuration) String() string {
	return fmt.Sprintf(
		"duration(%v) volume(%v)",
		a.Duration,
		a.Volume,
	)
}

func (a AuctionDuration) DeepClone() *AuctionDuration {
	return &AuctionDuration{
		Duration: a.Duration,
		Volume:   a.Volume,
	}
}

type rmType int

const (
	SimpleRiskModelType rmType = iota
	LogNormalRiskModelType
)

type TradableInstrument struct {
	Instrument       *Instrument
	MarginCalculator *MarginCalculator
	RiskModel        isTRM
	rmt              rmType
}

type isTRM interface {
	isTRM()
	trmIntoProto() interface{}
	rmType() rmType
	String() string
	Equal(isTRM) bool
}

func TradableInstrumentFromProto(ti *proto.TradableInstrument) *TradableInstrument {
	if ti == nil {
		return nil
	}
	rm := isTRMFromProto(ti.RiskModel)
	return &TradableInstrument{
		Instrument:       InstrumentFromProto(ti.Instrument),
		MarginCalculator: MarginCalculatorFromProto(ti.MarginCalculator),
		RiskModel:        rm,
		rmt:              rm.rmType(),
	}
}

func (t TradableInstrument) IntoProto() *proto.TradableInstrument {
	var (
		i *proto.Instrument
		m *proto.MarginCalculator
	)
	if t.Instrument != nil {
		i = t.Instrument.IntoProto()
	}
	if t.MarginCalculator != nil {
		m = t.MarginCalculator.IntoProto()
	}
	r := &proto.TradableInstrument{
		Instrument:       i,
		MarginCalculator: m,
	}
	if t.RiskModel == nil {
		return r
	}
	rmp := t.RiskModel.trmIntoProto()
	switch rm := rmp.(type) {
	case *proto.TradableInstrument_SimpleRiskModel:
		r.RiskModel = rm
	case *proto.TradableInstrument_LogNormalRiskModel:
		r.RiskModel = rm
	}
	return r
}

func (t TradableInstrument) GetSimpleRiskModel() *SimpleRiskModel {
	if t.rmt == SimpleRiskModelType {
		srm, ok := t.RiskModel.(*TradableInstrumentSimpleRiskModel)
		if !ok || srm == nil {
			return nil
		}
		return srm.SimpleRiskModel
	}
	return nil
}

func (t TradableInstrument) GetLogNormalRiskModel() *LogNormalRiskModel {
	if t.rmt == LogNormalRiskModelType {
		lrm, ok := t.RiskModel.(*TradableInstrumentLogNormalRiskModel)
		if !ok || lrm == nil {
			return nil
		}
		return lrm.LogNormalRiskModel
	}
	return nil
}

func (t TradableInstrument) String() string {
	return fmt.Sprintf(
		"instrument(%s) marginCalculator(%s) riskModel(%s)",
		reflectPointerToString(t.Instrument),
		reflectPointerToString(t.MarginCalculator),
		reflectPointerToString(t.RiskModel),
	)
}

func (t TradableInstrument) DeepClone() *TradableInstrument {
	ti := &TradableInstrument{
		Instrument: t.Instrument.DeepClone(),
		RiskModel:  t.RiskModel,
		rmt:        t.rmt,
	}
	if t.MarginCalculator != nil {
		ti.MarginCalculator = t.MarginCalculator.DeepClone()
	}
	return ti
}

type InstrumentSpot struct {
	Spot *Spot
}

func (InstrumentSpot) Type() ProductType {
	return ProductTypeSpot
}

func (i InstrumentSpot) String() string {
	return fmt.Sprintf(
		"spot(%s)",
		reflectPointerToString(i.Spot),
	)
}

type Spot struct {
	Name       string
	BaseAsset  string
	QuoteAsset string
}

func SpotFromProto(s *proto.Spot) *Spot {
	return &Spot{
		Name:       s.Name,
		BaseAsset:  s.BaseAsset,
		QuoteAsset: s.QuoteAsset,
	}
}

func (s Spot) IntoProto() *proto.Spot {
	return &proto.Spot{
		Name:       s.Name,
		BaseAsset:  s.BaseAsset,
		QuoteAsset: s.QuoteAsset,
	}
}

func (s Spot) String() string {
	return fmt.Sprintf(
		"baseAsset(%s) quoteAsset(%s)",
		s.BaseAsset,
		s.QuoteAsset,
	)
}

type InstrumentFuture struct {
	Future *Future
}

func (InstrumentFuture) Type() ProductType {
	return ProductTypeFuture
}

func (i InstrumentFuture) String() string {
	return fmt.Sprintf(
		"future(%s)",
		reflectPointerToString(i.Future),
	)
}

type Future struct {
	SettlementAsset                     string
	QuoteName                           string
	DataSourceSpecForSettlementData     *DataSourceSpec
	DataSourceSpecForTradingTermination *DataSourceSpec
	DataSourceSpecBinding               *DataSourceSpecBindingForFuture
}

func FutureFromProto(f *proto.Future) *Future {
	return &Future{
		SettlementAsset:                     f.SettlementAsset,
		QuoteName:                           f.QuoteName,
		DataSourceSpecForSettlementData:     DataSourceSpecFromProto(f.DataSourceSpecForSettlementData),
		DataSourceSpecForTradingTermination: DataSourceSpecFromProto(f.DataSourceSpecForTradingTermination),
		DataSourceSpecBinding:               DataSourceSpecBindingForFutureFromProto(f.DataSourceSpecBinding),
	}
}

func (f Future) IntoProto() *proto.Future {
	return &proto.Future{
		SettlementAsset:                     f.SettlementAsset,
		QuoteName:                           f.QuoteName,
		DataSourceSpecForSettlementData:     f.DataSourceSpecForSettlementData.IntoProto(),
		DataSourceSpecForTradingTermination: f.DataSourceSpecForTradingTermination.IntoProto(),
		DataSourceSpecBinding:               f.DataSourceSpecBinding.IntoProto(),
	}
}

func (f Future) String() string {
	return fmt.Sprintf(
		"quoteName(%s) settlementAsset(%s) dataSourceSpec(settlementData(%s) tradingTermination(%s) binding(%s))",
		f.QuoteName,
		f.SettlementAsset,
		reflectPointerToString(f.DataSourceSpecForSettlementData),
		reflectPointerToString(f.DataSourceSpecForTradingTermination),
		reflectPointerToString(f.DataSourceSpecBinding),
	)
}

func iInstrumentFromProto(pi interface{}) iProto {
	switch i := pi.(type) {
	case proto.Instrument_Future:
		return InstrumentFutureFromProto(&i)
	case *proto.Instrument_Future:
		return InstrumentFutureFromProto(i)
	case proto.Instrument_Spot:
		return InstrumentSpotFromProto(&i)
	case *proto.Instrument_Spot:
		return InstrumentSpotFromProto(i)
	}
	return nil
}

func InstrumentSpotFromProto(f *proto.Instrument_Spot) *InstrumentSpot {
	return &InstrumentSpot{
		Spot: SpotFromProto(f.Spot),
	}
}

func (i InstrumentSpot) IntoProto() *proto.Instrument_Spot {
	return &proto.Instrument_Spot{
		Spot: i.Spot.IntoProto(),
	}
}

func (i InstrumentSpot) getAssets() ([]string, error) {
	if i.Spot == nil {
		return []string{}, ErrUnknownAsset
	}
	return []string{i.Spot.BaseAsset, i.Spot.QuoteAsset}, nil
}

func (i InstrumentSpot) iIntoProto() interface{} {
	return i.IntoProto()
}

func InstrumentFutureFromProto(f *proto.Instrument_Future) *InstrumentFuture {
	return &InstrumentFuture{
		Future: FutureFromProto(f.Future),
	}
}

func (i InstrumentFuture) IntoProto() *proto.Instrument_Future {
	return &proto.Instrument_Future{
		Future: i.Future.IntoProto(),
	}
}

func (i InstrumentFuture) getAssets() ([]string, error) {
	if i.Future == nil {
		return []string{}, ErrUnknownAsset
	}
	return []string{i.Future.SettlementAsset}, nil
}

func (m *Market) GetAssets() ([]string, error) {
	if m.TradableInstrument == nil {
		return []string{}, ErrNilTradableInstrument
	}
	if m.TradableInstrument.Instrument == nil {
		return []string{}, ErrNilInstrument
	}
	if m.TradableInstrument.Instrument.Product == nil {
		return []string{}, ErrNilProduct
	}

	return m.TradableInstrument.Instrument.Product.getAssets()
}

func (m *Market) ProductType() ProductType {
	return m.TradableInstrument.Instrument.Product.Type()
}

func (m *Market) GetFuture() *InstrumentFuture {
	if m.ProductType() == ProductTypeFuture {
		f, _ := m.TradableInstrument.Instrument.Product.(*InstrumentFuture)
		return f
	}
	return nil
}

func (m *Market) GetSpot() *InstrumentSpot {
	if m.ProductType() == ProductTypeSpot {
		s, _ := m.TradableInstrument.Instrument.Product.(*InstrumentSpot)
		return s
	}
	return nil
}

func (i InstrumentFuture) iIntoProto() interface{} {
	return i.IntoProto()
}

type iProto interface {
	iIntoProto() interface{}
	getAssets() ([]string, error)
	String() string
	Type() ProductType
}

type Instrument struct {
	ID       string
	Code     string
	Name     string
	Metadata *InstrumentMetadata
	// Types that are valid to be assigned to Product:
	//	*InstrumentFuture
	//	*InstrumentSpot
	Product iProto
}

func InstrumentFromProto(i *proto.Instrument) *Instrument {
	if i == nil {
		return nil
	}
	return &Instrument{
		ID:       i.Id,
		Code:     i.Code,
		Name:     i.Name,
		Metadata: InstrumentMetadataFromProto(i.Metadata),
		Product:  iInstrumentFromProto(i.Product),
	}
}

func (i Instrument) GetSpot() *Spot {
	switch p := i.Product.(type) {
	case *InstrumentSpot:
		return p.Spot
	default:
		return nil
	}
}

func (i Instrument) GetFuture() *Future {
	switch p := i.Product.(type) {
	case *InstrumentFuture:
		return p.Future
	default:
		return nil
	}
}

func (i Instrument) IntoProto() *proto.Instrument {
	p := i.Product.iIntoProto()
	r := &proto.Instrument{
		Id:       i.ID,
		Code:     i.Code,
		Name:     i.Name,
		Metadata: i.Metadata.IntoProto(),
	}
	switch pt := p.(type) {
	case *proto.Instrument_Future:
		r.Product = pt

	case *proto.Instrument_Spot:
		r.Product = pt
	}
	return r
}

func (i Instrument) DeepClone() *Instrument {
	cpy := &Instrument{
		ID:      i.ID,
		Code:    i.Code,
		Name:    i.Name,
		Product: i.Product,
	}

	if i.Metadata != nil {
		cpy.Metadata = i.Metadata.DeepClone()
	}
	return cpy
}

func (i Instrument) String() string {
	return fmt.Sprintf(
		"ID(%s) name(%s) code(%s) product(%s) metadata(%s)",
		i.ID,
		i.Name,
		i.Code,
		reflectPointerToString(i.Product),
		reflectPointerToString(i.Metadata),
	)
}

type MarketData struct {
	MarkPrice                 *num.Uint
	LastTradedPrice           *num.Uint
	BestBidPrice              *num.Uint
	BestBidVolume             uint64
	BestOfferPrice            *num.Uint
	BestOfferVolume           uint64
	BestStaticBidPrice        *num.Uint
	BestStaticBidVolume       uint64
	BestStaticOfferPrice      *num.Uint
	BestStaticOfferVolume     uint64
	MidPrice                  *num.Uint
	StaticMidPrice            *num.Uint
	Market                    string
	Timestamp                 int64
	OpenInterest              uint64
	AuctionEnd                int64
	AuctionStart              int64
	IndicativePrice           *num.Uint
	IndicativeVolume          uint64
	MarketTradingMode         MarketTradingMode
	MarketState               MarketState
	Trigger                   AuctionTrigger
	ExtensionTrigger          AuctionTrigger
	TargetStake               string
	SuppliedStake             string
	PriceMonitoringBounds     []*PriceMonitoringBounds
	MarketValueProxy          string
	LiquidityProviderFeeShare []*LiquidityProviderFeeShare
	NextMTM                   int64
	MarketGrowth              num.Decimal
}

func (m MarketData) DeepClone() *MarketData {
	cpy := m
	cpy.MarkPrice = m.MarkPrice.Clone()
	cpy.LastTradedPrice = m.LastTradedPrice.Clone()
	cpy.BestBidPrice = m.BestBidPrice.Clone()
	cpy.BestOfferPrice = m.BestOfferPrice.Clone()
	cpy.BestStaticBidPrice = m.BestStaticBidPrice.Clone()
	cpy.BestStaticOfferPrice = m.BestStaticOfferPrice.Clone()
	cpy.MidPrice = m.MidPrice.Clone()
	cpy.StaticMidPrice = m.StaticMidPrice.Clone()
	cpy.IndicativePrice = m.IndicativePrice.Clone()

	cpy.PriceMonitoringBounds = make([]*PriceMonitoringBounds, 0, len(m.PriceMonitoringBounds))
	for _, pmb := range m.PriceMonitoringBounds {
		cpy.PriceMonitoringBounds = append(cpy.PriceMonitoringBounds, pmb.DeepClone())
	}
	lpfs := make([]*LiquidityProviderFeeShare, 0, len(m.LiquidityProviderFeeShare))
	for _, fs := range m.LiquidityProviderFeeShare {
		lpfs = append(lpfs, fs.DeepClone())
	}
	cpy.LiquidityProviderFeeShare = lpfs
	return &cpy
}

func (m MarketData) IntoProto() *proto.MarketData {
	r := &proto.MarketData{
		MarkPrice:                 num.UintToString(m.MarkPrice),
		LastTradedPrice:           num.UintToString(m.LastTradedPrice),
		BestBidPrice:              num.UintToString(m.BestBidPrice),
		BestBidVolume:             m.BestBidVolume,
		BestOfferPrice:            num.UintToString(m.BestOfferPrice),
		BestOfferVolume:           m.BestOfferVolume,
		BestStaticBidPrice:        num.UintToString(m.BestStaticBidPrice),
		BestStaticBidVolume:       m.BestStaticBidVolume,
		BestStaticOfferPrice:      num.UintToString(m.BestStaticOfferPrice),
		BestStaticOfferVolume:     m.BestStaticOfferVolume,
		MidPrice:                  num.UintToString(m.MidPrice),
		StaticMidPrice:            num.UintToString(m.StaticMidPrice),
		Market:                    m.Market,
		Timestamp:                 m.Timestamp,
		OpenInterest:              m.OpenInterest,
		AuctionEnd:                m.AuctionEnd,
		AuctionStart:              m.AuctionStart,
		IndicativePrice:           num.UintToString(m.IndicativePrice),
		IndicativeVolume:          m.IndicativeVolume,
		MarketTradingMode:         m.MarketTradingMode,
		MarketState:               m.MarketState,
		Trigger:                   m.Trigger,
		ExtensionTrigger:          m.ExtensionTrigger,
		TargetStake:               m.TargetStake,
		SuppliedStake:             m.SuppliedStake,
		PriceMonitoringBounds:     make([]*proto.PriceMonitoringBounds, 0, len(m.PriceMonitoringBounds)),
		MarketValueProxy:          m.MarketValueProxy,
		LiquidityProviderFeeShare: make([]*proto.LiquidityProviderFeeShare, 0, len(m.LiquidityProviderFeeShare)),
		NextMarkToMarket:          m.NextMTM,
		MarketGrowth:              m.MarketGrowth.String(),
	}
	for _, pmb := range m.PriceMonitoringBounds {
		r.PriceMonitoringBounds = append(r.PriceMonitoringBounds, pmb.IntoProto())
	}
	for _, lpfs := range m.LiquidityProviderFeeShare {
		r.LiquidityProviderFeeShare = append(r.LiquidityProviderFeeShare, lpfs.DeepClone()) // call IntoProto if this type gets updated
	}
	return r
}

func (m MarketData) String() string {
	return fmt.Sprintf(
		"markPrice(%s) lastTradedPrice(%s) bestBidPrice(%s) bestBidVolume(%v) bestOfferPrice(%s) bestOfferVolume(%v) bestStaticBidPrice(%s) bestStaticBidVolume(%v) bestStaticOfferPrice(%s) bestStaticOfferVolume(%v) midPrice(%s) staticMidPrice(%s) market(%s) timestamp(%v) openInterest(%v) auctionEnd(%v) auctionStart(%v) indicativePrice(%s) indicativeVolume(%v) marketTradingMode(%s) marketState(%s) trigger(%s) extensionTrigger(%s) targetStake(%s) suppliedStake(%s) priceMonitoringBounds(%s) marketValueProxy(%s) liquidityProviderFeeShare(%v) nextMTM(%v) marketGrowth(%v)",
		uintPointerToString(m.MarkPrice),
		uintPointerToString(m.LastTradedPrice),
		m.BestBidPrice.String(),
		m.BestBidVolume,
		uintPointerToString(m.BestOfferPrice),
		m.BestOfferVolume,
		uintPointerToString(m.BestStaticBidPrice),
		m.BestStaticBidVolume,
		uintPointerToString(m.BestStaticOfferPrice),
		m.BestStaticOfferVolume,
		uintPointerToString(m.MidPrice),
		uintPointerToString(m.StaticMidPrice),
		m.Market,
		m.Timestamp,
		m.OpenInterest,
		m.AuctionEnd,
		m.AuctionStart,
		uintPointerToString(m.IndicativePrice),
		m.IndicativeVolume,
		m.MarketTradingMode.String(),
		m.MarketState.String(),
		m.Trigger.String(),
		m.ExtensionTrigger.String(),
		m.TargetStake,
		m.SuppliedStake,
		PriceMonitoringBoundsList(m.PriceMonitoringBounds).String(),
		m.MarketValueProxy,
		LiquidityProviderFeeShares(m.LiquidityProviderFeeShare).String(),
		m.NextMTM,
		m.MarketGrowth,
	)
}

type Market struct {
	ID                            string
	TradableInstrument            *TradableInstrument
	DecimalPlaces                 uint64
	PositionDecimalPlaces         int64
	Fees                          *Fees
	OpeningAuction                *AuctionDuration
	PriceMonitoringSettings       *PriceMonitoringSettings
	LiquidityMonitoringParameters *LiquidityMonitoringParameters
	LPPriceRange                  num.Decimal
	LinearSlippageFactor          num.Decimal
	QuadraticSlippageFactor       num.Decimal
	LiquiditySLAParams            *LiquiditySLAParams

	TradingMode           MarketTradingMode
	State                 MarketState
	MarketTimestamps      *MarketTimestamps
	ParentMarketID        string
	InsurancePoolFraction num.Decimal
}

func MarketFromProto(mkt *proto.Market) (*Market, error) {
	lppr, _ := num.DecimalFromString(mkt.LpPriceRange)
	linearSlippageFactor, _ := num.DecimalFromString(mkt.LinearSlippageFactor)
	quadraticSlippageFactor, _ := num.DecimalFromString(mkt.QuadraticSlippageFactor)
	liquidityParameters, err := LiquidityMonitoringParametersFromProto(mkt.LiquidityMonitoringParameters)
	if err != nil {
		return nil, err
	}
	insFraction := num.DecimalZero()
	if mkt.InsurancePoolFraction != nil && len(*mkt.InsurancePoolFraction) > 0 {
		insFraction = num.MustDecimalFromString(*mkt.InsurancePoolFraction)
	}
	parent := ""
	if mkt.ParentMarketId != nil {
		parent = *mkt.ParentMarketId
	}

	m := &Market{
		ID:                            mkt.Id,
		TradableInstrument:            TradableInstrumentFromProto(mkt.TradableInstrument),
		DecimalPlaces:                 mkt.DecimalPlaces,
		PositionDecimalPlaces:         mkt.PositionDecimalPlaces,
		Fees:                          FeesFromProto(mkt.Fees),
		OpeningAuction:                AuctionDurationFromProto(mkt.OpeningAuction),
		PriceMonitoringSettings:       PriceMonitoringSettingsFromProto(mkt.PriceMonitoringSettings),
		LiquidityMonitoringParameters: liquidityParameters,
		TradingMode:                   mkt.TradingMode,
		State:                         mkt.State,
		MarketTimestamps:              MarketTimestampsFromProto(mkt.MarketTimestamps),
		LPPriceRange:                  lppr,
		LinearSlippageFactor:          linearSlippageFactor,
		QuadraticSlippageFactor:       quadraticSlippageFactor,
		ParentMarketID:                parent,
		InsurancePoolFraction:         insFraction,
	}

	if mkt.LiquiditySlaParams != nil {
		m.LiquiditySLAParams = LiquiditySLAParamsFromProto(mkt.LiquiditySlaParams)
	}

	return m, nil
}

func (m Market) IntoProto() *proto.Market {
	var (
		openAuct *proto.AuctionDuration
		mktTS    *proto.MarketTimestamps
		ti       *proto.TradableInstrument
		fees     *proto.Fees
		pms      *proto.PriceMonitoringSettings
		lms      *proto.LiquidityMonitoringParameters
	)
	if m.OpeningAuction != nil {
		openAuct = m.OpeningAuction.IntoProto()
	}
	if m.MarketTimestamps != nil {
		mktTS = m.MarketTimestamps.IntoProto()
	}
	if m.TradableInstrument != nil {
		ti = m.TradableInstrument.IntoProto()
	}
	if m.Fees != nil {
		fees = m.Fees.IntoProto()
	}
	if m.PriceMonitoringSettings != nil {
		pms = m.PriceMonitoringSettings.IntoProto()
	}
	if m.LiquidityMonitoringParameters != nil {
		lms = m.LiquidityMonitoringParameters.IntoProto()
	}
	var parent, insPoolFrac *string
	if len(m.ParentMarketID) != 0 {
		pid, insf := m.ParentMarketID, m.InsurancePoolFraction.String()
		parent = &pid
		insPoolFrac = &insf
	}

	var lpSLA *proto.LiquiditySLAParameters
	if m.LiquiditySLAParams != nil {
		lpSLA = m.LiquiditySLAParams.IntoProto()
	}

	r := &proto.Market{
		Id:                            m.ID,
		TradableInstrument:            ti,
		DecimalPlaces:                 m.DecimalPlaces,
		PositionDecimalPlaces:         m.PositionDecimalPlaces,
		Fees:                          fees,
		OpeningAuction:                openAuct,
		PriceMonitoringSettings:       pms,
		LiquidityMonitoringParameters: lms,
		TradingMode:                   m.TradingMode,
		State:                         m.State,
		MarketTimestamps:              mktTS,
		LpPriceRange:                  m.LPPriceRange.String(),
		LiquiditySlaParams:            lpSLA,
		LinearSlippageFactor:          m.LinearSlippageFactor.String(),
		QuadraticSlippageFactor:       m.QuadraticSlippageFactor.String(),
		InsurancePoolFraction:         insPoolFrac,
		ParentMarketId:                parent,
	}
	return r
}

func (m Market) GetID() string {
	return m.ID
}

func (m Market) String() string {
	return fmt.Sprintf(
		"ID(%s) tradableInstrument(%s) decimalPlaces(%v) positionDecimalPlaces(%v) fees(%s) openingAuction(%s) priceMonitoringSettings(%s) liquidityMonitoringParameters(%s) tradingMode(%s) state(%s) marketTimestamps(%s)",
		m.ID,
		reflectPointerToString(m.TradableInstrument),
		m.DecimalPlaces,
		m.PositionDecimalPlaces,
		reflectPointerToString(m.Fees),
		reflectPointerToString(m.OpeningAuction),
		reflectPointerToString(m.PriceMonitoringSettings),
		reflectPointerToString(m.LiquidityMonitoringParameters),
		m.TradingMode.String(),
		m.State.String(),
		reflectPointerToString(m.MarketTimestamps),
	)
}

func (m Market) DeepClone() *Market {
	cpy := &Market{
		ID:                      m.ID,
		DecimalPlaces:           m.DecimalPlaces,
		PositionDecimalPlaces:   m.PositionDecimalPlaces,
		TradingMode:             m.TradingMode,
		State:                   m.State,
		LPPriceRange:            m.LPPriceRange,
		LinearSlippageFactor:    m.LinearSlippageFactor,
		QuadraticSlippageFactor: m.QuadraticSlippageFactor,
		ParentMarketID:          m.ParentMarketID,
		InsurancePoolFraction:   m.InsurancePoolFraction,
	}

	if m.LiquiditySLAParams != nil {
		cpy.LiquiditySLAParams = m.LiquiditySLAParams.DeepClone()
	}

	if m.TradableInstrument != nil {
		cpy.TradableInstrument = m.TradableInstrument.DeepClone()
	}

	if m.Fees != nil {
		cpy.Fees = m.Fees.DeepClone()
	}

	if m.OpeningAuction != nil {
		cpy.OpeningAuction = m.OpeningAuction.DeepClone()
	}

	if m.PriceMonitoringSettings != nil {
		cpy.PriceMonitoringSettings = m.PriceMonitoringSettings.DeepClone()
	}

	if m.LiquidityMonitoringParameters != nil {
		cpy.LiquidityMonitoringParameters = m.LiquidityMonitoringParameters.DeepClone()
	}

	if m.MarketTimestamps != nil {
		cpy.MarketTimestamps = m.MarketTimestamps.DeepClone()
	}
	return cpy
}

type Tags []string

func (t Tags) String() string {
	return "[" + strings.Join(t, ", ") + "]"
}

func toPtr[T any](t T) *T { return &t }

type MarketCounters struct {
	StopOrderCounter    uint64
	PeggedOrderCounter  uint64
	LPShapeCount        uint64
	PositionCount       uint64
	OrderbookLevelCount uint64
}
