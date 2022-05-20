package sqlsubscribers

import (
	"context"
	"time"

	"code.vegaprotocol.io/data-node/entities"
	"code.vegaprotocol.io/data-node/logging"
	types "code.vegaprotocol.io/protos/vega"
	"code.vegaprotocol.io/vega/events"
	"github.com/pkg/errors"
)

type MarketDataEvent interface {
	events.Event
	MarketData() types.MarketData
}

//go:generate go run github.com/golang/mock/mockgen -destination mocks/market_data_mock.go -package mocks code.vegaprotocol.io/data-node/sqlsubscribers MarketDataStore
type MarketDataStore interface {
	Add(*entities.MarketData) error
	Flush(context.Context) error
}

type MarketData struct {
	subscriber
	log       *logging.Logger
	store     MarketDataStore
	dbTimeout time.Duration
}

func (md *MarketData) Flush(ctx context.Context) error {
	return md.store.Flush(ctx)
}

func (md *MarketData) Push(ctx context.Context, evt events.Event) error {
	return md.consume(evt.(MarketDataEvent))
}

func (md *MarketData) Types() []events.Type {
	return []events.Type{events.MarketDataEvent}
}

func NewMarketData(store MarketDataStore, log *logging.Logger) *MarketData {
	return &MarketData{
		log:   log,
		store: store,
	}
}

func (md *MarketData) consume(event MarketDataEvent) error {

	var record *entities.MarketData
	var err error
	mdProto := event.MarketData()

	if record, err = md.convertMarketDataProto(&mdProto, event.Sequence()); err != nil {
		errors.Wrap(err, "converting market data proto for persistence failed")
	}

	return errors.Wrap(md.store.Add(record), "inserting market data to SQL store failed")
}

func (md *MarketData) convertMarketDataProto(data *types.MarketData, seqNum uint64) (*entities.MarketData, error) {
	record, err := entities.MarketDataFromProto(data)
	if err != nil {
		return nil, err
	}

	record.SyntheticTime = md.vegaTime.Add(time.Duration(record.SeqNum) * time.Microsecond)
	record.VegaTime = md.vegaTime
	record.SeqNum = seqNum

	return record, nil
}
