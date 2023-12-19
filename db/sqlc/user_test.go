package db

import (
	"context"
	"testing"

	"github.com/NightShop/fasunga-project/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	groupKey := createRandomGroup(t)

	arg := CreateUserParams{
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomString(20),
		GroupKey:       groupKey,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.GroupKey, user.GroupKey)

	require.NotZero(t, user.Email)
}

func TestDeleteUser(t *testing.T) {
	groupKey := createRandomGroup(t)

	arg := CreateUserParams{
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomString(20),
		GroupKey:       groupKey,
	}

	user, _ := testQueries.CreateUser(context.Background(), arg)

	err := testQueries.DeleteUser(context.Background(), user.Email)
	require.NoError(t, err)

	user2, _ := testQueries.GetUser(context.Background(), user.Email)
	require.Empty(t, user2)
}

func TestUpdateUser(t *testing.T) {
	groupKey := createRandomGroup(t)
	groupKey2 := createRandomGroup(t)

	arg := CreateUserParams{
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomString(20),
		GroupKey:       groupKey,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	updateUserParams := UpdateUserParams{
		Email:    user.Email,
		GroupKey: groupKey2,
	}

	user2, err := testQueries.UpdateUser(context.Background(), updateUserParams)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, updateUserParams.GroupKey, user2.GroupKey)
}

func TestGetUser(t *testing.T) {
	groupKey := createRandomGroup(t)

	arg := CreateUserParams{
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomString(20),
		GroupKey:       groupKey,
	}

	user1, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	user2, err := testQueries.GetUser(context.Background(), arg.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.GroupKey, user2.GroupKey)

}
