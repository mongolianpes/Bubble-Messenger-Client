package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"client/config"
	"client/crypto"
)

type RegRequest struct {
	Login    string
	Name     string
	Password string
	Device   string
}

type AuthRequest struct {
	Login    string
	Password string
	Device   string
}

func regRequest(login, name, password, key, hashedDeviceInfo, deviceInfo string) string {
	loginEncypted, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return "$ неизвестная ошибка"
	}
	nameEncrypted, err := crypto.StringEncrypt([]byte(name), key)
	if err != nil {
		return "$ неизвестная ошибка"
	}
	passwordEncrypted, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return "$ неизвестная ошибка"
	}

	requestBody := RegRequest{
		Login:    loginEncypted,
		Name:     nameEncrypted,
		Password: passwordEncrypted,
		Device:   hashedDeviceInfo,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return "$ ошибка создания json"
	}

	req, err := http.NewRequest("POST", config.ForServerReq+config.MainServicePort+"/reg", bytes.NewBuffer(requestJSON))
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
		return "Не удалось прочитать ответ от сервера"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return ""
	case http.StatusBadRequest:
		return "$ " + string(body)
	case http.StatusInternalServerError:
		return "$ " + string(body)
	default:
		decryptedResp, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ попробуйте зарегистрироваться заново"
		}
		return "$ " + decryptedResp
	}
}

func authRequest(login, password, key, hashedDeviceInfo, deviceInfo string) string {
	loginEncypted, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return "$ неизвестная ошибка"
	}
	passwordEncrypted, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return "$ неизвестная ошибка"
	}

	requestBody := AuthRequest{
		Login:    loginEncypted,
		Password: passwordEncrypted,
		Device:   hashedDeviceInfo,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return "$ ошибка создания json"
	}

	req, err := http.NewRequest("POST", config.ForServerReq+config.MainServicePort+"/auth", bytes.NewBuffer(requestJSON))
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
		return "Не удается обработать ответ сервера"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		decryptedResp, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ попробуйте авторизоваться заново"
		}
		return decryptedResp
	case http.StatusBadRequest:
		return "$ " + string(body)
	case http.StatusInternalServerError:
		return "$ " + string(body)
	default:
		return ""
	}
}
