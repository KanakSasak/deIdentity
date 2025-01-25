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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DeidentityKeeper(t)
	deidentity.InitGenesis(ctx, k, genesisState)
	got := deidentity.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
