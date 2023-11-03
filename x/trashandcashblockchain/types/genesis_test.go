package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"trashandcash-blockchain/x/trashandcashblockchain/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				SpinList: []types.Spin{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				SpinCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated spin",
			genState: &types.GenesisState{
				SpinList: []types.Spin{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid spin count",
			genState: &types.GenesisState{
				SpinList: []types.Spin{
					{
						Id: 1,
					},
				},
				SpinCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
