package auth

import (
	"bytes"
	"context"
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"client/config"
)

// KeyExchangeRequest определяет структуру JSON запроса
type KeyExchangeRequest struct {
	ID              string `json:"id"`
	ClientPublicKey string `json:"public_key"`
	IsRegistring    bool   `json:"is_registring"`
}

// KeyExchangeResponse определяет структуру JSON ответа от сервера
type KeyExchangeResponse struct {
	ServerPublicKey string `json:"public_key"`
}

// KeyExchangeResult представляет результат обмена ключами
type KeyExchangeResult struct {
	SharedSecret []byte
	Error        error
}

func exchangeKeys(deviceInfo string, isRegistring bool) *KeyExchangeResult {
	timeout := 10 * time.Second
	result := make(chan *KeyExchangeResult, 1)

	go func() {
		defer close(result)

		// Генерация пары ключей для кривой P256
		privateKey, err := ecdh.P256().GenerateKey(rand.Reader)
		if err != nil {
			result <- &KeyExchangeResult{Error: fmt.Errorf("ошибка генерации ключей: %v", err)}
			return
		}

		// Получение открытого ключа в виде байт
		publicKeyBytes := privateKey.PublicKey().Bytes()

		// Подготовка запроса
		requestBody := KeyExchangeRequest{
			ID:              deviceInfo,
			ClientPublicKey: base64.RawURLEncoding.EncodeToString(publicKeyBytes),
			IsRegistring:    isRegistring,
		}

		requestJSON, err := json.Marshal(requestBody)
		if err != nil {
			result <- &KeyExchangeResult{Error: fmt.Errorf("ошибка маршалинга JSON: %v", err)}
			return
		}

		// Создание контекста с таймаутом
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Создание HTTP запроса с контекстом
		req, err := http.NewRequestWithContext(ctx, "POST", config.ForServerReq+config.MainServicePort+"/exchangekey", bytes.NewBuffer(requestJSON))
		if err != nil {
			result <- &KeyExchangeResult{Error: fmt.Errorf("ошибка создания запроса: %v", err)}
			return
		}
		req.Header.Set("Content-Type", "application/json")

		// Выполнение запроса
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			// Проверяем, был ли это таймаут
			if ctx.Err() == context.DeadlineExceeded {
				result <- &KeyExchangeResult{Error: fmt.Errorf("Сервер не отвечает")}
			} else {
				result <- &KeyExchangeResult{Error: fmt.Errorf("ошибка отправки запроса: %v", err)}
			}
			return
		}
		defer resp.Body.Close()

		// Чтение ответа от сервера
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			result <- &KeyExchangeResult{Error: fmt.Errorf("ошибка чтения ответа: %v", err)}
			return
		}

		switch resp.StatusCode {
		case http.StatusOK:
			var serverResponse KeyExchangeResponse
			if err := json.Unmarshal(body, &serverResponse); err != nil {
				result <- &KeyExchangeResult{Error: fmt.Errorf("ошибка парсинга ответа: %v", err)}
				return
			}

			// Декодирование публичного ключа сервера из hex-строки
			serverPubKeyBytes, err := base64.RawURLEncoding.DecodeString(serverResponse.ServerPublicKey)
			if err != nil {
				result <- &KeyExchangeResult{Error: fmt.Errorf("ошибка декодирования ключа сервера: %v", err)}
				return
			}

			// Преобразование байт в ecdh.PublicKey
			serverPublicKey, err := ecdh.P256().NewPublicKey(serverPubKeyBytes)
			if err != nil {
				result <- &KeyExchangeResult{Error: fmt.Errorf("ошибка создания публичного ключа сервера: %v", err)}
				return
			}

			// ВЫЧИСЛЕНИЕ ОБЩЕГО СЕКРЕТА
			sharedSecret, err := privateKey.ECDH(serverPublicKey)
			if err != nil {
				result <- &KeyExchangeResult{Error: fmt.Errorf("ошибка вычисления общего секрета: %v", err)}
				return
			}

			result <- &KeyExchangeResult{SharedSecret: sharedSecret}
		case http.StatusBadRequest:
			result <- &KeyExchangeResult{Error: fmt.Errorf("Не удалось обработать запрос")}
		case http.StatusInternalServerError:
			result <- &KeyExchangeResult{Error: fmt.Errorf("Внутреняя ошибка сервера")}
		default:
			result <- &KeyExchangeResult{Error: fmt.Errorf("Ошибка сервера: %s\nТело ответа: %s", resp.Status, string(body))}
		}
	}()

	// Ожидаем результат или таймаут
	select {
	case res := <-result:
		return res
	case <-time.After(timeout):
		return &KeyExchangeResult{Error: fmt.Errorf("Сервер не отвечает")}
	}
}
