package messages

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"client/config"
	"client/crypto"
)

type RequestCheckMessages struct {
	Gzip     bool   `json:"Gzip"`
	Messages []byte `json:"Messages"` // или string, если хочешь base64
}

type Message struct {
	Sender   string    `json:"SenderLogin"`
	Message  string    `json:"Text"`
	SendTime time.Time `json:"SendTime"`
}

func SendMessage(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key string) string {
	resp := sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key)
	return resp
}

func sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key string) string {
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

func CheckMessages(login, password, deviceInfo, key string) string {
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
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	formData := url.Values{}
	formData.Add("login", loginEcnrypted)
	formData.Add("password", passwordEncrypted)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("forserver", crypto.GenKey(deviceInfo))

	resp, err := client.PostForm(config.ForServerReq+config.MainServicePort+"/checkmessage", formData)
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
		var response RequestCheckMessages
		if err := json.Unmarshal(body, &response); err != nil {
			return "$ Некорректный формат ответа от сервера"
		}
		decryptedResp, err := crypto.StringDecryptByte(response.Messages, key)
		if err != nil {
			return "$ Ошибка расшифровки ответа"
		}

		// if response.Gzip {
		// 	decryptedResp, err = decompressGzip(decryptedResp)
		// 	if err != nil {
		// 		return "$ не удалось распаковать сообщения"
		// 	}
		// }

		lines := strings.Split(string(decryptedResp), "\n")
		var result strings.Builder

		for li, line := range lines {
			fmt.Println(line)
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			var msg Message
			if err := json.Unmarshal([]byte(line), &msg); err != nil {
				return "$ ошибка обработки ответа, попробуй еще раз " + err.Error()
			}

			date := msg.SendTime.Format("02.01.2006 15:04")
			_, offsetSeconds := msg.SendTime.Zone()
			offsetHours := offsetSeconds / 3600
			offsetShort := fmt.Sprintf("%+03d", offsetHours)

			result.WriteString(msg.Sender)
			result.WriteString("\\\\")
			result.WriteString(msg.Message)
			result.WriteString("\\\\")
			result.WriteString(date)
			result.WriteString(" ")
			result.WriteString(offsetShort)

			// добавляем перевод строки между записями
			if li < len(lines)-1 {
				result.WriteString("\n")
			}
		}

		out := result.String()

		return out
	case http.StatusBadRequest:
		return "$ Проверьте правильность ввода данных"
	case http.StatusForbidden:
		return "$ Не правильный логин или пароль"
	case http.StatusInternalServerError:
		resp, _ := crypto.StringDecrypt(string(body), key)
		return "$ Внутренняя ошибка сервера" + resp
	default:
		return "$ Неизвестная ошибка"
	}
}

func DelMessages(login, password, deviceInfo, key string) string {
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
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	formData := url.Values{}
	formData.Add("login", loginEcnrypted)
	formData.Add("password", passwordEncrypted)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("forserver", crypto.GenKey(deviceInfo))

	resp, err := client.PostForm(config.ForServerReq+config.MediumSizeDataServicePort+"/delmessages", formData)
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
		resp, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ ошибка при расшифровке " + err.Error()
		}
		return "$ Проверьте правильность ввода данных " + resp
	case http.StatusForbidden:
		return "$ Не правильный логин или пароль" + string(body)
	case http.StatusInternalServerError:
		return "$ Внутренняя ошибка сервера" + string(body)
	default:
		return "$ Неизвестная ошибка" + string(body)
	}
}

func compressGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decompressGzip(compressed []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r)
}
