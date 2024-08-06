package arbitrum

import (
	"context"

	"github.com/codecodeorg/arbgeth/arbitrum_types"
	"github.com/codecodeorg/arbgeth/core"
	"github.com/codecodeorg/arbgeth/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
