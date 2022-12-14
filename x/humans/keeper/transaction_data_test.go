package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/humansdotai/humans/testutil/keeper"
	"github.com/humansdotai/humans/testutil/nullify"
	"github.com/humansdotai/humans/x/humans/keeper"
	"github.com/humansdotai/humans/x/humans/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTransactionData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TransactionData {
	items := make([]types.TransactionData, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTransactionData(ctx, items[i])
	}
	return items
}

func TestTransactionDataGet(t *testing.T) {
	keeper, ctx := keepertest.HumansKeeper(t)
	items := createNTransactionData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTransactionData(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTransactionDataRemove(t *testing.T) {
	keeper, ctx := keepertest.HumansKeeper(t)
	items := createNTransactionData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTransactionData(ctx,
			item.Index,
		)
		_, found := keeper.GetTransactionData(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTransactionDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.HumansKeeper(t)
	items := createNTransactionData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTransactionData(ctx)),
	)
}
