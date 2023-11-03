package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"trashandcash-blockchain/x/trashandcashblockchain/types"
)

func (k Keeper) Spins(goCtx context.Context, req *types.QuerySpinsRequest) (*types.QuerySpinsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var spins []types.Spin
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	spinStore := prefix.NewStore(store, types.KeyPrefix(types.SpinKey))

	iterator := spinStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var spin types.Spin
		if err := k.cdc.Unmarshal(iterator.Value(), &spin); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if len(req.Address) > 0 {
			if spin.Address == req.Address {
				spins = append(spins, spin)
			}
		} else {
			spins = append(spins, spin)
		}
	}
	return &types.QuerySpinsResponse{Spin: spins}, nil
	// var pageRes *query.PageResponse
	// var err error

	// switch {
	// case len(req.Address) > 0:
	// 	pageRes, err = query.Paginate(spinStore, req.Pagination, func(key []byte, value []byte) error {
	// 		var spin types.Spin
	// 		if err := k.cdc.Unmarshal(value, &spin); err != nil {
	// 			return err
	// 		}

	// 		has := spin.Address == req.Address
	// 		if has {
	// 			spins = append(spins, spin)
	// 		}
	// 		return nil
	// 	})
	// default:
	// 	pageRes, err = query.Paginate(spinStore, req.Pagination, func(key []byte, value []byte) error {
	// 		var spin types.Spin
	// 		if err := k.cdc.Unmarshal(value, &spin); err != nil {
	// 			return err
	// 		}

	// 		spins = append(spins, spin)
	// 		return nil
	// 	})
	// }

	// if err != nil {
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }

	// return &types.QuerySpinsResponse{Spin: spins, Pagination: pageRes}, nil
}

func (k Keeper) Spin(goCtx context.Context, req *types.QueryGetSpinRequest) (*types.QueryGetSpinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	spin, found := k.GetSpin(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSpinResponse{Spin: spin}, nil
}
