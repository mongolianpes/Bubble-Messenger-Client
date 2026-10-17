package messages

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendCheckAndDelMessages(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		senderLogin, _, senderPassword, senderDevice, senderKey, err := RegAndAuthTestUser()
		require.NoError(t, err)
		receiverLogin, _, receiverPassword, receiverDevice, receiverKey, err := RegAndAuthTestUser()
		require.NoError(t, err)
		messageText := generateTestRandValidStr()

		send := SendMessage(senderLogin, senderPassword, receiverLogin, senderDevice, messageText, senderKey)
		require.Equal(t, send, "Сообщение отправлено")

		check := CheckMessages(receiverLogin, receiverPassword, receiverDevice, receiverKey)
		checkSplitted := strings.Split(check, "\\\\")
		require.Equal(t, checkSplitted[0], senderLogin)
		require.Equal(t, checkSplitted[1], messageText)

		del := DelMessages(receiverLogin, receiverPassword, receiverDevice, receiverKey)
		require.Equal(t, del, "Успешно удалено")

		check2 := CheckMessages(receiverLogin, receiverPassword, receiverDevice, receiverKey)
		require.Empty(t, check2)
	})
}
