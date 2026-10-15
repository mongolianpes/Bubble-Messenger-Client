package messages

import (
	"client/crypto"
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"log"
)

const (
	designationOfInitDialog   = "i"
	designationOfAnswerDialog = "a"
)

func initiatorKeyGeneration() (string, string, error) {
	curve := ecdh.P256()

	privKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}
	pubKey := privKey.PublicKey()

	// сериализуем приватный ключ в base64
	privKeyBase64 := base64.StdEncoding.EncodeToString(privKey.Bytes())
	// сериализуем публичный ключ в base64
	pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKey.Bytes())

	return privKeyBase64, pubKeyBase64, nil
}

func InitDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, key string) string {
	privKeyBase64, pubKeyBase64, err := initiatorKeyGeneration()
	if err != nil {
		return "Произошла ошибка при генерации ключей"
	}

	resp := sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, designationOfInitDialog+"\\"+pubKeyBase64, key)
	if resp == "Сообщение отправлено" {
		return privKeyBase64
	}
	return "$ " + resp
}

func answerKeyGeneration(user1PubBase64 string) (string, string) {
	curve := ecdh.P256()

	// генерируем ключи юзера2
	privKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("User2: ошибка генерации ключа: %v", err)
	}
	pubKey := privKey.PublicKey()

	// декодируем публичный ключ юзера1
	user1PubBytes, err := base64.StdEncoding.DecodeString(user1PubBase64)
	if err != nil {
		log.Fatalf("User2: ошибка декодирования публичного ключа юзера1: %v", err)
	}
	user1PubKey, err := curve.NewPublicKey(user1PubBytes)
	if err != nil {
		log.Fatalf("User2: ошибка восстановления публичного ключа юзера1: %v", err)
	}

	// вычисляем общий секрет
	sharedSecret, err := privKey.ECDH(user1PubKey)
	if err != nil {
		log.Fatalf("User2: ошибка вычисления общего секрета: %v", err)
	}

	// сериализуем публичный ключ юзера2
	pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKey.Bytes())
	return pubKeyBase64, hex.EncodeToString(sharedSecret)
}

func AnswerInitDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, pubKeyBase64, key string) string {
	pubKeyBase64ForSend, sharedSecret := answerKeyGeneration(pubKeyBase64)

	resp := sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, designationOfAnswerDialog+"\\"+pubKeyBase64ForSend, key)
	if resp == "Сообщение отправлено" {
		return sharedSecret + "\\" + crypto.HashLogin(receiverLogin, sharedSecret)
	}
	return "$ " + resp

}

func EndInitDialog(privBase64, user2PubBase64, loginCommunicating string) string {
	curve := ecdh.P256()

	// восстанавливаем приватный ключ юзера1
	privBytes, err := base64.StdEncoding.DecodeString(privBase64)
	if err != nil {
		log.Fatalf("User1: ошибка декодирования приватного ключа: %v", err)
	}
	privKey, err := curve.NewPrivateKey(privBytes)
	if err != nil {
		log.Fatalf("User1: ошибка восстановления приватного ключа: %v", err)
	}

	// восстанавливаем публичный ключ юзера2
	pubBytes, err := base64.StdEncoding.DecodeString(user2PubBase64)
	if err != nil {
		log.Fatalf("User1: ошибка декодирования публичного ключа юзера2: %v", err)
	}
	pubKey, err := curve.NewPublicKey(pubBytes)
	if err != nil {
		log.Fatalf("User1: ошибка восстановления публичного ключа юзера2: %v", err)
	}

	// вычисляем общий секрет
	sharedSecret, err := privKey.ECDH(pubKey)
	if err != nil {
		log.Fatalf("User1: ошибка вычисления общего секрета: %v", err)
	}
	return hex.EncodeToString(sharedSecret) + "\\" + crypto.HashLogin(loginCommunicating, hex.EncodeToString(sharedSecret))
}
