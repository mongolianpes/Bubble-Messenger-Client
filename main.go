package http

import (
	"client/audiodialog"
	"client/auth"
	"client/files"
	"client/messages"
	"client/search"
)

//export Reg
func Reg(login, name, password, deviceInfo string) string {
	return auth.Reg(login, name, password, deviceInfo)
}

//export Auth
func Auth(login, password, deviceInfo string) string {
	return auth.Auth(login, password, deviceInfo)
}

//export SearchUser
func SearchUser(login, password, deviceInfo, loginForSearch, key string, numSeaerchRequest int) string {
	return search.SearchUser(login, password, deviceInfo, loginForSearch, key, numSeaerchRequest)
}

//export SendMessage
func SendMessage(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key string) string {
	return messages.SendMessage(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key)
}

//export CheckMessages
func CheckMessages(login, password, deviceInfo, key string) string {
	return messages.CheckMessages(login, password, deviceInfo, key)
}

//export DelMessages
func DelMessages(login, password, deviceInfo, key string) string {
	return messages.DelMessages(login, password, deviceInfo, key)
}

//export InitDialog
func InitDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, key string) string {
	return messages.InitDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, key)
}

//export AnswerInitDialog
func AnswerInitDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, pubKeyBase64, key string) string {
	return messages.AnswerInitDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, pubKeyBase64, key)
}

//export EndInitDialog
func EndInitDialog(privBase64, user2PubBase64, loginCommunicating string) string {
	return messages.EndInitDialog(privBase64, user2PubBase64, loginCommunicating)
}

//export SetAvatar
func SetAvatar(login, password, deviceInfo, key string, avatar []byte) string {
	return files.SetAvatar(login, password, deviceInfo, key, avatar)
}

//export GetAvatar
func GetAvatar(login, password, loginForSearch, deviceInfo, key string) []byte {
	return files.GetAvatar(login, password, loginForSearch, deviceInfo, key)
}

//export InitAudioDialog
func InitAudioDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, key string) string {
	return audiodialog.InitAudioDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, key)
}

//export StopAudioDialog
func StopAudioDialog(senderLogin, senderPassword, receiverLogin, dialogUUID, deviceInfo, key string) string {
	return audiodialog.StopAudioDialog(senderLogin, senderPassword, receiverLogin, dialogUUID, deviceInfo, key)
}

func AudioDialog(UUID, UserID, key string, message []byte) string {
	return audiodialog.AudioDialog(UUID, UserID, key, message)
}

//export SendFile
func SendFile(login, password, receiverLogin, fileName, deviceInfo, key string, file []byte) string {
	return files.SendFile(login, password, receiverLogin, fileName, deviceInfo, key, file)
}

//export GetFile
func GetFile(login, password, fileName, deviceInfo, key string) []byte {
	return files.GetFile(login, password, fileName, deviceInfo, key)
}

//export DelFile
func DelFile(login, password, fileName, deviceInfo, key string) string {
	return files.DelFile(login, password, fileName, deviceInfo, key)
}
