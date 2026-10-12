package auth

import (
	"fmt"
	"regexp"
	"strings"

	"client/crypto"
)

func Reg(login, name, password, deviceInfo string) string {
	if !isValidStr(login) && !isValidStr(name) && !isValidStr(password) {
		return "$ Используйте только символы: a - z, A - Z или 0 - 9"
	}

	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	result := exchangeKeys(hashedDeviceInfo, true)

	if result.Error != nil {
		return "$ Не удалось совершить обмен ключами: " + result.Error.Error()
	}

	secretKey := fmt.Sprintf("%x", result.SharedSecret)
	res := regRequest(login, name, password, secretKey, hashedDeviceInfo, deviceInfo)
	if !strings.Contains(res, "$") {
		return secretKey
	} else {
		return res
	}
}

func Auth(login, password, deviceInfo string) string {
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)
	result := exchangeKeys(hashedDeviceInfo, false)

	if result.Error != nil {
		return result.Error.Error()
	}

	secretKey := fmt.Sprintf("%x", result.SharedSecret)
	res := authRequest(login, password, secretKey, hashedDeviceInfo, deviceInfo)
	if strings.Contains(res, "$") {
		return res
	} else {

		return secretKey + "\n" + res
	}
}

var re *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

func isValidStr(str string) bool {
	if len(str) > 20 {
		return false
	}

	return re.MatchString(str)
}
