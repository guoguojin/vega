package banking

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"
)

var (
	ErrUnsupportedTransferKind = errors.New("unsupported transfer kind")
)

type scheduledTransfer struct {
	transfer    *types.Transfer
	accountType types.AccountType
	reference   string
}

func (e *Engine) OnTransferFeeFactorUpdate(ctx context.Context, f num.Decimal) error {
	e.transferFeeFactor = f
	return nil
}

func (e *Engine) TransferFunds(
	ctx context.Context,
	transfer *types.TransferFunds,
) error {
	switch transfer.Kind {
	case types.TransferCommandKindOneOff:
		return e.oneOffTransfer(ctx, transfer.OneOff)
	case types.TransferCommandKindRecurring:
		return e.recurringTransfer(ctx, transfer.Recurring)
	default:
		return ErrUnsupportedTransferKind
	}
}

func (e *Engine) oneOffTransfer(
	ctx context.Context,
	transfer *types.OneOffTransfer,
) error {
	// ensure asset exists
	if _, err := e.assets.Get(transfer.Asset); err != nil {
		e.log.Debug("cannot transfer funds, invalid asset", logging.Error(err))
		return fmt.Errorf("could not transfer funds: %w", err)
	}

	if err := transfer.IsValid(false); err != nil {
		return err
	}

	// ensure the party have enough funds for both the
	// amount and the fee for the transfer
	feeTransfer, err := e.ensureFeeForTransferFunds(
		transfer.Amount, transfer.From, transfer.Asset, transfer.FromAccountType)
	if err != nil {
		return fmt.Errorf("could not pay the fee for transfer: %w", err)
	}
	feeTransferAccountType := []types.AccountType{transfer.FromAccountType}

	fromTransfer, toTransfer := e.makeTransfers(
		transfer.From, transfer.To, transfer.Asset, transfer.Amount)
	transfers := []*types.Transfer{fromTransfer}
	accountTypes := []types.AccountType{transfer.FromAccountType}
	references := []string{transfer.Reference}

	// does the transfer needs to be finalized now?
	if transfer.DeliverOn == nil ||
		(transfer.DeliverOn != nil && transfer.DeliverOn.Before(e.currentTime)) {
		transfers = append(transfers, toTransfer)
		accountTypes = append(accountTypes, transfer.ToAccountType)
		references = append(references, transfer.Reference)
	} else {
		// schedule the transfer
		e.scheduleTransfer(
			toTransfer, transfer.ToAccountType, transfer.Reference, *transfer.DeliverOn)
	}

	// process the transfer
	tresps, err := e.col.TransferFunds(
		ctx, transfers, accountTypes, references, []*types.Transfer{feeTransfer}, feeTransferAccountType,
	)
	if err != nil {
		return err
	}

	e.broker.Send(events.NewTransferResponse(ctx, tresps))

	return nil
}

func (e *Engine) makeTransfers(
	from, to, asset string,
	amount *num.Uint,
) (*types.Transfer, *types.Transfer) {
	return &types.Transfer{
			Owner: from,
			Amount: &types.FinancialAmount{
				Amount: amount.Clone(),
				Asset:  asset,
			},
			Type:      types.TransferTypeTransferFundsSend,
			MinAmount: amount.Clone(),
		}, &types.Transfer{
			Owner: to,
			Amount: &types.FinancialAmount{
				Amount: amount.Clone(),
				Asset:  asset,
			},
			Type:      types.TransferTypeTransferFundsDistribute,
			MinAmount: amount.Clone(),
		}
}

func (e *Engine) ensureFeeForTransferFunds(
	amount *num.Uint,
	from, asset string,
	fromAccountType types.AccountType,
) (*types.Transfer, error) {
	// first we calculate the fee
	feeAmount, _ := num.UintFromDecimal(amount.ToDecimal().Mul(e.transferFeeFactor))

	// now we get the total amount and ensure we have enough funds
	// if the source account
	var (
		totalAmount = num.Zero().Add(feeAmount, amount)
		account     *types.Account
		err         error
	)
	switch fromAccountType {
	case types.AccountTypeGeneral:
		account, err = e.col.GetPartyGeneralAccount(from, asset)
		if err != nil {
			return nil, err
		}

	default:
		e.log.Panic("from account not supported",
			logging.String("account-type", fromAccountType.String()),
			logging.String("asset", asset),
			logging.String("from", from),
		)
	}

	if account.Balance.LT(totalAmount) {
		e.log.Debug("not enough funds to transfer",
			logging.BigUint("amount", amount),
			logging.BigUint("fee", feeAmount),
			logging.BigUint("total-amount", totalAmount),
			logging.BigUint("account-balance", account.Balance),
			logging.String("account-type", fromAccountType.String()),
			logging.String("asset", asset),
			logging.String("from", from),
		)
		return nil, ErrNotEnoughFundsToTransfer
	}

	return &types.Transfer{
		Owner: from,
		Amount: &types.FinancialAmount{
			Amount: feeAmount.Clone(),
			Asset:  asset,
		},
		Type:      types.TransferTypeInfrastructureFeePay,
		MinAmount: feeAmount.Clone(),
	}, nil
}

type timesToTransfers struct {
	deliverOn time.Time
	transfer  []scheduledTransfer
}

func (e *Engine) distributeScheduledTransfers(ctx context.Context) error {
	var ttfs = []timesToTransfers{}

	// iterate over those scheduled transfers to sort them by time
	for k, v := range e.scheduledTransfers {
		if k.Before(e.currentTime) {
			ttfs = append(ttfs, timesToTransfers{k, v})
			delete(e.scheduledTransfers, k)
		}
	}

	// sort slice by time.
	// no need to sort transfers they are going out as first in first out.
	sort.Slice(ttfs, func(i, j int) bool {
		return ttfs[i].deliverOn.Before(ttfs[j].deliverOn)
	})

	transfers := []*types.Transfer{}
	accountTypes := []types.AccountType{}
	references := []string{}
	for _, v := range ttfs {
		for _, t := range v.transfer {
			transfers = append(transfers, t.transfer)
			accountTypes = append(accountTypes, t.accountType)
			references = append(references, t.reference)
		}
	}

	if len(transfers) <= 0 {
		// nothing to do yeay
		return nil
	}

	tresps, err := e.col.TransferFunds(
		ctx, transfers, accountTypes, references, nil, nil, // no fees required there, they've been paid already
	)
	if err != nil {
		return err
	}

	e.broker.Send(events.NewTransferResponse(ctx, tresps))

	return nil
}

func (e *Engine) scheduleTransfer(
	t *types.Transfer, ty types.AccountType, reference string, deliverOn time.Time,
) {
	sts, ok := e.scheduledTransfers[deliverOn]
	if !ok {
		e.scheduledTransfers[deliverOn] = []scheduledTransfer{}
		sts = e.scheduledTransfers[deliverOn]
	}

	sts = append(sts, scheduledTransfer{
		transfer:    t,
		accountType: ty,
		reference:   reference,
	})
	e.scheduledTransfers[deliverOn] = sts
}

func (e *Engine) recurringTransfer(
	ctx context.Context,
	transfer *types.RecurringTransfer,
) error {
	return nil
}
