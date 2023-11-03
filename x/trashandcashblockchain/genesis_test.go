package trashandcashblockchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "trashandcash-blockchain/testutil/keeper"
	"trashandcash-blockchain/testutil/nullify"
	"trashandcash-blockchain/x/trashandcashblockchain"
	"trashandcash-blockchain/x/trashandcashblockchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SpinList: []types.Spin{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SpinCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TrashandcashblockchainKeeper(t)
	trashandcashblockchain.InitGenesis(ctx, *k, genesisState)
	got := trashandcashblockchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SpinList, got.SpinList)
	require.Equal(t, genesisState.SpinCount, got.SpinCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
