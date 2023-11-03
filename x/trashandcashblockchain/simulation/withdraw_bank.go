package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"trashandcash-blockchain/x/trashandcashblockchain/keeper"
	"trashandcash-blockchain/x/trashandcashblockchain/types"
)

func SimulateMsgWithdrawBank(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgWithdrawBank{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the WithdrawBank simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "WithdrawBank simulation not implemented"), nil, nil
	}
}
