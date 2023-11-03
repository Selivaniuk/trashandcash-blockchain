package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "trashandcash-blockchain/testutil/keeper"
	"trashandcash-blockchain/testutil/nullify"
	"trashandcash-blockchain/x/trashandcashblockchain/keeper"
	"trashandcash-blockchain/x/trashandcashblockchain/types"
)

func createNSpin(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Spin {
	items := make([]types.Spin, n)
	for i := range items {
		items[i].Id = keeper.AppendSpin(ctx, items[i])
	}
	return items
}

func TestSpinGet(t *testing.T) {
	keeper, ctx := keepertest.TrashandcashblockchainKeeper(t)
	items := createNSpin(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSpin(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSpinRemove(t *testing.T) {
	keeper, ctx := keepertest.TrashandcashblockchainKeeper(t)
	items := createNSpin(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSpin(ctx, item.Id)
		_, found := keeper.GetSpin(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSpinGetAll(t *testing.T) {
	keeper, ctx := keepertest.TrashandcashblockchainKeeper(t)
	items := createNSpin(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSpin(ctx)),
	)
}

func TestSpinCount(t *testing.T) {
	keeper, ctx := keepertest.TrashandcashblockchainKeeper(t)
	items := createNSpin(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSpinCount(ctx))
}
