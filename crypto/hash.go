package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"strings"
)

const (
	saltSize = 32
	iter     = 5
)

var mainsalt = []byte("jhgogofzfhadfhjikhbnuoik]")
var mainsalt2 = []byte("damgagagnagnasfnoignasjofsngo")

func HashLogin(loginToHash, key string) string {
	return "/" + hashString512(key) + hashString512(loginToHash)[:10]
}

func HashDeviceInfo(stringToHash, userLogin string) string {
	result := hashString512(stringToHash)

	splittedDeviceInfo := strings.Split(stringToHash, ".")
	result += hashString512(splittedDeviceInfo[0])

	result += hashString512(userLogin)[:80]

	result += hashString256(stringToHash)

	// fmt.Println(result[:172] + result[252:]) // devices
	// fmt.Println(result[:252])                //ids

	return result
}

func hashString512(stringToHash string) string {
	splittedDeviceInfo := strings.Split(stringToHash, ".")

	h := sha512.New()
	h.Write(mainsalt)
	for numInfo := range splittedDeviceInfo {
		h.Write([]byte(splittedDeviceInfo[numInfo]))
	}
	sum := h.Sum(nil)

	for i := 1; i < iter; i++ {
		h2 := sha512.New()
		h2.Write(sum)
		sum = h2.Sum(nil)
	}
	result := base64.RawURLEncoding.EncodeToString(sum)

	return result
}

func hashString256(stringToHash string) string {
	splittedDeviceInfo := strings.Split(stringToHash, ".")

	h := sha256.New()
	h.Write(mainsalt2)
	for numInfo := range splittedDeviceInfo {
		h.Write([]byte(splittedDeviceInfo[numInfo]))
	}
	sum2 := h.Sum(nil)

	for i := 1; i < iter; i++ {
		h4 := sha256.New()
		h4.Write(sum2)
		sum2 = h4.Sum(nil)
	}
	result := base64.RawURLEncoding.EncodeToString(sum2)

	return result
}
