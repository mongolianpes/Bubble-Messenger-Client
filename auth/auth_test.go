package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

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

func TestRegAndAuth(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_, _, _, _, err := RegTestUser()
		require.NoError(t, err)
	})

	t.Run("duplicate login", func(t *testing.T) {
		login, name, password, device, err := RegTestUser()
		require.NoError(t, err)

		reg2 := Reg(login, name, password, device)
		require.Contains(t, reg2, "$")
	})
}

func TestAuth(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		login, name, password, device, err := RegTestUser()
		require.NoError(t, err)

		auth := Auth(login, password, device)
		require.NotContains(t, auth, "$")
		authSplitted := strings.Split(auth, "\n")
		require.Equal(t, authSplitted[1], name)
	})

	t.Run("non-existent user", func(t *testing.T) {
		fmt.Println("В случае возникновение неудачного теста: возможно сгенерированные рандомный логин уже существует")
		fmt.Println("Стоит очистить БД и провести тест еще несколько раз")
		login := generateTestRandValidStr()
		password := generateTestRandValidStr()
		device := generateTestRandValidStr()

		auth := Auth(login, password, device)
		require.Contains(t, auth, "$")
	})
}

func RegTestUser() (login, name, password, device string, err error) {
	login = generateTestRandValidStr()
	name = generateTestRandValidStr()
	password = generateTestRandValidStr()
	device = generateTestRandValidStr()

	result := Reg(login, name, password, device)
	if strings.Contains(result, "$") {
		err = errors.New(result)
	}

	return
}
