package mockbridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type GlobalInboxWatcher struct {
	client arbbridge.ArbClient
}

func NewGlobalInboxWatcher(client arbbridge.ArbClient) (*EthRollupWatcher, error) {
	//vm := &EthRollupWatcher{Client: client.(*ArbClient).client, address: address}
	//err := vm.setupContracts()
	//return vm, err
	return &EthRollupWatcher{client: client}, nil
}

func (vm *GlobalInboxWatcher) GetEvents(ctx context.Context, blockId *common.BlockId, timestamp *big.Int) ([]arbbridge.Event, error) {
	return nil, nil
}

func (vm *GlobalInboxWatcher) GetAllReceived(
	ctx context.Context,
	fromBlock *big.Int,
	toBlock *big.Int,
) ([]message.Received, error) {
	return nil, nil
}
