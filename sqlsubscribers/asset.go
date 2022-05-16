package sqlsubscribers

import (
	"context"
	"math"
	"strconv"
	"time"

	"code.vegaprotocol.io/data-node/entities"
	"code.vegaprotocol.io/data-node/logging"
	"code.vegaprotocol.io/protos/vega"
	"code.vegaprotocol.io/vega/events"
	"github.com/pkg/errors"

	"github.com/shopspring/decimal"
)

type AssetEvent interface {
	events.Event
	Asset() vega.Asset
}

type AssetStore interface {
	Add(context.Context, entities.Asset) error
}

type Asset struct {
	subscriber
	store AssetStore
	log   *logging.Logger
}

func NewAsset(store AssetStore, log *logging.Logger) *Asset {
	return &Asset{
		store: store,
		log:   log,
	}
}

func (a *Asset) Types() []events.Type {
	return []events.Type{events.AssetEvent}
}

func (as *Asset) Push(ctx context.Context, evt events.Event) error {
	return as.consume(ctx, evt.(AssetEvent))
}

func (as *Asset) consume(ctx context.Context, ae AssetEvent) error {
	err := as.addAsset(ctx, ae.Asset(), as.vegaTime)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (as *Asset) addAsset(ctx context.Context, va vega.Asset, vegaTime time.Time) error {
	totalSupply, err := decimal.NewFromString(va.Details.TotalSupply)
	if err != nil {
		return errors.Errorf("bad total supply '%v'", va.Details.TotalSupply)
	}

	quantum, err := strconv.Atoi(va.Details.Quantum)
	if err != nil {
		return errors.Errorf("bad quantum '%v'", va.Details.Quantum)
	}

	var source, erc20Contract, lifetimeLimit, withdrawalThreshold string
	switch src := va.Details.Source.(type) {
	case *vega.AssetDetails_BuiltinAsset:
		source = src.BuiltinAsset.MaxFaucetAmountMint
	case *vega.AssetDetails_Erc20:
		erc20Contract = src.Erc20.ContractAddress
		lifetimeLimit = src.Erc20.LifetimeLimit
		withdrawalThreshold = src.Erc20.WithdrawThreshold
	default:
		return errors.Errorf("unknown asset source: %v", source)
	}

	if va.Details.Decimals > math.MaxInt {
		return errors.Errorf("decimals value will cause integer overflow: %d", va.Details.Decimals)
	}

	decimals := int(va.Details.Decimals)

	asset := entities.Asset{
		ID:                entities.NewAssetID(va.Id),
		Name:              va.Details.Name,
		Symbol:            va.Details.Symbol,
		TotalSupply:       totalSupply,
		Decimals:          decimals,
		Quantum:           quantum,
		Source:            source,
		ERC20Contract:     erc20Contract,
		VegaTime:          vegaTime,
		LifetimeLimit:     lifetimeLimit,
		WithdrawThreshold: withdrawalThreshold,
	}

	return errors.WithStack(as.store.Add(ctx, asset))
}
