package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "deIdentity/testutil/keeper"
	"deIdentity/x/deidentity/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DeidentityKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
