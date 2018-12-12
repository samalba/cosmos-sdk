//nolint
package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Expose the hooks if present
func (k Keeper) PostValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) {
	if k.hooks != nil {
		k.hooks.PostValidatorCreated(ctx, valAddr)
	}
}

func (k Keeper) PreValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) {
	if k.hooks != nil {
		k.hooks.PreValidatorModified(ctx, valAddr)
	}
}

func (k Keeper) PostValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	if k.hooks != nil {
		k.hooks.PostValidatorRemoved(ctx, consAddr, valAddr)
	}
}

func (k Keeper) PostValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	if k.hooks != nil {
		k.hooks.PostValidatorBonded(ctx, consAddr, valAddr)
	}
}

func (k Keeper) PostValidatorPowerDidChange(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	if k.hooks != nil {
		k.hooks.PostValidatorPowerDidChange(ctx, consAddr, valAddr)
	}
}

func (k Keeper) PostValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	if k.hooks != nil {
		k.hooks.PostValidatorBeginUnbonding(ctx, consAddr, valAddr)
	}
}

func (k Keeper) PreDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	if k.hooks != nil {
		k.hooks.PreDelegationCreated(ctx, delAddr, valAddr)
	}
}

func (k Keeper) PreDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	if k.hooks != nil {
		k.hooks.PreDelegationSharesModified(ctx, delAddr, valAddr)
	}
}

func (k Keeper) PreDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	if k.hooks != nil {
		k.hooks.PreDelegationRemoved(ctx, delAddr, valAddr)
	}
}
