package leadscanner

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go-onchain-leads/internal/domain"
	"math/big"
)

type BlockchainReader interface {
	NetworkID(ctx context.Context) (*big.Int, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

type IdentityResolver interface {
	ResolveIdentity(walletAddress string) domain.Identity
}
