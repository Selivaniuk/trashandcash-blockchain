package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBankReplenishment{}, "trashandcashblockchain/BankReplenishment", nil)
	cdc.RegisterConcrete(&MsgWithdrawBank{}, "trashandcashblockchain/WithdrawBank", nil)
	cdc.RegisterConcrete(&MsgSpinReels{}, "trashandcashblockchain/SpinReels", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBankReplenishment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawBank{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSpinReels{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
