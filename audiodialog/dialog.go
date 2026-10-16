package audiodialog

import (
	"bytes"
	"client/config"
	"client/crypto"
	"context"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type RequestAudio struct {
	UUID    string `json:"dialog_id"`
	UserID  string `json:"user_id"`
	Message []byte `json:"message"`
}

type MessageAudioDialog struct {
	Time  int64  `json:"time"`
	Audio []byte `json:"audio"`
}

func AudioDialog(UUID, UserID, key string, message []byte) string {
	encrypted, err := crypto.StringEncryptByte(message, key)
	if err != nil {
		return "$ ошибка шифрования сообщения"
	}

	payload := RequestAudio{
		UUID:    UUID,
		UserID:  UserID,
		Message: encrypted,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "$ Ошибка маршалинга JSON"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AudioDialogService+"/exchangeaudio", bytes.NewReader(body))
	if err != nil {
		return "$ Ошибка создания запроса"
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "$ Ошибка выполнения запроса"
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "$ Ошибка чтения ответа"
	}

	if resp.StatusCode != http.StatusOK {
		return "$ Ошибка сервера: " + string(raw) + ":"
	}

	var data map[string][]MessageAudioDialog
	if err := json.Unmarshal(raw, &data); err != nil {
		return "$ Ошибка парсинга ответа: " + err.Error()
	}

	result := ""
	for user, msgs := range data {
		if user == UserID {
			continue
		}
		for _, msg := range msgs {
			dec, err := crypto.StringDecryptByte(msg.Audio, key)
			if err == nil {
				if len(result) == 0 {
					result = hex.EncodeToString(dec)
				} else {
					result += "\\" + hex.EncodeToString(dec)
				}
			}
		}
	}

	return result
}
