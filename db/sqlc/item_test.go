package db

import (
	"context"
	"testing"

	"github.com/NightShop/fasunga-project/util"
	"github.com/stretchr/testify/require"
)

func TestCreateItem(t *testing.T) {
	groupKey := createRandomGroup(t)

	arg := CreateItemParams{
		UserEmail:   util.RandomEmail(),
		Description: util.RandomString(10),
		GroupKey:    groupKey,
	}

	item, err := testQueries.CreateItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.UserEmail, item.UserEmail)
	require.Equal(t, arg.Description, item.Description)
	require.Equal(t, arg.GroupKey, item.GroupKey)
	require.Equal(t, false, item.Checked)

	require.NotZero(t, item.ID)
}

func TestDeleteItem(t *testing.T) {
	groupKey := createRandomGroup(t)

	arg := CreateItemParams{
		UserEmail:   util.RandomEmail(),
		Description: util.RandomString(10),
		GroupKey:    groupKey,
	}

	item, _ := testQueries.CreateItem(context.Background(), arg)

	err := testQueries.DeleteItem(context.Background(), item.ID)
	require.NoError(t, err)

	items, _ := testQueries.ListItems(context.Background(), groupKey)
	require.Empty(t, items)
}

func TestListItems(t *testing.T) {
	groupKey := createRandomGroup(t)

	for i := 0; i < 3; i++ {
		testQueries.CreateItem(context.Background(), CreateItemParams{
			UserEmail:   util.RandomEmail(),
			Description: util.RandomString(5),
			GroupKey:    groupKey,
		})
	}

	for i := 0; i < 3; i++ {
		testQueries.CreateItem(context.Background(), CreateItemParams{
			UserEmail:   util.RandomEmail(),
			Description: util.RandomString(5),
			GroupKey:    util.RandomString(10),
		})
	}

	items, err := testQueries.ListItems(context.Background(), groupKey)

	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 3)

	for _, item := range items {
		require.NotEmpty(t, item)
		require.Equal(t, groupKey, item.GroupKey)
	}
}

func TestUpdateItem(t *testing.T) {
	groupKey := createRandomGroup(t)
	arg := CreateItemParams{
		UserEmail:   util.RandomEmail(),
		Description: util.RandomString(10),
		GroupKey:    groupKey,
	}

	item1, err := testQueries.CreateItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item1)

	updateItemParams := UpdateItemParams{
		ID:      item1.ID,
		Checked: true,
	}

	item2, err := testQueries.UpdateItem(context.Background(), updateItemParams)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.UserEmail, item2.UserEmail)
	require.Equal(t, item1.Description, item2.Description)
	require.Equal(t, item1.GroupKey, item2.GroupKey)
	require.Equal(t, updateItemParams.Checked, item2.Checked)
}
