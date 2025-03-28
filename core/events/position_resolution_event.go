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

package events

import (
	"context"
	"fmt"

	"code.vegaprotocol.io/vega/libs/num"
	eventspb "code.vegaprotocol.io/vega/protos/vega/events/v1"
)

type PosRes struct {
	*Base
	distressed, closed int
	marketID           string
	markPrice          *num.Uint
}

func NewPositionResolution(ctx context.Context, distressed, closed int, markPrice *num.Uint, marketID string) *PosRes {
	base := newBase(ctx, PositionResolution)
	return &PosRes{
		Base:       base,
		distressed: distressed,
		closed:     closed,
		markPrice:  markPrice,
		marketID:   marketID,
	}
}

// MarketEvent implement the MarketEvent interface.
func (p PosRes) MarketEvent() string {
	return fmt.Sprintf("Market %s entered position resolution, %d parties were distressed, %d of which were closed out at mark price %s", p.marketID, p.distressed, p.closed, p.markPrice.String())
}

func (p PosRes) MarketID() string {
	return p.marketID
}

func (p PosRes) MarkPrice() *num.Uint {
	return p.markPrice.Clone()
}

func (p PosRes) Distressed() int {
	return p.distressed
}

func (p PosRes) Closed() int {
	return p.closed
}

func (p PosRes) Proto() eventspb.PositionResolution {
	return eventspb.PositionResolution{
		MarketId:   p.marketID,
		Closed:     int64(p.closed),
		Distressed: int64(p.distressed),
		MarkPrice:  p.markPrice.String(),
	}
}

func (p PosRes) MarketProto() eventspb.MarketEvent {
	return eventspb.MarketEvent{
		MarketId: p.marketID,
		Payload:  p.MarketEvent(),
	}
}

func (p PosRes) StreamMessage() *eventspb.BusEvent {
	pr := p.Proto()

	busEvent := newBusEventFromBase(p.Base)
	busEvent.Event = &eventspb.BusEvent_PositionResolution{
		PositionResolution: &pr,
	}

	return busEvent
}

func (p PosRes) StreamMarketMessage() *eventspb.BusEvent {
	msg := p.MarketProto()

	busEvent := newBusEventFromBase(p.Base)
	busEvent.Type = eventspb.BusEventType_BUS_EVENT_TYPE_MARKET
	busEvent.Event = &eventspb.BusEvent_Market{
		Market: &msg,
	}

	return busEvent
}

func PositionResolutionEventFromStream(ctx context.Context, be *eventspb.BusEvent) *PosRes {
	base := newBaseFromBusEvent(ctx, PositionResolution, be)
	mp, _ := num.UintFromString(be.GetPositionResolution().GetMarkPrice(), 10)
	return &PosRes{
		Base:       base,
		distressed: int(be.GetPositionResolution().Distressed),
		closed:     int(be.GetPositionResolution().Closed),
		markPrice:  mp,
		marketID:   be.GetPositionResolution().MarketId,
	}
}
