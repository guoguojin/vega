package api

import (
	"context"
	"vega/blockchain"
	"vega/core"
	"vega/datastore"
	"vega/msg"
	"github.com/pkg/errors"
	"vega/log"
)

type OrderService interface {
	Init(vega *core.Vega, orderStore datastore.OrderStore, eventStore datastore.EventStore)
	CreateOrder(ctx context.Context, order *msg.Order) (success bool, err error)
	CancelOrder(ctx context.Context, order *msg.Order) (success bool, err error)
	GetByMarket(ctx context.Context, market string, limit uint64) (orders []*msg.Order, err error)
	GetByParty(ctx context.Context, party string, limit uint64) (orders []*msg.Order, err error)
	GetByMarketAndId(ctx context.Context, market string, id string) (order *msg.Order, err error)
	GetByPartyAndId(ctx context.Context, market string, id string) (order *msg.Order, err error)
	GetMarkets(ctx context.Context) ([]string, error)
	GetMarketDepth(ctx context.Context, market string) (marketDepth *msg.MarketDepth, err error)
	Subscribe(ctx context.Context) (orders <-chan msg.Order, ref uint64)
	Unsubscribe(ctx context.Context, ref uint64) error
}

type orderService struct {
	app        *core.Vega
	orderStore datastore.OrderStore
	eventStore datastore.EventStore
	blockchain blockchain.Client
}

func NewOrderService() OrderService {
	return &orderService{}
}

func (p *orderService) Init(app *core.Vega, orderStore datastore.OrderStore, eventStore datastore.EventStore) {
	p.app = app
	p.orderStore = orderStore
	p.eventStore = eventStore
	p.blockchain = blockchain.NewClient()
}

func (p *orderService) CreateOrder(ctx context.Context, order *msg.Order) (success bool, err error) {
	// Set defaults, prevent unwanted external manipulation
	order.Remaining = order.Size
	order.Status = msg.Order_Active
	order.Type = msg.Order_GTC // VEGA only supports GTC at present
	order.Timestamp = 0
	order.RiskFactor = 0

	// TODO validate

	// Call out to the blockchain package/layer and use internal client to gain consensus
	return p.blockchain.CreateOrder(ctx, order)
}

// CancelOrder requires valid ID, Market, Party on an attempt to cancel the given active order via consensus
func (p *orderService) CancelOrder(ctx context.Context, order *msg.Order) (success bool, err error) {
	// Validate order exists using read store
	o, err := p.orderStore.GetByMarketAndId(order.Market, order.Id)
	if err != nil {
		return false, err
	}
	if o.Status == msg.Order_Cancelled {
		return false, errors.New("order has already been cancelled")
	}
	if o.Remaining == 0 {
		return false, errors.New("order has been fully filled")
	}
	if o.Party != order.Party {
		return false, errors.New("party mis-match cannot cancel order")
	}
	// Send cancellation request by consensus 
	return p.blockchain.CancelOrder(ctx, o.ToProtoMessage())
}

func (p *orderService) GetByMarket(ctx context.Context, market string, limit uint64) (orders []*msg.Order, err error) {
	o, err := p.orderStore.GetByMarket(market, datastore.GetOrderParams{Limit: limit})
	if err != nil {
		return nil, err
	}
	result := make([]*msg.Order, 0)
	for _, order := range o {
		//if order.Remaining == 0 {
		//	continue
		//}
		o := &msg.Order{
			Id:        order.Id,
			Market:    order.Market,
			Party:     order.Party,
			Side:      order.Side,
			Price:     order.Price,
			Size:      order.Size,
			Remaining: order.Remaining,
			Timestamp: order.Timestamp,
			Type:      order.Type,
			Status:    order.Status,
			RiskFactor:order.RiskFactor,
		}
		result = append(result, o)
	}
	return result, err
}

func (p *orderService) GetByParty(ctx context.Context, party string, limit uint64) (orders []*msg.Order, err error) {
	o, err := p.orderStore.GetByParty(party, datastore.GetOrderParams{Limit: limit})
	if err != nil {
		return nil, err
	}
	result := make([]*msg.Order, 0)
	for _, order := range o {
		//if order.Remaining == 0 {
		//	continue
		//}
		o := &msg.Order{
			Id:        order.Id,
			Market:    order.Market,
			Party:     order.Party,
			Side:      order.Side,
			Price:     order.Price,
			Size:      order.Size,
			Remaining: order.Remaining,
			Timestamp: order.Timestamp,
			Type:      order.Type,
			Status:    order.Status,
			RiskFactor:order.RiskFactor,
		}
		result = append(result, o)
	}
	return result, err
}

func (p *orderService) GetByMarketAndId(ctx context.Context, market string, id string) (order *msg.Order, err error) {
	o, err := p.orderStore.GetByMarketAndId(market, id)
	if err != nil {
		return &msg.Order{}, err
	}
	return o.ToProtoMessage(), err
}

func (p *orderService) GetByPartyAndId(ctx context.Context, market string, id string) (order *msg.Order, err error) {
	o, err := p.orderStore.GetByPartyAndId(market, id)
	if err != nil {
		return &msg.Order{}, err
	}
	return o.ToProtoMessage(), err
}

func (p *orderService) GetMarkets(ctx context.Context) ([]string, error) {
	markets, err := p.orderStore.GetMarkets()
	if err != nil {
		return []string{}, err
	}
	return markets, err
}

func (p *orderService) GetMarketDepth(ctx context.Context, marketName string) (orderBookDepth *msg.MarketDepth, err error) {
	return p.orderStore.GetMarketDepth(marketName)
}

func (p *orderService) Subscribe(ctx context.Context) (<-chan msg.Order, uint64) {
	orders := make(chan msg.Order)
	internal := make(chan []datastore.Order)
	ref := p.orderStore.Subscribe(internal)

	go func(id uint64, internal chan []datastore.Order) {
		<-ctx.Done()
		log.Debugf("OrderService -> Subscriber closed connection: %d", id)
		err := p.orderStore.Unsubscribe(id)
		if err != nil {
			log.Errorf("Error un-subscribing when context.Done() on OrderService for id: %d", id)
		}
		close(internal)
	}(ref, internal)

	go func(id uint64) {
		for v := range internal {
			for _, item := range v {
				orders <- *item.ToProtoMessage()
			}
		}
		log.Debugf("OrderService -> Channel for subscriber %d has been closed", ref)
	}(ref)

	return orders, ref
}

func (p *orderService) Unsubscribe(ctx context.Context, ref uint64) error {
	return p.orderStore.Unsubscribe(ref)
}