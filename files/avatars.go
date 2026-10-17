package files

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"client/config"
	"client/crypto"
)

type Avatar struct {
	Avatar               []byte `json:"avatar"`
	Login                string `json:"login"`
	Password             string `json:"password"`
	DeviceInfo           string `json:"deviceinfo"`
	KeyForServerDataBase string `json:"forserver"`
}

func SetAvatar(login, password, deviceInfo, key string, avatar []byte) string {
	loginEcnrypted, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	passwordEncrypted, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	avatarEncrypted, err := crypto.StringEncryptByte(avatar, key)
	if err != nil {
		return "$ ошибка зашифровки 1"
	}
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	payload := Avatar{
		Avatar:               avatarEncrypted,
		Login:                loginEcnrypted,
		Password:             passwordEncrypted,
		DeviceInfo:           hashedDeviceInfo,
		KeyForServerDataBase: crypto.GenKey(deviceInfo),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "$ Ошибка при маршалинге JSON:" + err.Error()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.ForServerReq+config.MediumSizeDataServicePort+"/setavatar", bytes.NewReader(body))
	if err != nil {
		return "$ Ошибка при создании запроса:" + err.Error()
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "$ Сервер не ответил в течение 10 секунд. Завершение программы."
		}
		return "$ Ошибка при выполнении запроса:" + err.Error()
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return "Аватар усшпено загружен"
	case http.StatusBadRequest:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "Не удается обработать ответ сервера"
		}
		resp, _ := crypto.StringDecrypt(string(body), key)
		return "$ " + string(resp)
	case http.StatusForbidden:
		return "$ Не правильный логин или пароль"
	case http.StatusInternalServerError:
		resp, _ := crypto.StringDecrypt(string(body), key)
		return "$ Внутренняя ошибка сервера" + resp
	default:
		return "$ Неизвестная ошибка"
	}
}

func GetAvatar(login, password, loginForSearch, deviceInfo, key string) []byte {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	loginForSearchEcnrypted, err := crypto.StringEncrypt([]byte(loginForSearch), key)
	if err != nil {
		return []byte{}
	}
	loginEncrypted, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return []byte{}
	}
	passwordEncrypted, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return []byte{}
	}

	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	formData := url.Values{}
	formData.Add("loginforsearch", loginForSearchEcnrypted)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("forserver", crypto.GenKey(deviceInfo))
	formData.Add("login", loginEncrypted)
	formData.Add("password", passwordEncrypted)

	respStoragePath, err := client.PostForm(config.ForServerReq+config.MediumSizeDataServicePort+"/getavatar", formData)
	if err != nil {
		fmt.Println(err)
		if os.IsTimeout(err) {
			return []byte{}
		}
		return []byte{}
	}
	defer respStoragePath.Body.Close()

	respBody, err := io.ReadAll(respStoragePath.Body)
	if err != nil {
		return nil
	}

	switch respStoragePath.StatusCode {
	case http.StatusOK:
		// var encAvatar []byte
		// if err := json.Unmarshal(respBody, &encAvatar); err != nil {
		// 	// Если не получилось, возможно сервер вернул JSON-строку или plain string с зашифрованными байтами
		// 	var s string
		// 	if err2 := json.Unmarshal(respBody, &s); err2 == nil {
		// 		encAvatar = []byte(s)
		// 	} else {
		// 		// как запасной вариант — попробовать использовать сырое тело без JSON
		// 		encAvatar = respBody
		// 	}
		// }

		// Расшифровываем полученные байты
		storagePath, err := crypto.StringDecrypt(string(respBody), key)
		if err != nil {
			fmt.Println("decrypt avatar:", err)
			return nil
		}

		fmt.Println(storagePath)

		respAvatarBytes, err := client.Get(config.FilesService + "/" + storagePath)
		if err != nil {
			return []byte{}
		}
		defer respAvatarBytes.Body.Close()

		respBody, err := io.ReadAll(respAvatarBytes.Body)
		if err != nil {
			return []byte{}
		}

		return respBody
	case http.StatusNoContent:
		return []byte{}
	case http.StatusBadRequest:
		return []byte{}
	case http.StatusInternalServerError:
		return []byte{}
	default:
		return []byte{}
	}
}
