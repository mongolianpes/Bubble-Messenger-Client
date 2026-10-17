package audiodialog

import (
	"client/messages"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAudioDialog(t *testing.T) {
	loginInitor, _, passwordInitor, deviceInitor, keyInitor, err := RegAndAuthTestUser()
	require.NoError(t, err)
	loginReceiver, _, passwordReceiver, deviceReceiver, keyReceiver, err := RegAndAuthTestUser()
	require.NoError(t, err)

	// в реальном диалоге просходит обмен между пользователями через сообщения
	keyDialog := keyInitor

	init := InitAudioDialog(loginInitor, passwordInitor, loginReceiver, deviceInitor, keyInitor)
	initSlice := strings.Split(init, "\\")

	init2 := messages.CheckMessages(loginReceiver, passwordReceiver, deviceReceiver, keyReceiver)
	initSlice2 := strings.Split(init2, "\\\\")
	initSlice2 = strings.Split(initSlice2[1], "\\")

	initorDialogUUID := initSlice[0]
	initorDialodUserID := initSlice[1]
	receiverDialogUUID := initSlice2[1]
	receiverDialogUserID := initSlice2[2]

	initorMessages := make(chan string, 10)
	receiverMessages := make(chan string, 10)

	stopDialog := false

	go func() {
		i := 1
		for {
			text := fmt.Sprint(i)
			i++
			initorMessages <- text
			r := AudioDialog(initorDialogUUID, initorDialodUserID, keyDialog, []byte(text))
			if stopDialog {
				return
			}

			responses := strings.Split(r, "\\")
			for _, resp := range responses {
				rByte, err := hex.DecodeString(resp)
				require.NoError(t, err)
				if string(rByte) != "" {
					require.Equal(t, <-receiverMessages, string(rByte))
				}
			}
		}
	}()

	go func() {
		i := 1
		for {
			text := fmt.Sprint(i)
			i++
			receiverMessages <- text
			r := AudioDialog(receiverDialogUUID, receiverDialogUserID, keyDialog, []byte(text))
			if stopDialog {
				return
			}

			responses := strings.Split(r, "\\")
			for _, resp := range responses {
				rByte, err := hex.DecodeString(resp)
				require.NoError(t, err)
				if string(rByte) != "" {
					require.Equal(t, <-initorMessages, string(rByte))
				}
			}
			time.Sleep(time.Microsecond * 2000)
		}
	}()

	time.Sleep(time.Second * 3)

	stopDialog = true
	stopDialogResult := StopAudioDialog(loginInitor, passwordInitor, loginReceiver, initorDialogUUID, deviceInitor, keyInitor)
	require.Equal(t, stopDialogResult, "Audio dialog stop")
}
