package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SpinList: []Spin{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in spin
	spinIdMap := make(map[uint64]bool)
	spinCount := gs.GetSpinCount()
	for _, elem := range gs.SpinList {
		if _, ok := spinIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for spin")
		}
		if elem.Id >= spinCount {
			return fmt.Errorf("spin id should be lower or equal than the last id")
		}
		spinIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
