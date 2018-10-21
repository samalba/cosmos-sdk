package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// keeper of the stake store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        sdk.Codec
	paramSpace params.Subspace
	sk         StakeKeeper
	fck        FeeCollectionKeeper
}

func NewKeeper(cdc sdk.Codec, key sdk.StoreKey,
	paramSpace params.Subspace, sk StakeKeeper, fck FeeCollectionKeeper) Keeper {

	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		paramSpace: paramSpace.WithTypeTable(ParamTypeTable()),
		sk:         sk,
		fck:        fck,
	}
	return keeper
}

//____________________________________________________________________
// Keys

var (
	minterKey = []byte{0x00} // the one key to use for the keeper store

	// params store for inflation params
	ParamStoreKeyParams = []byte("params")
)

// ParamTable for stake module
func ParamTypeTable() params.TypeTable {
	return params.NewTypeTable(
		ParamStoreKeyParams, Params{},
	)
}

const (
	// default paramspace for params keeper
	DefaultParamspace = "mint"
)

//______________________________________________________________________

// get the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(minterKey)
	if b == nil {
		panic("Stored fee pool should not have been nil")
	}
	k.cdc.MustUnmarshalBinary(b, &minter)
	return
}

// set the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinary(minter)
	store.Set(minterKey, b)
}

//______________________________________________________________________

// get inflation params from the global param store
func (k Keeper) GetParams(ctx sdk.Context) Params {
	var params Params
	k.paramSpace.Get(ctx, ParamStoreKeyParams, &params)
	return params
}

// set inflation params from the global param store
func (k Keeper) SetParams(ctx sdk.Context, params Params) {
	k.paramSpace.Set(ctx, ParamStoreKeyParams, &params)
}