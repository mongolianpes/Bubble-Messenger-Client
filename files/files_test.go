package files

import (
	"os"
	"strings"
	"testing"

	"client/messages"

	"github.com/stretchr/testify/require"
)

func TestSendAndGetFile(t *testing.T) {
	fileBytes, err := os.ReadFile("F:/GolangProjects/BubleMessenger/client/192.webp")
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		senderLogin, _, senderPassword, senderDevice, senderKey, err := RegAndAuthTestUser()
		require.NoError(t, err)
		receiverLogin, _, receiverPassword, receiverDevice, receiverKey, err := RegAndAuthTestUser()
		require.NoError(t, err)
		fileName := generateTestRandValidStr()

		sendFile := SendFile(senderLogin, senderPassword, receiverLogin, fileName, senderDevice, senderKey, fileBytes)
		require.NotContains(t, sendFile, "$")

		checkMessages := messages.CheckMessages(receiverLogin, receiverPassword, receiverDevice, receiverKey)
		checkMessagesSplitted := strings.Split(checkMessages, "\\\\")
		textMessageSplitted := strings.Split(checkMessagesSplitted[1], "\\")
		require.Equal(t, textMessageSplitted[0], "p")
		require.Equal(t, textMessageSplitted[1], fileName)
		fileStoragePath := textMessageSplitted[2]

		gettedFile := GetFile(receiverLogin, receiverPassword, fileStoragePath, receiverDevice, receiverKey)
		require.Equal(t, gettedFile, fileBytes)

		del := DelFile(receiverLogin, receiverPassword, fileStoragePath, receiverDevice, receiverKey)
		require.Equal(t, del, "Успешно удалено")

		gettedFile = GetFile(receiverLogin, receiverPassword, fileStoragePath, receiverDevice, receiverKey)
		require.Contains(t, string(gettedFile), "$")
	})

	t.Run("get not exist file", func(t *testing.T) {
		login, _, password, device, key, err := RegAndAuthTestUser()
		require.NoError(t, err)
		notExistStoragePath := generateTestRandValidStr()

		gettedFile := GetFile(login, password, notExistStoragePath, device, key)
		require.Contains(t, string(gettedFile), "$")
	})

	t.Run("del not exist file", func(t *testing.T) {
		login, _, password, device, key, err := RegAndAuthTestUser()
		require.NoError(t, err)
		notExistStoragePath := generateTestRandValidStr()

		del := DelFile(login, password, notExistStoragePath, device, key)
		require.Contains(t, del, "$")
	})
}
