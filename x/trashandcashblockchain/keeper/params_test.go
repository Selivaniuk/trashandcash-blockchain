package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "trashandcash-blockchain/testutil/keeper"
	"trashandcash-blockchain/x/trashandcashblockchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TrashandcashblockchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
