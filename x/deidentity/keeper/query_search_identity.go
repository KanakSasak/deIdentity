package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"

	"deIdentity/x/deidentity/types"
)

// GetNationalID returns the identity from its NationalId
func (k Keeper) SearchByNationalID(ctx context.Context, NationalIDHash string) (val types.Identity, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IdentityKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Identity
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		if val.NationalId == NationalIDHash {
			return val, true
		}
	}

	return val, false
}
