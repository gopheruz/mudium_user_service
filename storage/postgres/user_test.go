package postgres_test

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	_ "github.com/lib/pq"
	"github.com/nurmuhammaddeveloper/mudium_user_service/storage/repo"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *repo.User {
	u, err := strg.User().Create(&repo.User{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  faker.Password(),
		Type:      repo.UserTypeUser,
	})
	require.NoError(t, err)
	require.NotEmpty(t, u)

	return u
}

func deleteUser(t *testing.T){
	u, err := strg.User().Create(&repo.User{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  faker.Password(),
		Type:      repo.UserTypeUser,
	})
	require.NoError(t, err)
	err = strg.User().Delete(u.ID)
	require.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	createUser(t)
}

func TestDeleteUser(t *testing.T) {
	deleteUser(t)
}