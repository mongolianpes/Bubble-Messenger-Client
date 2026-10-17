package messages

import (
	"client/config"
	"client/crypto"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func SendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

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

	formData := url.Values{}
	formData.Add("senderlogin", senderLoginEncrypted)
	formData.Add("senderpassword", senderPasswordEncrypted)
	formData.Add("receiverlogin", receiverLoginEncrypted)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("message", messageEncrypted)
	formData.Add("forserver", crypto.GenKey(deviceInfo))
	// if !(message[0] == designationOfAnswerDialog[0] || message[0] == designationOfInitDialog[0]) {
	// 	formData.Add("senderLoginHashed", crypto.HashLogin(senderLogin, dialogKey))
	// }

	resp, err := client.PostForm(config.ForServerReq+config.MainServicePort+"/sendmessage", formData)
	if err != nil {
		if os.IsTimeout(err) {
			return "$ Сервер не отвечает"
		}
		return "$ Не удается установить соединение с сервером"
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
