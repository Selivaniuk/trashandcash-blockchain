package keeper

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"trashandcash-blockchain/x/trashandcashblockchain/types"
)

// GetSpinCount get the total number of spin
func (k Keeper) GetSpinCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SpinCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSpinCount set the total number of spin
func (k Keeper) SetSpinCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SpinCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSpin appends a spin in the store with a new id and update the count
func (k Keeper) AppendSpin(
	ctx sdk.Context,
	spin types.Spin,
) uint64 {
	// Create the spin
	count := k.GetSpinCount(ctx)

	// Set the ID of the appended value
	spin.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpinKey))
	appendedValue := k.cdc.MustMarshal(&spin)
	store.Set(GetSpinIDBytes(spin.Id), appendedValue)

	// Update spin count
	k.SetSpinCount(ctx, count+1)

	return count
}

// SetSpin set a specific spin in the store
func (k Keeper) SetSpin(ctx sdk.Context, spin types.Spin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpinKey))
	b := k.cdc.MustMarshal(&spin)
	store.Set(GetSpinIDBytes(spin.Id), b)
}

// GetSpin returns a spin from its id
func (k Keeper) GetSpin(ctx sdk.Context, id uint64) (val types.Spin, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpinKey))
	b := store.Get(GetSpinIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSpin removes a spin from the store
func (k Keeper) RemoveSpin(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpinKey))
	store.Delete(GetSpinIDBytes(id))
}

// GetAllSpin returns all spin
func (k Keeper) GetAllSpin(ctx sdk.Context) (list []types.Spin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpinKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Spin
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSpinIDBytes returns the byte representation of the ID
func GetSpinIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSpinIDFromBytes returns ID in uint64 format from a byte array
func GetSpinIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func generateRandomNumber(ratio float64) (result string, err error) {
	base := 8.0
	exponent := 3.0
	max := (math.Pow(base, exponent)) - 1

	randomNumber := math.Floor(ratio * max)
	octal := fmt.Sprintf("%03o", int(randomNumber))

	return octal, nil
}

func (k Keeper) GenerateResult(ctx sdk.Context) (result string, err error) {
	ratio, err := GetRatioByTxBytes(ctx.TxBytes())
	if err != nil {
		return "", err
	}

	randomNumber, err := generateRandomNumber(ratio)
	if err != nil {
		return "", err
	}

	fmt.Println("Random Number: ", randomNumber)

	return randomNumber, nil
}
