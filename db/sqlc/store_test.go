package db

import (
	"context"
	"testing"

	"github.com/NightShop/fasunga-project/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUserTx(t *testing.T) {
	store := NewStore(testDB)
	groupKey := util.RandomString(10)

	arg := CreateUserTxParams{
		Email:    util.RandomEmail(),
		GroupKey: groupKey,
	}

	// Group shouldn't exists yet
	createdGroup, err := testQueries.GetGroup(context.Background(), arg.GroupKey)
	require.Error(t, err)
	require.Empty(t, createdGroup)

	_, err = store.CreateUserTx(context.Background(), arg)
	require.NoError(t, err)

	// Group should be created
	createdGroup, err = testQueries.GetGroup(context.Background(), arg.GroupKey)
	require.NoError(t, err)
	require.NotEmpty(t, createdGroup)
	require.Equal(t, arg.GroupKey, createdGroup)

	// New user with already existing group should be created
	_, err = store.CreateUserTx(context.Background(), CreateUserTxParams{
		Email:    util.RandomEmail(),
		GroupKey: groupKey,
	})
	require.NoError(t, err)

}
