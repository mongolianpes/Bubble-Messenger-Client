package files

import (
	"errors"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"client/auth"

	"github.com/stretchr/testify/require"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateTestRandValidStr() string {
	charset := "abcABC123"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RegAndAuthTestUser() (login, name, password, device, key string, err error) {
	login = generateTestRandValidStr()
	name = generateTestRandValidStr()
	password = generateTestRandValidStr()
	device = generateTestRandValidStr()

	reg := auth.Reg(login, name, password, device)
	if strings.Contains(reg, "$") {
		err = errors.New(reg)
		return
	}

	auth := auth.Auth(login, password, device)
	if strings.Contains(auth, "$") {
		err = errors.New(auth)
		return
	}

	key = strings.Split(auth, "\n")[0]

	return
}

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
