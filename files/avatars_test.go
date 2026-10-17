package files

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetAndGetAvatar(t *testing.T) {
	avatarBytes, err := os.ReadFile("F:/GolangProjects/BubleMessenger/client/192.webp")
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		login, _, password, device, key, err := RegAndAuthTestUser()
		require.NoError(t, err)

		setAvatar := SetAvatar(login, password, device, key, avatarBytes)
		require.NotContains(t, setAvatar, "$")

		gettedAvatar := GetAvatar(login, password, login, device, key)
		require.Equal(t, gettedAvatar, avatarBytes)
	})

	t.Run("not exist user login", func(t *testing.T) {
		_, _, password, device, key, err := RegAndAuthTestUser()
		require.NoError(t, err)

		randStr := generateTestRandValidStr()

		setAvatar := SetAvatar(randStr, password, device, key, avatarBytes)
		require.Contains(t, setAvatar, "$")
	})
}
