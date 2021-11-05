package models

import (
	"context"
	"database/sql"
	"testing"
	"time"
	"user-service/util"

	"github.com/stretchr/testify/require"
)

var user = new(User)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParam{
		Username: util.RandomString(5),
		Password: util.RandomString(5),
		Email:    util.RandomEmail(),
	}

	result, err := user.Create(context.Background(), arg)

	// Check if method executed correctly.
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, arg.Username, result.Username)
	require.Equal(t, arg.Password, result.Password)
	require.Equal(t, arg.Email, result.Email)

	require.NotZero(t, result.ID)
	require.NotZero(t, result.CreatedAt)

	return result
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := user.GetByID(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestListUsers(t *testing.T) {

	// Create a list of users in database.
	var createdUsers [10]User
	for i := 0; i < 10; i++ {
		createdUsers[i] = createRandomUser(t)
	}

	arg := ListUserParam{
		Limit:  10,
		Offset: 0,
	}

	// Retrieve list of users.
	users, err := user.GetAll(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, u := range users {
		require.NotEmpty(t, u)
	}
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParam{
		Username: util.RandomString(5),
		Password: util.RandomString(5),
		Email:    util.RandomEmail(),
	}

	user2, err := user.Update(context.Background(), arg, user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := user.Delete(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := user.GetByID(context.Background(), user1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
