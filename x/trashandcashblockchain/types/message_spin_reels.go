package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSpinReels = "spin_reels"

var _ sdk.Msg = &MsgSpinReels{}

func NewMsgSpinReels(creator string, amount string) *MsgSpinReels {
	return &MsgSpinReels{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgSpinReels) Route() string {
	return RouterKey
}

func (msg *MsgSpinReels) Type() string {
	return TypeMsgSpinReels
}

func (msg *MsgSpinReels) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSpinReels) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSpinReels) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
