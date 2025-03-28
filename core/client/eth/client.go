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

package eth

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"code.vegaprotocol.io/vega/core/types"
	vgcrypto "code.vegaprotocol.io/vega/libs/crypto"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrUnexpectedContractHash   = errors.New("hash of contract bytecode not as expected")
	ErrUnexpectedSolidityFormat = errors.New("unexpected format of solidity bytecode")
)

// ContractHashes the sha3-256(contract-bytecode stripped of metadata).
var ContractHashes = map[string]string{
	"staking":    "d66948e12817f8ae6ca94d56b43ca12e66416e7e9bc23bb09056957b25afc6bd",
	"vesting":    "5278802577f4aca315b9524bfa78790f8f0fae08939ec58bc9e8f0ea40123b09",
	"collateral": "1cd7f315188baf26f70c77a764df361c5d01bd365b109b96033b8755ee2b2750",
	"multisig":   "5b7070e6159628455b38f5796e8d0dc08185aaaa1fb6073767c88552d396c6c2",
}

// ETHClient ...
//
//go:generate go run github.com/golang/mock/mockgen -destination mocks/eth_client_mock.go -package mocks code.vegaprotocol.io/vega/core/client/eth ETHClient
type ETHClient interface { //revive:disable:exported
	bind.ContractBackend
	ethereum.ChainReader
	ChainID(context.Context) (*big.Int, error)
	NetworkID(context.Context) (*big.Int, error)
}

type Client struct {
	ETHClient
	ethConfig *types.EthereumConfig

	// this is all just to prevent spamming the infura just
	// to get the last height of the blockchain
	mu                      sync.Mutex
	currentHeightLastUpdate time.Time
	currentHeight           uint64

	cfg Config
}

func Dial(ctx context.Context, cfg Config) (*Client, error) {
	if len(cfg.RPCEndpoint) <= 0 {
		return nil, errors.New("no ethereum rpc endpoint configured. the configuration have move from the NodeWallet section to the Ethereum section, please make sure your vega configuration is up to date")
	}

	ethClient, err := ethclient.DialContext(ctx, cfg.RPCEndpoint)
	if err != nil {
		return nil, fmt.Errorf("couldn't instantiate Ethereum client: %w", err)
	}

	return &Client{ETHClient: ethClient, cfg: cfg}, nil
}

func (c *Client) UpdateEthereumConfig(ethConfig *types.EthereumConfig) error {
	if c == nil {
		return nil
	}

	netID, err := c.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't retrieve the network ID form the ethereum client: %w", err)
	}

	chainID, err := c.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't retrieve the chain ID form the ethereum client: %w", err)
	}

	if netID.String() != ethConfig.NetworkID() {
		return fmt.Errorf("updated network ID does not match the one set during start up, expected %s got %v", ethConfig.NetworkID(), netID)
	}

	if chainID.String() != ethConfig.ChainID() {
		return fmt.Errorf("updated chain ID does not matchthe one set during start up, expected %v got %v", ethConfig.ChainID(), chainID)
	}

	// if err := c.verifyStakingContract(context.Background(), ethConfig); err != nil {
	// 	return fmt.Errorf("failed to verify staking bridge contract: %w", err)
	// }

	// if err := c.verifyVestingContract(context.Background(), ethConfig); err != nil {
	// 	return fmt.Errorf("failed to verify vesting bridge contract: %w", err)
	// }

	if err := c.verifyCollateralContract(context.Background(), ethConfig); err != nil {
		return fmt.Errorf("failed to verify collateral bridge contract: %w", err)
	}

	if err := c.verifyMultisigContract(context.Background(), ethConfig); err != nil {
		return fmt.Errorf("failed to verify multisig control contract: %w", err)
	}

	c.ethConfig = ethConfig

	return nil
}

func (c *Client) CollateralBridgeAddress() ethcommon.Address {
	return c.ethConfig.CollateralBridge().Address()
}

func (c *Client) CollateralBridgeAddressHex() string {
	return c.ethConfig.CollateralBridge().HexAddress()
}

func (c *Client) CurrentHeight(ctx context.Context) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if now := time.Now(); c.currentHeightLastUpdate.Add(c.cfg.RetryDelay.Get()).Before(now) {
		lastBlockHeader, err := c.HeaderByNumber(ctx, nil)
		if err != nil {
			return c.currentHeight, err
		}
		c.currentHeightLastUpdate = now
		c.currentHeight = lastBlockHeader.Number.Uint64()
	}

	return c.currentHeight, nil
}

func (c *Client) ConfirmationsRequired() uint64 {
	return c.ethConfig.Confirmations()
}

// VerifyContract takes the address of a contract in hex and checks the hash of the byte-code is as expected.
func (c *Client) VerifyContract(ctx context.Context, address ethcommon.Address, expectedHash string) error {
	// nil block number means latest block
	b, err := c.CodeAt(ctx, address, nil)
	if err != nil {
		return err
	}

	// the bytecode of the contract is appended which is deployment specific. We only care about
	// the contract code itself and so we need to strip this meta-data before hashing it. For the version
	// of Solidity we use, the format is [contract-bytecode]a264[CBOR-encoded meta-data]
	asHex := strings.Split(hex.EncodeToString(b), "a264")
	if len(asHex) != 2 {
		return fmt.Errorf("%w: address: %s", ErrUnexpectedSolidityFormat, address)
	}

	// Back to bytes for hashing
	b, err = hex.DecodeString(asHex[0])
	if err != nil {
		return err
	}

	h := hex.EncodeToString(vgcrypto.Hash(b))
	if h != expectedHash {
		return fmt.Errorf("%w: address: %s, hash: %s, expected: %s", ErrUnexpectedContractHash, address, h, expectedHash)
	}

	return nil
}

// func (c *Client) verifyStakingContract(ctx context.Context, ethConfig *types.EthereumConfig) error {
// 	if address := ethConfig.StakingBridge(); address.HasAddress() {
// 		return c.VerifyContract(ctx, address.Address(), ContractHashes["staking"])
// 	}
// 	return nil
// }

// func (c *Client) verifyVestingContract(ctx context.Context, ethConfig *types.EthereumConfig) error {
// 	if address := ethConfig.VestingBridge(); address.HasAddress() {
// 		return c.VerifyContract(ctx, address.Address(), ContractHashes["vesting"])
// 	}
// 	return nil
// }

func (c *Client) verifyCollateralContract(ctx context.Context, ethConfig *types.EthereumConfig) error {
	if address := ethConfig.CollateralBridge(); address.HasAddress() {
		return c.VerifyContract(ctx, address.Address(), ContractHashes["collateral"])
	}

	return nil
}

func (c *Client) verifyMultisigContract(ctx context.Context, ethConfig *types.EthereumConfig) error {
	if address := ethConfig.MultiSigControl(); address.HasAddress() {
		return c.VerifyContract(ctx, address.Address(), ContractHashes["multisig"])
	}

	return nil
}
