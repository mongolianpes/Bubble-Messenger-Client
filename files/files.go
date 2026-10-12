package files

import (
	"bytes"
	"client/config"
	"client/crypto"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type SendFileRequest struct {
	Sender               string
	Receiver             string
	SenderPassword       string
	FileName             string
	Device               string
	KeyForServerDataBase string
	File                 []byte
}

func SendFile(login, password, receiverLogin, fileName, deviceInfo, key string, file []byte) string {
	timeout := 10 * time.Second

	loginEcnrypted, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return "$ ошибка зашифровки данных"
	}
	passwordEcnrypted, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return "$ ошибка зашифровки данных"
	}
	receivedLoginEcnrypted, err := crypto.StringEncrypt([]byte(receiverLogin), key)
	if err != nil {
		return "$ ошибка зашифровки данных"
	}
	fileNameEcnrypted, err := crypto.StringEncrypt([]byte(fileName), key)
	if err != nil {
		return "$ ошибка зашифровки данных"
	}
	fileEcnrypted, err := crypto.StringEncryptByte(file, key)
	if err != nil {
		return "$ ошибка зашифровки данных"
	}
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	requestBody := SendFileRequest{
		Sender:               loginEcnrypted,
		Receiver:             receivedLoginEcnrypted,
		SenderPassword:       passwordEcnrypted,
		FileName:             fileNameEcnrypted,
		Device:               hashedDeviceInfo,
		KeyForServerDataBase: crypto.GenKey(deviceInfo),
		File:                 fileEcnrypted,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return "$ ошибка создания json"
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", config.ForServerReq+config.MediumSizeDataServicePort+"/sendfile", bytes.NewBuffer(requestJSON))
	if err != nil {
		return "$ ошибка создания запроса"
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "$ сервер не отвечает"
		} else {
			return "$ ошибка отправки запроса"
		}
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
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	loginEcnrypted, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	passwordEncrypted, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	fileNameEncrypted, err := crypto.StringEncrypt([]byte(fileName), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	formData := url.Values{}
	formData.Add("login", loginEcnrypted)
	formData.Add("password", passwordEncrypted)
	formData.Add("filename", fileNameEncrypted)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("forserver", crypto.GenKey(deviceInfo))

	resp, err := client.PostForm(config.ForServerReq+config.MediumSizeDataServicePort+"/delfile", formData)
	if err != nil {
		if os.IsTimeout(err) {
			return "$ Сервер не отвечает"
		}
		return "$ Не удается установить соединение с сервером"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
