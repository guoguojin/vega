package entities

import (
	"encoding/json"
	"fmt"
	"time"

	v2 "code.vegaprotocol.io/protos/data-node/api/v2"
	"code.vegaprotocol.io/protos/vega"
	"github.com/shopspring/decimal"
)

type DepositID struct{ ID }

func NewDepositID(id string) DepositID {
	return DepositID{ID: ID(id)}
}

type Deposit struct {
	ID                DepositID
	Status            DepositStatus
	PartyID           PartyID
	Asset             AssetID
	Amount            decimal.Decimal
	TxHash            string
	CreditedTimestamp time.Time
	CreatedTimestamp  time.Time
	VegaTime          time.Time
}

func DepositFromProto(deposit *vega.Deposit, vegaTime time.Time) (*Deposit, error) {
	var err error
	var amount decimal.Decimal

	if amount, err = decimal.NewFromString(deposit.Amount); err != nil {
		return nil, fmt.Errorf("invalid amount: %w", err)
	}

	return &Deposit{
		ID:                NewDepositID(deposit.Id),
		Status:            DepositStatus(deposit.Status),
		PartyID:           NewPartyID(deposit.PartyId),
		Asset:             NewAssetID(deposit.Asset),
		Amount:            amount,
		TxHash:            deposit.TxHash,
		CreditedTimestamp: NanosToPostgresTimestamp(deposit.CreditedTimestamp),
		CreatedTimestamp:  NanosToPostgresTimestamp(deposit.CreatedTimestamp),
		VegaTime:          vegaTime,
	}, nil
}

func (d Deposit) ToProto() *vega.Deposit {
	return &vega.Deposit{
		Id:                d.ID.String(),
		Status:            vega.Deposit_Status(d.Status),
		PartyId:           d.PartyID.String(),
		Asset:             d.Asset.String(),
		Amount:            d.Amount.String(),
		TxHash:            d.TxHash,
		CreditedTimestamp: d.CreditedTimestamp.UnixNano(),
		CreatedTimestamp:  d.CreatedTimestamp.UnixNano(),
	}
}

func (d Deposit) Cursor() *Cursor {
	cursor := DepositCursor{
		VegaTime: d.VegaTime,
		ID:       d.ID.String(),
	}
	return NewCursor(cursor.String())
}

func (d Deposit) ToProtoEdge(_ ...any) *v2.DepositEdge {
	return &v2.DepositEdge{
		Node:   d.ToProto(),
		Cursor: d.Cursor().Encode(),
	}
}

type DepositCursor struct {
	VegaTime time.Time `json:"vegaTime"`
	ID       string    `json:"id"`
}

func (dc DepositCursor) String() string {
	bs, err := json.Marshal(dc)
	if err != nil {
		return fmt.Sprintf(`{"vegaTime":"%s","id":"%s"}`, dc.VegaTime.Format(time.RFC3339Nano), dc.ID)
	}
	return string(bs)
}

func (dc *DepositCursor) Parse(cursorString string) error {
	if cursorString == "" {
		return nil
	}
	return json.Unmarshal([]byte(cursorString), dc)
}
