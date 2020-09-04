package abci

import (
	"context"

	"code.vegaprotocol.io/vega/blockchain"

	"github.com/tendermint/tendermint/abci/types"
)

type Tx interface {
	Command() blockchain.Command
	Payload() []byte
	PubKey() []byte
	Hash() []byte
	Validate() error
}

type Codec interface {
	Decode(in []byte) (Tx, error)
}

// ABCI hooks
type OnInitChainHandler func(types.RequestInitChain) types.ResponseInitChain
type OnBeginBlockHandler func(types.RequestBeginBlock) types.ResponseBeginBlock
type OnCheckTxHandler func(context.Context, types.RequestCheckTx) (context.Context, types.ResponseCheckTx)
type OnDeliverTxHandler func(context.Context, types.RequestDeliverTx) (context.Context, types.ResponseDeliverTx)
type OnCommitHandler func(types.RequestCommit) types.ResponseCommit
