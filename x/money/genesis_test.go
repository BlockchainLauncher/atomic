package money_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "money/testutil/keeper"
	"money/testutil/nullify"
	"money/x/money"
	"money/x/money/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MoneyKeeper(t)
	money.InitGenesis(ctx, *k, genesisState)
	got := money.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
