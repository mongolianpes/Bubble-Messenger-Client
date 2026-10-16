package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"client/audiodialog"
	"client/auth"
	"client/files"
	"client/messages"
	"client/search"
)

func main() {
	device := "g35c47d72f1c8c8d.SM-A155F.mt6789"
	// resp := Reg("Mongol", "mongoliskii", "1234567890", device)
	// fmt.Println("new reg", resp)
	// resp = Reg("Pes", "pesik", "1234567890", device)
	// fmt.Println(resp)
	// resp := Reg("Mongoliseker", "mongolisek", "1234567890", device)
	// fmt.Println(resp)
	authMongol := Auth("Mongol", "1234567890", device)
	fmt.Println(authMongol)
	authPes := Auth("Pes", "1234567890", device)
	fmt.Println(authPes)
	keyMongol := strings.Split(authMongol, "\n")[0]
	keyPes := strings.Split(authPes, "\n")[0]

	// ТЕСТ ФАЙЛОВ: аватарки
	// avatarBytes, err := os.ReadFile("F:/GolangProjects/BubleMessenger/client/192.webp")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("set avatar pes", SetAvatar("Pes", "1234567890", device, keyPes, avatarBytes))
	// fmt.Println("set avatar mongol", SetAvatar("Mongol", "1234567890", device, keyMongol, avatarBytes))
	// mongolAvatar := GetAvatar("Pes", "1234567890", "Mongol", device, keyPes)
	// fmt.Println(os.WriteFile("mongolAvatar.webp", mongolAvatar, 0644))
	// pesAvatar := GetAvatar("Pes", "1234567890", "Pes", device, keyPes)
	// fmt.Println(os.WriteFile("pesAvatar.webp", pesAvatar, 0644))

	// ТЕСТ ФАЙЛОВ: файлы
	fmt.Println("del messages (mongol): ", DelMessages("Pes", "1234567890", device, keyPes))
	file, err := os.ReadFile("F:/GolangProjects/BubleMessenger/client/192.webp")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("send file", SendFile("Pes", "1234567890", "Mongol", "newfile.webp", device, keyPes, file))
	checkMessages := CheckMessages("Mongol", "1234567890", device, keyMongol)
	checkMessageSplit := strings.Split(checkMessages, "\\\\")
	messageSplit := strings.Split(checkMessageSplit[1], "\\")
	getFile := GetFile("Mongol", "1234567890", messageSplit[2], device, keyMongol)
	fmt.Println(os.WriteFile("get.webp", getFile, 0600))

	// ТЕСТ сообщений
	// fmt.Println("send message: ", SendMessage("Pes", "1234567890", "Mongol", device, "pon", keyPes))
	// fmt.Println("check messages: ", CheckMessages("Mongol", "1234567890", device, keyMongol))
	// fmt.Println("del: ", DelMessages("Mongol", "1234567890", device, keyMongol))
	// fmt.Println("check messages: ", CheckMessages("Mongol", "1234567890", device, keyMongol))

	// ТЕСТ инициализации аудио диалога
	// fmt.Println("del messages (pes): ", DelMessages("Pes", "1234567890", device, keyPes))
	// fmt.Println("del messages (mongol): ", DelMessages("Mongol", "1234567890", device, keyMongol))
	// testECDHAlgorithm()

	// ТЕСТ поиска
	// fmt.Println("search: ", SearchUser("Mongol", "1234567890", device, "Pes", keyMongol, 1))
	// fmt.Println("search: ", SearchUser("Mongol", "1234567890", device, "@Mongoliseker", keyMongol, 1))
	// fmt.Println("search: ", SearchUser("Mongol", "1234567890", device, "Mongol", keyMongol, 1))

	// ТЕСТ аудио диалога
	// fmt.Println("del messages (pes): ", DelMessages("Pes", "1234567890", device, keyPes))
	// fmt.Println("del messages (mongol): ", DelMessages("Mongol", "1234567890", device, keyMongol))
	// testAlgorithmAudioDialog()
}

func testAlgorithmAudioDialog() {
	device := "g35c47d72f1c8c8d.SM-A155F.mt6789"

	authMongol := Auth("Mongol", "1234567890", device)
	fmt.Println(authMongol)
	authPes := Auth("Pes", "1234567890", device)
	fmt.Println(authPes)
	keyMongol := strings.Split(authMongol, "\n")[0]
	keyPes := strings.Split(authPes, "\n")[0]

	init := InitAudioDialog("Mongol", "1234567890", "Pes", device, keyMongol)
	fmt.Println("init audio dialog: ", init)
	initSlice := strings.Split(init, "\\")
	fmt.Println("init slice (mongol): ", initSlice)
	firstAudioDialog := AudioDialog(initSlice[0], initSlice[1], keyMongol, []byte("messeeafge"))
	fmt.Println("send first message (mongol): ", firstAudioDialog)

	fmt.Println("........")
	init2 := CheckMessages("Pes", "1234567890", device, keyPes)
	fmt.Println("check message (pes): ", init2)
	time.Sleep(time.Second * 2)
	initSlice2 := strings.Split(init2, "\\\\")
	initSlice2 = strings.Split(initSlice2[1], "\\")
	fmt.Println("init slice (pes): ", initSlice2)

	go func() {
		i := 1
		for {
			text := fmt.Sprint(i)
			i++
			r := AudioDialog(initSlice[0], initSlice[1], keyMongol, []byte(text))
			responses := strings.Split(r, "\\")
			for _, resp := range responses {
				rByte, err := hex.DecodeString(resp)
				fmt.Println("mongol: ", string(rByte), err)
			}
		}
	}()
	go func() {
		i := 1
		for {
			text := fmt.Sprint(i)
			i++
			r := AudioDialog(initSlice2[1], initSlice2[2], keyMongol, []byte(text))
			responses := strings.Split(r, "\\")
			for _, resp := range responses {
				rByte, err := hex.DecodeString(resp)
				fmt.Println("pes: ", string(rByte), err)
			}
			time.Sleep(time.Microsecond * 2000)
		}
	}()

	select {}

	//для теста остановки диалога
	fmt.Println(StopAudioDialog("Mongol", "1234567890", "Pes", initSlice[0], device, keyMongol))
}

func testECDHAlgorithm() {
	deviceMongol := "fa8a882dasdfasa76daa6c90d8.SM-A155F.mt6789"
	devicePes := "fa8aa882a76daafadsfasdfaf6c90d8.SM-A155F.mt6789"
	passMongol := "1234567890"
	passPes := "1234567890"

	keyMongol := Auth("Mongol", passMongol, deviceMongol)
	keyMongol = strings.Split(keyMongol, "\n")[0]
	fmt.Println("auth Mongol: ", keyMongol)
	keyPes := Auth("Pes", passPes, devicePes)
	keyPes = strings.Split(keyPes, "\n")[0]
	fmt.Println("auth Pes: ", keyPes)

	privKeyMongol := InitDialog("Mongol", passMongol, "Pes", deviceMongol, keyMongol)
	fmt.Println("priv key Mongol: ", privKeyMongol)

	messagesPes := CheckMessages("Pes", passPes, devicePes, keyPes)
	fmt.Println("Messages PES: " + messagesPes)

	messagesPesSplitted := strings.Split(messagesPes, "\\\\")
	fmt.Println("mes splitted pes: ", messagesPesSplitted[1][2:])
	answerPesResult := AnswerInitDialog("Pes", passPes, "Mongol", devicePes, messagesPesSplitted[1][2:], keyPes)
	fmt.Println(answerPesResult)

	messagesMongol := CheckMessages("Mongol", passMongol, deviceMongol, keyMongol)
	fmt.Println("MEssages Mongol: " + messagesMongol)

	messagesMongolSplitted := strings.Split(messagesMongol, "\\\\")
	fmt.Println("mes splitted mongol: ", messagesMongolSplitted[1][2:])
	endInit := EndInitDialog(privKeyMongol, messagesMongolSplitted[1][2:], "Pes")
	fmt.Println(endInit)

	if strings.Split(answerPesResult, "\\")[0] == strings.Split(endInit, "\\")[0] {
		fmt.Println("Успешно")
	} else {
		fmt.Println("Общий секрет не совпадает")
	}
}

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
	return messages.InitAudioDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, key)
}

//export StopAudioDialog
func StopAudioDialog(senderLogin, senderPassword, receiverLogin, dialogUUID, deviceInfo, key string) string {
	return messages.StopAudioDialog(senderLogin, senderPassword, receiverLogin, dialogUUID, deviceInfo, key)
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
