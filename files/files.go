package files

import (
	"client/config"
	"client/crypto"
	"client/request"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type SendFileRequest struct {
	Sender   string
	Receiver string
	FileName string
	File     []byte
}

type DelFileRequest struct {
	Filename string
}

func SendFile(login, password, receiverLogin, fileName, deviceInfo, key string, file []byte) string {
	req := SendFileRequest{
		Sender:   login,
		Receiver: receiverLogin,
		FileName: fileName,
		File:     file,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "$ не удалось преобразовать запрос в JSON" + err.Error()
	}

	resp, err := request.SendRequest(login, password, deviceInfo, key, config.ForServerReq+config.MediumSizeDataServicePort+"/", "sendfile", reqBody)
	if err != nil {
		return "$ не удалось отправить реквест " + err.Error()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "$ ошибка чтения ответа"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return "Успешно отправлен"
	case http.StatusBadRequest:
		return "$ Проверьте правильность ввода данных "
	case http.StatusForbidden:
		return "$ Не правильный логин или пароль" + string(body)
	case http.StatusInternalServerError:
		resp, _ := crypto.StringDecrypt(string(body), key)
		return "$ Внутренняя ошибка сервера" + resp
	default:
		return "$ Неизвестная ошибка" + string(body)
	}
}

func GetFile(login, password, fileStoragePath, deviceInfo, key string) []byte {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(config.FilesService + "/" + fileStoragePath)
	if err != nil {
		if os.IsTimeout(err) {
			return []byte("$ Сервер не отвечает")
		}
		return []byte("$ Не удается установить соединение с сервером")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte("$ Не удается обработать ответ сервера")
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return body
	case http.StatusBadRequest:
		return []byte("$ Проверьте правильность ввода данных")
	case http.StatusForbidden:
		return []byte("$ Не правильный логин или пароль")
	case http.StatusInternalServerError:
		resp, _ := crypto.StringDecrypt(string(body), key)
		return []byte("$ Внутренняя ошибка сервера" + resp)
	default:
		return []byte("$ Неизвестная ошибка")
	}
}

func DelFile(login, password, fileName, deviceInfo, key string) string {
	req := DelFileRequest{
		Filename: fileName,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "$ Не удалось преобразовать запрос в JSON" + err.Error()
	}

	resp, err := request.SendRequest(login, password, deviceInfo, key, config.ForServerReq+config.MediumSizeDataServicePort+"/", "delfile", reqBody)
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
		return "Успешно удалено"
	case http.StatusBadRequest:
		return "$ Проверьте правильность ввода данных " + string(body)
	case http.StatusForbidden:
		return "$ Не правильный логин или пароль" + string(body)
	case http.StatusInternalServerError:
		return "$ Внутренняя ошибка сервера" + string(body)
	default:
		return "$ Неизвестная ошибка" + string(body)
	}
}
