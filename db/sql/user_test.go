package sql

import (
	"context"
	"testing"

	"beebeewijaya.com/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	randomPassword := util.RandomString(8)
	hashed, err := util.HashPassword(randomPassword)
	require.NoError(t, err)
	require.NotEmpty(t, hashed)

	args := CreateUserArgs{
		Email:    util.RandomEmail(),
		Password: hashed,
		Fullname: util.RandomString(12),
	}

	user, err := testQuery.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Email, args.Email)
	require.Equal(t, user.Password, hashed)
	require.Equal(t, user.Fullname, args.Fullname)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
