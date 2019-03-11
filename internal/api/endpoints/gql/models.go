// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"fmt"
	"io"
	"strconv"

	"code.vegaprotocol.io/vega/proto"
)

type TradingMode interface {
	IsTradingMode()
}

type ContinuousTrading struct {
	Void *bool `json:"void"`
}

func (ContinuousTrading) IsTradingMode() {}

type DiscreteTrading struct {
	Duration *int `json:"duration"`
}

func (DiscreteTrading) IsTradingMode() {}

type Market struct {
	Name        string            `json:"name"`
	TradingMode TradingMode       `json:"tradingMode"`
	Orders      []proto.Order     `json:"orders"`
	Trades      []proto.Trade     `json:"trades"`
	Depth       proto.MarketDepth `json:"depth"`
	Candles     []*proto.Candle   `json:"candles"`
}

type OrderFilter struct {
	And           []OrderFilter `json:"AND"`
	Or            []OrderFilter `json:"OR"`
	Open          *bool         `json:"open"`
	ID            *string       `json:"id"`
	IDNeq         *string       `json:"id_neq"`
	Market        *string       `json:"market"`
	MarketNeq     *string       `json:"market_neq"`
	Party         *string       `json:"party"`
	PartyNeq      *string       `json:"party_neq"`
	Side          *Side         `json:"side"`
	SideNeq       *Side         `json:"side_neq"`
	Price         *string       `json:"price"`
	PriceNeq      *string       `json:"price_neq"`
	PriceFrom     *string       `json:"price_from"`
	PriceTo       *string       `json:"price_to"`
	Size          *string       `json:"size"`
	SizeNeq       *string       `json:"size_neq"`
	SizeFrom      *string       `json:"size_from"`
	SizeTo        *string       `json:"size_to"`
	Remaining     *string       `json:"remaining"`
	RemainingNeq  *string       `json:"remaining_neq"`
	RemainingFrom *string       `json:"remaining_from"`
	RemainingTo   *string       `json:"remaining_to"`
	Type          *OrderType    `json:"type"`
	TypeNeq       *OrderType    `json:"type_neq"`
	Timestamp     *string       `json:"timestamp"`
	TimestampNeq  *string       `json:"timestamp_neq"`
	TimestampFrom *string       `json:"timestamp_from"`
	TimestampTo   *string       `json:"timestamp_to"`
	Status        *OrderStatus  `json:"status"`
	StatusNeq     *OrderStatus  `json:"status_neq"`
}

type Party struct {
	Name      string                 `json:"name"`
	Orders    []proto.Order          `json:"orders"`
	Trades    []proto.Trade          `json:"trades"`
	Positions []proto.MarketPosition `json:"positions"`
}

type PreConsensus struct {
	Accepted  bool   `json:"accepted"`
	Reference string `json:"reference"`
}

type TradeFilter struct {
	And           []TradeFilter `json:"AND"`
	Or            []TradeFilter `json:"OR"`
	ID            *string       `json:"id"`
	IDNeq         *string       `json:"id_neq"`
	Market        *string       `json:"market"`
	MarketNeq     *string       `json:"market_neq"`
	Buyer         *string       `json:"buyer"`
	BuyerNeq      *string       `json:"buyer_neq"`
	Seller        *string       `json:"seller"`
	SellerNeq     *string       `json:"seller_neq"`
	Aggressor     *Side         `json:"aggressor"`
	AggressorNeq  *Side         `json:"aggressor_neq"`
	Price         *string       `json:"price"`
	PriceNeq      *string       `json:"price_neq"`
	PriceFrom     *string       `json:"price_from"`
	PriceTo       *string       `json:"price_to"`
	Size          *string       `json:"size"`
	SizeNeq       *string       `json:"size_neq"`
	SizeFrom      *string       `json:"size_from"`
	SizeTo        *string       `json:"size_to"`
	Timestamp     *string       `json:"timestamp"`
	TimestampNeq  *string       `json:"timestamp_neq"`
	TimestampFrom *string       `json:"timestamp_from"`
	TimestampTo   *string       `json:"timestamp_to"`
}

type Vega struct {
	Markets []Market `json:"markets"`
	Market  *Market  `json:"market"`
	Parties []Party  `json:"parties"`
	Party   *Party   `json:"party"`
}

type Interval string

const (
	IntervalI1M  Interval = "I1M"
	IntervalI5M  Interval = "I5M"
	IntervalI15M Interval = "I15M"
	IntervalI1H  Interval = "I1H"
	IntervalI6H  Interval = "I6H"
	IntervalI1D  Interval = "I1D"
)

var AllInterval = []Interval{
	IntervalI1M,
	IntervalI5M,
	IntervalI15M,
	IntervalI1H,
	IntervalI6H,
	IntervalI1D,
}

func (e Interval) IsValid() bool {
	switch e {
	case IntervalI1M, IntervalI5M, IntervalI15M, IntervalI1H, IntervalI6H, IntervalI1D:
		return true
	}
	return false
}

func (e Interval) String() string {
	return string(e)
}

func (e *Interval) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Interval(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Interval", str)
	}
	return nil
}

func (e Interval) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderStatus string

const (
	OrderStatusActive    OrderStatus = "Active"
	OrderStatusCancelled OrderStatus = "Cancelled"
	OrderStatusExpired   OrderStatus = "Expired"
	OrderStatusStopped   OrderStatus = "Stopped"
	OrderStatusFilled    OrderStatus = "Filled"
)

var AllOrderStatus = []OrderStatus{
	OrderStatusActive,
	OrderStatusCancelled,
	OrderStatusExpired,
	OrderStatusStopped,
	OrderStatusFilled,
}

func (e OrderStatus) IsValid() bool {
	switch e {
	case OrderStatusActive, OrderStatusCancelled, OrderStatusExpired, OrderStatusStopped, OrderStatusFilled:
		return true
	}
	return false
}

func (e OrderStatus) String() string {
	return string(e)
}

func (e *OrderStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderStatus", str)
	}
	return nil
}

func (e OrderStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderType string

const (
	OrderTypeFok OrderType = "FOK"
	OrderTypeEne OrderType = "ENE"
	OrderTypeGtc OrderType = "GTC"
	OrderTypeGtt OrderType = "GTT"
)

var AllOrderType = []OrderType{
	OrderTypeFok,
	OrderTypeEne,
	OrderTypeGtc,
	OrderTypeGtt,
}

func (e OrderType) IsValid() bool {
	switch e {
	case OrderTypeFok, OrderTypeEne, OrderTypeGtc, OrderTypeGtt:
		return true
	}
	return false
}

func (e OrderType) String() string {
	return string(e)
}

func (e *OrderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderType", str)
	}
	return nil
}

func (e OrderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Side string

const (
	SideBuy  Side = "Buy"
	SideSell Side = "Sell"
)

var AllSide = []Side{
	SideBuy,
	SideSell,
}

func (e Side) IsValid() bool {
	switch e {
	case SideBuy, SideSell:
		return true
	}
	return false
}

func (e Side) String() string {
	return string(e)
}

func (e *Side) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Side(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Side", str)
	}
	return nil
}

func (e Side) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ValueDirection string

const (
	ValueDirectionPositive ValueDirection = "Positive"
	ValueDirectionNegative ValueDirection = "Negative"
)

var AllValueDirection = []ValueDirection{
	ValueDirectionPositive,
	ValueDirectionNegative,
}

func (e ValueDirection) IsValid() bool {
	switch e {
	case ValueDirectionPositive, ValueDirectionNegative:
		return true
	}
	return false
}

func (e ValueDirection) String() string {
	return string(e)
}

func (e *ValueDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ValueDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ValueDirection", str)
	}
	return nil
}

func (e ValueDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
