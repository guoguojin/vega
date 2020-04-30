package orders_test

import (
	"context"
	"testing"
	"time"

	"code.vegaprotocol.io/vega/proto"
	"code.vegaprotocol.io/vega/vegatime"
	"github.com/stretchr/testify/assert"
)

var (
	amend = proto.OrderAmendment{
		OrderID:   "order_id",
		PartyID:   "party",
		Price:     &proto.Price{Value: 10000},
		SizeDelta: 1,
		MarketID:  "market",
	}
)

type amendMatcher struct {
	e proto.OrderAmendment
}

func TestPrepareAmendOrder(t *testing.T) {
	t.Run("Prepare amend order price - success", testPrepareAmendOrderJustPriceSuccess)
	t.Run("Prepare amend order reduce - success", testPrepareAmendOrderJustReduceSuccess)
	t.Run("Prepare amend order increase - success", testPrepareAmendOrderJustIncreaseSuccess)
	t.Run("Prepare amend order expiry - success", testPrepareAmendOrderJustExpirySuccess)
	t.Run("Prepare amend order tif - success", testPrepareAmendOrderJustTIFSuccess)

	t.Run("Prepare amend order empty - fail", testPrepareAmendOrderEmptyFail)
	t.Run("Prepare amend order nil - fail", testPrepareAmendOrderNilFail)
	t.Run("Prepare amend order invalid expiry type - fail", testPrepareAmendOrderInvalidExpiryFail)
}

func testPrepareAmendOrderJustPriceSuccess(t *testing.T) {
	arg := proto.OrderAmendment{
		OrderID: "orderid",
		PartyID: "partyid",
		Price:   &proto.Price{Value: 1000},
	}
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	err := svc.svc.PrepareAmendOrder(context.Background(), &arg)
	assert.NoError(t, err)
}

func testPrepareAmendOrderJustReduceSuccess(t *testing.T) {
	arg := proto.OrderAmendment{
		OrderID:   "orderid",
		PartyID:   "partyid",
		SizeDelta: -10,
	}
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	err := svc.svc.PrepareAmendOrder(context.Background(), &arg)
	assert.NoError(t, err)
}

func testPrepareAmendOrderJustIncreaseSuccess(t *testing.T) {
	arg := proto.OrderAmendment{
		OrderID:   "orderid",
		PartyID:   "partyid",
		SizeDelta: 10,
	}
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	err := svc.svc.PrepareAmendOrder(context.Background(), &arg)
	assert.NoError(t, err)
}

func testPrepareAmendOrderJustExpirySuccess(t *testing.T) {
	now := vegatime.Now()
	expires := now.Add(-2 * time.Hour)
	arg := proto.OrderAmendment{
		OrderID:   "orderid",
		PartyID:   "partyid",
		ExpiresAt: &proto.Timestamp{Value: expires.UnixNano()},
	}
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	err := svc.svc.PrepareAmendOrder(context.Background(), &arg)
	assert.NoError(t, err)
}

func testPrepareAmendOrderJustTIFSuccess(t *testing.T) {
	arg := proto.OrderAmendment{
		OrderID:     "orderid",
		PartyID:     "partyid",
		TimeInForce: proto.Order_GTC,
	}
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	err := svc.svc.PrepareAmendOrder(context.Background(), &arg)
	assert.NoError(t, err)
}

func testPrepareAmendOrderEmptyFail(t *testing.T) {
	arg := proto.OrderAmendment{}
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	err := svc.svc.PrepareAmendOrder(context.Background(), &arg)
	assert.Error(t, err)

	arg2 := proto.OrderAmendment{
		OrderID: "orderid",
		PartyID: "partyid",
	}
	err = svc.svc.PrepareAmendOrder(context.Background(), &arg2)
	assert.Error(t, err)
}

func testPrepareAmendOrderNilFail(t *testing.T) {
	var arg proto.OrderAmendment
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	err := svc.svc.PrepareAmendOrder(context.Background(), &arg)
	assert.Error(t, err)
}

func testPrepareAmendOrderInvalidExpiryFail(t *testing.T) {
	arg := proto.OrderAmendment{
		OrderID:     "orderid",
		PartyID:     "partyid",
		TimeInForce: proto.Order_GTC,
		ExpiresAt:   &proto.Timestamp{Value: 10},
	}
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	err := svc.svc.PrepareAmendOrder(context.Background(), &arg)
	assert.Error(t, err)

	arg.TimeInForce = proto.Order_FOK
	err = svc.svc.PrepareAmendOrder(context.Background(), &arg)
	assert.Error(t, err)

	arg.TimeInForce = proto.Order_IOC
	err = svc.svc.PrepareAmendOrder(context.Background(), &arg)
	assert.Error(t, err)
}
