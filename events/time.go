package events

import (
	"context"
	"time"

	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
)

// Time event indicating a change in block time (ie time update).
type Time struct {
	*Base
	blockTime time.Time
}

// NewTime returns a new time Update event.
func NewTime(ctx context.Context, t time.Time) *Time {
	return &Time{
		Base:      newBase(ctx, TimeUpdate),
		blockTime: t,
	}
}

// Time returns the new blocktime.
func (t Time) Time() time.Time {
	return t.blockTime
}

func (t Time) Proto() eventspb.TimeUpdate {
	return eventspb.TimeUpdate{
		Timestamp: t.blockTime.UnixNano(),
	}
}

func (t Time) StreamMessage() *eventspb.BusEvent {
	p := t.Proto()
	busEvent := newBusEventFromBase(t.Base)
	busEvent.Event = &eventspb.BusEvent_TimeUpdate{
		TimeUpdate: &p,
	}

	return busEvent
}

func TimeEventFromStream(ctx context.Context, be *eventspb.BusEvent) *Time {
	return &Time{
		Base:      newBaseFromBusEvent(ctx, TimeUpdate, be),
		blockTime: time.Unix(0, be.GetTimeUpdate().Timestamp),
	}
}
