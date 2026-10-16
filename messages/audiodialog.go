package messages

import "fmt"

func InitAudioDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, key string) string {
	return sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, "a\\", key)
}

func StopAudioDialog(senderLogin, senderPassword, receiverLogin, dialogUUID, deviceInfo, key string) string {
	return sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, fmt.Sprintf("as\\%v", dialogUUID), key)
}
