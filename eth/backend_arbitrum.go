package eth

import (
	"context"

	"github.com/codecodeorg/arbgeth/core"
	"github.com/codecodeorg/arbgeth/core/state"
	"github.com/codecodeorg/arbgeth/core/types"
	"github.com/codecodeorg/arbgeth/core/vm"
	"github.com/codecodeorg/arbgeth/eth/tracers"
	"github.com/codecodeorg/arbgeth/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}
