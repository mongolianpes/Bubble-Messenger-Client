package messages

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExchangeSharedSecret(t *testing.T) {
	loginInitor, _, passwordInitor, deviceInitor, keyInitor, err := RegAndAuthTestUser()
	require.NoError(t, err)
	loginReceiver, _, passwordReceiver, deviceReceiver, keyReceiver, err := RegAndAuthTestUser()
	require.NoError(t, err)

	privKeyInitor := InitDialog(loginInitor, passwordInitor, loginReceiver, deviceInitor, keyInitor)

	messagesReceiver := CheckMessages(loginReceiver, passwordReceiver, deviceReceiver, keyReceiver)
	messagesReceiverSplitted := strings.Split(messagesReceiver, "\\\\")

	answerReceiverResult := AnswerInitDialog(loginReceiver, passwordReceiver, loginInitor, deviceReceiver, messagesReceiverSplitted[1][2:], keyReceiver)

	messagesInitor := CheckMessages(loginInitor, passwordInitor, deviceInitor, keyInitor)

	messagesInitorSplitted := strings.Split(messagesInitor, "\\\\")
	endInit := EndInitDialog(privKeyInitor, messagesInitorSplitted[1][2:], loginReceiver)

	require.Equal(t, strings.Split(answerReceiverResult, "\\")[0], strings.Split(endInit, "\\")[0])
}
