package messages

import (
	"client/config"
	"client/crypto"
	"client/request"
	"encoding/json"
	"io"
	"net/http"
)

type SendMessageRequest struct {
	SenderLogin   string
	ReceiverLogin string
	Message       string
}

func SendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key string) string {
	req := SendMessageRequest{
		SenderLogin:   senderLogin,
		ReceiverLogin: receiverLogin,
		Message:       message,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "$ не удалось преобразовать запрос в JSON" + err.Error()
	}

	resp, err := request.SendRequest(senderLogin, senderPassword, deviceInfo, key, config.ForServerReq+config.MainServicePort+"/", "sendmessage", reqBody)
	if err != nil {
		return "$ не удалось отправить реквест " + err.Error()
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
