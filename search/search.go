package search

import (
	"client/config"
	"client/crypto"
	"client/request"
	"encoding/json"
	"io"
	"net/http"
)

type SearchUserRequest struct {
	LoginForSearch string
}

type FindUser struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

func SearchUser(login, password, deviceInfo, loginForSearch, key string, numSeaerchRequest int) string {
	req := SearchUserRequest{
		LoginForSearch: loginForSearch,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "$ не удалось преобразовать запрос в JSON" + err.Error()
	}

	resp, err := request.SendRequest(login, password, deviceInfo, key, config.ForServerReq+config.MainServicePort+"/", "searchuser", reqBody)
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
		body, err := crypto.StringDecryptByte(body, key)
		if err != nil {
			return "$decrypt error " + err.Error()
		}

		var findUsers []FindUser
		var out string
		if err := json.Unmarshal(body, &findUsers); err != nil {
			return "$ошибка обработки ответа, попробуй еще раз " + err.Error()
		}

		// out += strconv.Itoa(resp.NumSeaerchRequest)

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
