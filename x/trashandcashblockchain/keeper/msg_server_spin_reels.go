package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"trashandcash-blockchain/x/trashandcashblockchain/types"
)

func GetMultiplier(result string) (float64, error) {
	digit1, err := strconv.Atoi(string(result[0]))
	if err != nil {
		return 0, err
	}
	digit2, err := strconv.Atoi(string(result[1]))
	if err != nil {
		return 0, err
	}
	digit3, err := strconv.Atoi(string(result[2]))
	if err != nil {
		return 0, err
	}

	if result == "777" {
		return 100.0, nil
	} else if result == "333" {
		return 50.0, nil
	} else if digit1 == digit2 && digit1 == digit3 {
		return 15.0, nil
	} else if digit1 == digit2 || digit2 == digit3 {
		return 5.0, nil
	}
	return 0, nil
}
func TransformDigit(digit int) int {
	if digit > 0 {
		return digit - 1
	}
	return 7
}

func GetLosingResult(result string) (string, error) {
	digit1, err := strconv.Atoi(string(result[0]))
	if err != nil {
		return result, err
	}

	digit2, err := strconv.Atoi(string(result[1]))
	if err != nil {
		return result, err
	}

	digit3, err := strconv.Atoi(string(result[2]))
	if err != nil {
		return result, err
	}

	if digit3 == digit2 {
		combination := strconv.Itoa(digit1) + strconv.Itoa(digit2) + strconv.Itoa(TransformDigit(digit3))
		return GetLosingResult(combination)
	} else if digit1 == digit2 {
		combination := strconv.Itoa(digit1) + strconv.Itoa(TransformDigit(digit2)) + strconv.Itoa(digit3)
		return GetLosingResult(combination)
	}
	return result, nil
}

func (k msgServer) SpinReels(goCtx context.Context, msg *types.MsgSpinReels) (*types.MsgSpinReelsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	player, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	moduleAcc := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAcc == nil {
		return nil, fmt.Errorf("module account does not exist")
	}

	bet, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}
	fmt.Println("bet ", bet.String())

	result, err := k.GenerateResult(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("result ", result)

	multiplier, err := GetMultiplier(result)
	if err != nil {
		return nil, err
	}

	fmt.Println("multiplier ", multiplier)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, player, types.ModuleName, sdk.NewCoins(bet))
	if err != nil {
		return nil, err
	}

	var winCoins sdk.Coins
	if multiplier > 0 {
		betFloat := sdk.NewDecFromInt(bet.Amount)
		winAmountFloat := betFloat.Mul(sdk.NewDecWithPrec(int64(multiplier*100), 2))
		winAmount := winAmountFloat.RoundInt()
		winCoin := sdk.NewCoin(bet.Denom, winAmount)
		winCoins = sdk.NewCoins(winCoin)

		if !k.bankKeeper.HasBalance(ctx, moduleAcc, winCoin) {
			result, err = GetLosingResult(result)
			if err != nil {
				return nil, err
			}

			multiplier = 0
			winCoin = sdk.NewCoin(bet.Denom, sdk.ZeroInt())
			winCoins = sdk.NewCoins(winCoin)
		} else {
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, player, winCoins)
			if err != nil {
				return nil, err
			}
		}
	} else {
		winCoins = sdk.NewCoins(sdk.NewCoin(bet.Denom, sdk.ZeroInt()))
	}

	var spin = types.Spin{
		Bet:     msg.Amount,
		Address: msg.Creator,
		Result:  result,
		Winning: winCoins.String(),
	}
	spinId := k.AppendSpin(ctx, spin)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.SpinReelsEventType,
			sdk.NewAttribute(types.SpinReelsEventId, fmt.Sprint(spinId)),
			sdk.NewAttribute(types.SpinReelsEventBet, spin.Bet),
			sdk.NewAttribute(types.SpinReelsEventAddress, spin.Address),
			sdk.NewAttribute(types.SpinReelsEventResult, spin.Result),
			sdk.NewAttribute(types.SpinReelsEventWinning, spin.Winning),
		),
	)
	return &types.MsgSpinReelsResponse{}, nil

}
