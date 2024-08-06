package arbitrum

import (
	"context"

	"github.com/codecodeorg/arbgeth/common/hexutil"
	"github.com/codecodeorg/arbgeth/core"
	"github.com/codecodeorg/arbgeth/internal/ethapi"
	"github.com/codecodeorg/arbgeth/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *ethapi.StateOverride, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, overrides, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}
