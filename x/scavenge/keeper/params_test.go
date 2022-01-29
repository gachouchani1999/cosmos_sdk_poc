package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmonaut/scavenge/testutil/keeper"
	"github.com/cosmonaut/scavenge/x/scavenge/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ScavengeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
