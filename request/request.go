package request

import (
	"bytes"
	"client/crypto"
	"crypto/rand"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type TLSData struct {
	Device      string
	Login       string
	Password    string
	SendTime    string
	Nonce       string
	RequestBody []byte
}

func SendRequest(login, password, device, key, host, path string, body []byte) (*http.Response, error) {
	hashedDeviceInfo := crypto.HashDeviceInfo(device, login)

	loginEncrypt, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return nil, errors.New("$ Необходимо заново авторизоваться")
	}
	passwordEncrypt, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return nil, errors.New("$ Необходимо заново авторизоваться")
	}
	bodyEncrypt, err := crypto.StringEncryptByte(body, key)
	if err != nil {
		return nil, errors.New("$ Необходимо заново авторизоваться")
	}
	pathEncrypt, err := crypto.StringEncrypt([]byte(path), key)
	if err != nil {
		return nil, errors.New("$ Необходимо заново авторизоваться")
	}
	nonceEncrypt, err := crypto.StringEncrypt([]byte(rand.Text()), key)
	if err != nil {
		return nil, errors.New("$ Необходимо заново авторизоваться")
	}
	sendTimeEncrypt, err := crypto.StringEncrypt([]byte(time.Now().Format(time.RFC3339)), key)
	if err != nil {
		return nil, errors.New("$ Необходимо заново авторизоваться")
	}

	payload := TLSData{
		Login:       loginEncrypt,
		Password:    passwordEncrypt,
		Device:      hashedDeviceInfo,
		SendTime:    sendTimeEncrypt,
		Nonce:       nonceEncrypt,
		RequestBody: []byte(bodyEncrypt),
	}

	requestJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, host+pathEncrypt, bytes.NewReader(requestJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
