package messages

import (
	"client/auth"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
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
	fmt.Println("В случае возникновение неудачного теста: возможно сгенерированные рандомный логин уже существует")
	fmt.Println("Стоит очистить БД и провести тест еще несколько раз")
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
