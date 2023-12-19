package db

import (
	"context"
	"testing"

	"github.com/NightShop/fasunga-project/util"
	"github.com/stretchr/testify/require"
)

func TestCreateGroup(t *testing.T) {
	groupKey := createRandomGroup(t)

	testQueries.DeleteGroup(context.Background(), groupKey)
}

func TestDeleteGroup(t *testing.T) {
	groupKey := createRandomGroup(t)

	err := testQueries.DeleteGroup(context.Background(), groupKey)
	require.NoError(t, err)
}

func createRandomGroup(t *testing.T) string {
	groupKeyParameter := util.RandomString(10)
	groupKey, err := testQueries.CreateGroup(context.Background(), groupKeyParameter)

	require.NoError(t, err)
	require.NotEmpty(t, groupKey)

	require.Equal(t, groupKey, groupKeyParameter)

	return groupKey
}

func TestGetGroup(t *testing.T) {
	groupKeyParameter := util.RandomString(10)
	groupKey1, err := testQueries.CreateGroup(context.Background(), groupKeyParameter)

	require.NoError(t, err)
	require.NotEmpty(t, groupKey1)

	groupKey2, err := testQueries.GetGroup(context.Background(), groupKeyParameter)
	require.NoError(t, err)
	require.NotEmpty(t, groupKey2)

	require.Equal(t, groupKey1, groupKey2)
}
