package trashandcashblockchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"trashandcash-blockchain/x/trashandcashblockchain/keeper"
	"trashandcash-blockchain/x/trashandcashblockchain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the spin
	for _, elem := range genState.SpinList {
		k.SetSpin(ctx, elem)
	}

	// Set spin count
	k.SetSpinCount(ctx, genState.SpinCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.SpinList = k.GetAllSpin(ctx)
	genesis.SpinCount = k.GetSpinCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
