package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"client/config"
	"client/crypto"
	"client/request"
)

type SetAvatarRequest struct {
	Avatar []byte `json:"avatar"`
	Login  string `json:"login"`
}

type GetAvatarRequest struct {
	LoginForSearch string
}

func SetAvatar(login, password, deviceInfo, key string, avatar []byte) string {
	req := SetAvatarRequest{
		Avatar: avatar,
		Login:  login,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "$ не удалось преобразовать запрос в JSON" + err.Error()
	}

	resp, err := request.SendRequest(login, password, deviceInfo, key, config.ForServerReq+config.MediumSizeDataServicePort+"/", "setavatar", reqBody)
	if err != nil {
		return "$ не удалось отправить реквест " + err.Error()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "$ Не удалось прочитать ответ от сервера"
	}

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
	req := GetAvatarRequest{
		LoginForSearch: loginForSearch,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return []byte("$ не удалось преобразовать запрос в JSON" + err.Error())
	}

	resp, err := request.SendRequest(login, password, deviceInfo, key, config.ForServerReq+config.MediumSizeDataServicePort+"/", "getavatar", reqBody)
	if err != nil {
		return []byte("$ не удалось отправить реквест " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte("$ ошибка чтения ответа")
	}

	switch resp.StatusCode {
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
		storagePath, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			fmt.Println("decrypt avatar:", err)
			return nil
		}

		fmt.Println(storagePath)

		client := &http.Client{
			Timeout: 10 * time.Second,
		}

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
