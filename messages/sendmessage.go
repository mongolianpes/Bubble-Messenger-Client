package messages

import (
	"bytes"
	"client/config"
	"client/crypto"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type SendMessageRequest struct {
	SenderLogin    string
	SenderPassword string
	ReceiverLogin  string
	Device         string
	Message        string
}

func SendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key string) string {
	senderLoginEncrypted, err := crypto.StringEncrypt([]byte(senderLogin), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	receiverLoginEncrypted, err := crypto.StringEncrypt([]byte(receiverLogin), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	messageEncrypted, err := crypto.StringEncrypt([]byte(message), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	senderPasswordEncrypted, err := crypto.StringEncrypt([]byte(senderPassword), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, senderLogin)

	requestBody := SendMessageRequest{
		SenderLogin:    senderLoginEncrypted,
		SenderPassword: senderPasswordEncrypted,
		ReceiverLogin:  receiverLoginEncrypted,
		Device:         hashedDeviceInfo,
		Message:        messageEncrypted,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return "$ ошибка создания json"
	}

	req, err := http.NewRequest("POST", config.ForServerReq+config.MainServicePort+"/sendmessage", bytes.NewBuffer(requestJSON))
	if err != nil {
		return "$ ошибка создания запроса"
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "$ ошибка отправки запроса"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "$ Не удается обработать ответ сервера"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		if string(body) == "" {
			return "Сообщение отправлено"
		}
		respDecrypted, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ ошибка расшифровки"
		}
		return respDecrypted
	case http.StatusBadRequest:
		return "$ Проверьте правильность ввода данных"
	case http.StatusForbidden:
		return "$ Не правильный логин или пароль"
	case http.StatusInternalServerError:
		resp, _ := crypto.StringDecrypt(string(body), key)
		return "$ Внутренняя ошибка сервера:" + resp
	default:
		return "$ " + string(body)
	}
}
