package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"entexample/pkg/ent"
	"entexample/pkg/ent/fixture"
	"entexample/pkg/ent/test"
)

func TestUser_Get(t *testing.T) {
	client := test.NewEntClient(t)
	defer client.Close()
	repository := user{client: client}

	expectedUser := fixture.NewUser(client, fixture.UserWithName("my user example"))

	t.Run("success", func(t *testing.T) {
		result, err := repository.Get(context.Background(), expectedUser.ID)
		require.NoError(t, err)
		require.Equal(t, expectedUser.ID, result.ID)
		require.Equal(t, expectedUser.Name, result.Name)
	})

	t.Run("not found", func(t *testing.T) {
		result, err := repository.Get(context.Background(), 404)
		require.Nil(t, result)
		require.True(t, ent.IsNotFound(err))
	})

	t.Run("error", func(t *testing.T) {
		errorClient := test.NewEntClient(t)
		require.NoError(t, errorClient.Close())

		result, err := repository.Get(context.Background(), 404)
		require.Nil(t, result)
		require.Error(t, err)
	})
}
