package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"client/config"
	"client/crypto"
)

func regRequest(login, name, password, key, hashedDeviceInfo, deviceInfo string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	keyForServer, err := crypto.StringEncrypt([]byte(crypto.GenKey(deviceInfo)), key)
	if err != nil {
		return "$ неизвестная ошибка"
	}
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

	formData := url.Values{}
	formData.Add("login", loginEncypted)
	formData.Add("name", nameEncrypted)
	formData.Add("password", passwordEncrypted)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("forserver", keyForServer)

	resp, err := client.PostForm(config.ForServerReq+config.MainServicePort+"/reg", formData)
	if err != nil {
		if os.IsTimeout(err) {
			return "Сервер не отвечает"
		}
		return "Сервер не отвечает по какой то херне"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Не удалось прочитать ответ от сервера"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return ""
	case http.StatusBadRequest:
		if len(string(body)) > 20 {
			decryptedResp, err := crypto.StringDecrypt(string(body), key)
			if err != nil {
				return "$ попробуйте зарегистрироваться заново"
			}
			return "$ " + decryptedResp
		}
		return "$ " + string(body)
		// return "Проверьте правильность ввода данных"
	case http.StatusInternalServerError:
		decryptedResp, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ попробуйте зарегистрироваться заново"
		}
		return "$ " + decryptedResp
	default:
		decryptedResp, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ попробуйте зарегистрироваться заново"
		}
		return "$ " + decryptedResp
	}
}

func authRequest(login, password, key, hashedDeviceInfo, deviceInfo string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	keyForServer, err := crypto.StringEncrypt([]byte(crypto.GenKey(deviceInfo)), key)
	if err != nil {
		return "$ неизвестная ошибка"
	}
	loginEncypted, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return "$ неизвестная ошибка"
	}
	passwordEncrypted, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return "$ неизвестная ошибка"
	}

	formData := url.Values{}
	formData.Add("login", loginEncypted)
	formData.Add("password", passwordEncrypted)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("forserver", keyForServer)

	resp, err := client.PostForm(config.ForServerReq+config.MainServicePort+"/auth", formData)
	if err != nil {
		if os.IsTimeout(err) {
			return "Сервер не отвечает"
		}
		return "Не удается установить соединение с сервером: " + err.Error()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
		if string(body) == "Did not receive all server data" || string(body) == "This request is not expected for you" {
			return "$" + string(body)
		}
		decryptedResp, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ попробуйте авторизоваться заново"
		}
		return "$ " + decryptedResp
	// 	return "Проверьте правильность ввода данных"
	// case http.StatusForbidden:
	// 	return "Не правильный логин или пароль"
	case http.StatusInternalServerError:
		fmt.Println(string(body))
		decryptedResp, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ попробуйте авторизоваться заново"
		}
		return "$ " + decryptedResp
	// 	return "Внутреняя ошибка сервера"
	default:
		return ""
	}
}
