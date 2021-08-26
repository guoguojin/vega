package coreapi

import (
	"context"
	"errors"

	coreapipb "code.vegaprotocol.io/protos/vega/coreapi/v1"
	"code.vegaprotocol.io/vega/broker"
	"code.vegaprotocol.io/vega/coreapi/services"
	"code.vegaprotocol.io/vega/logging"
)

var (
	ErrServiceDisabled = errors.New("service disabled")
)

type Service struct {
	ctx    context.Context
	broker broker.BrokerI
	cfg    Config
	log    *logging.Logger

	accounts  *services.Accounts
	assets    *services.Assets
	netparams *services.NetParams
}

func NewService(
	ctx context.Context, log *logging.Logger, cfg Config, broker broker.BrokerI,
) *Service {
	log = log.Named(namedLogger)
	log.SetLevel(cfg.LogLevel.Get())
	svc := &Service{
		broker: broker,
		cfg:    cfg,
		ctx:    ctx,
		log:    log,
	}

	if cfg.Accounts {
		log.Info("starting accounts core api")
		svc.accounts = services.NewAccounts(ctx)
		broker.SubscribeBatch(svc.accounts)
	}

	if cfg.Assets {
		log.Info("starting assets core api")
		svc.assets = services.NewAssets(ctx)
		broker.SubscribeBatch(svc.assets)
	}

	if cfg.NetworkParameters {
		log.Info("starting assets core api")
		svc.netparams = services.NewNetParams(ctx)
		broker.SubscribeBatch(svc.netparams)
	}

	return svc
}

func (s *Service) ListAccounts(
	ctx context.Context, in *coreapipb.ListAccountsRequest,
) (*coreapipb.ListAccountsResponse, error) {
	if !s.cfg.Accounts {
		return nil, ErrServiceDisabled
	}
	return &coreapipb.ListAccountsResponse{
		Accounts: s.accounts.List(in.Party, in.Market),
	}, nil
}

func (s *Service) ListAssets(
	ctx context.Context, in *coreapipb.ListAssetsRequest,
) (*coreapipb.ListAssetsResponse, error) {
	if !s.cfg.Assets {
		return nil, ErrServiceDisabled
	}
	return &coreapipb.ListAssetsResponse{
		Assets: s.assets.List(in.Asset),
	}, nil
}

func (s *Service) ListNetworkParameters(
	ctx context.Context, in *coreapipb.ListNetworkParametersRequest,
) (*coreapipb.ListNetworkParametersResponse, error) {
	if !s.cfg.NetworkParameters {
		return nil, ErrServiceDisabled
	}
	return &coreapipb.ListNetworkParametersResponse{
		NetworkParameters: s.netparams.List(in.NetworkParameterKey),
	}, nil
}
