package deidentity_test

import (
	"testing"

	keepertest "deIdentity/testutil/keeper"
	"deIdentity/testutil/nullify"
	deidentity "deIdentity/x/deidentity/module"
	"deIdentity/x/deidentity/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		IdentityList: []types.Identity{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		IdentityCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DeidentityKeeper(t)
	deidentity.InitGenesis(ctx, k, genesisState)
	got := deidentity.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.IdentityList, got.IdentityList)
	require.Equal(t, genesisState.IdentityCount, got.IdentityCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
