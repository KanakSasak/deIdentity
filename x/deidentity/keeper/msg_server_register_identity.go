package keeper

import (
	"context"

	"deIdentity/x/deidentity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterIdentity(goCtx context.Context, msg *types.MsgRegisterIdentity) (*types.MsgRegisterIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterIdentityResponse{}, nil
}
