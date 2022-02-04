package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

// GetTaCount get the total number of ta
func (k Keeper) GetTaCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TaCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTaCount set the total number of ta
func (k Keeper) SetTaCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TaCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTa appends a ta in the store with a new id and update the count
func (k Keeper) AppendTa(
	ctx sdk.Context,
	ta types.Ta,
) uint64 {
	// Create the ta
	count := k.GetTaCount(ctx)

	// Set the ID of the appended value
	ta.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TaKey))
	appendedValue := k.cdc.MustMarshal(&ta)
	store.Set(GetTaIDBytes(ta.Id), appendedValue)

	// Update ta count
	k.SetTaCount(ctx, count+1)

	return count
}

// SetTa set a specific ta in the store
func (k Keeper) SetTa(ctx sdk.Context, ta types.Ta) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TaKey))
	b := k.cdc.MustMarshal(&ta)
	store.Set(GetTaIDBytes(ta.Id), b)
}

// GetTa returns a ta from its id
func (k Keeper) GetTa(ctx sdk.Context, id uint64) (val types.Ta, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TaKey))
	b := store.Get(GetTaIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTa removes a ta from the store
func (k Keeper) RemoveTa(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TaKey))
	store.Delete(GetTaIDBytes(id))
}

// GetAllTa returns all ta
func (k Keeper) GetAllTa(ctx sdk.Context) (list []types.Ta) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TaKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Ta
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) CheckCompliance(trustAnchor, token string) (bool, error) {
	return true, nil
}

// GetTaIDBytes returns the byte representation of the ID
func GetTaIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTaIDFromBytes returns ID in uint64 format from a byte array
func GetTaIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
