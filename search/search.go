package search

import (
	"client/config"
	"client/crypto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type FindUser struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

func SearchUser(login, password, deviceInfo, loginForSearch, key string, numSeaerchRequest int) string {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	loginForSearchEncrypt, err := crypto.StringEncrypt([]byte(loginForSearch), key)
	if err != nil {
		return "$ Необходимо заново авторизоваться"
	}
	loginEncrypt, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return "$ Необходимо заново авторизоваться"
	}
	passwordEncrypt, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return "$ Необходимо заново авторизоваться"
	}

	formData := url.Values{}
	formData.Add("login", loginEncrypt)
	formData.Add("password", passwordEncrypt)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("loginforsearch", loginForSearchEncrypt)
	formData.Add("forserver", crypto.GenKey(deviceInfo))

	resp, err := client.PostForm(config.ForServerReq+config.MainServicePort+"/searchuser", formData)
	if err != nil {
		if os.IsTimeout(err) {
			return "$ Сервер не отвечает по причине долгого ответа"
		}
		return "$ Сервер не отвечает по какому то дерьму"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "$ Не удалось прочитать ответ от сервера"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		body, err = crypto.StringDecryptByte(body, key)
		if err != nil {
			return "$ " + err.Error()
		}

		var findUsers []FindUser
		var out string
		if err := json.Unmarshal(body, &findUsers); err != nil {
			return "$ ошибка обработки ответа, попробуй еще раз " + err.Error()
		}

		// out += strconv.Itoa(resp.NumSeaerchRequest)

		fmt.Println(findUsers)
		for findUser := range findUsers {
			if findUsers[findUser].Login == login {
				continue
			}

			if out != "" {
				out += "\n"
			}

			out += findUsers[findUser].Login + " "

			out += findUsers[findUser].Name
		}

		return out
	case http.StatusInternalServerError:
		decryptedResp, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ ошибка чтения ответа" + err.Error()
		}
		return "$ " + decryptedResp
	case http.StatusBadRequest:
		return "$ " + string(body)
	default:
		return "$ " + string(body)
	}
}
