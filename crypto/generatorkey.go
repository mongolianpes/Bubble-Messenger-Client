package crypto

import (
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

func GenKey(inputString string) string {
	splittedDeviceInfo := strings.Split(inputString, ".")

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
	result := hex.EncodeToString(sum)

	return result[:64]
}
