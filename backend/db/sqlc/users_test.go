package db

import (
	"context"
	"ourplaylist/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	hashedPassword, err := util.HashPassword("avbbsd")
	require.NoError(t, err)
	arg := CreateUserParams{
		Username: "1313131",
		Password: hashedPassword,
		Mail:     "cm@gmail.com",
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Mail, user.Mail)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.ID)
}
