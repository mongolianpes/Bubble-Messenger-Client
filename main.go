package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"client/audiodialog"
	"client/auth"
	"client/files"
	"client/messages"
	"client/search"
)

func main() {
	// // dialogKey := "345ee540ab4c6d32e733a6dd8fe3a9b393f52cf95bc89821209a2989f1c94b32"
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
	fmt.Println("del messages (pes): ", DelMessages("Pes", "1234567890", device, keyPes))
	fmt.Println("del messages (mongol): ", DelMessages("Mongol", "1234567890", device, keyMongol))
	testAlgorithmAudioDialog()

	// fmt.Println("resp: ", SendFile("Mongol", "1234567890", "Pes", "hello", device, keyMongol, []byte("hello")))
	// fmt.Println("resp get file: ", string(GetFile("Pes", "1234567890", "storage/users/2026/10/12/18/0b72d726-19db-4eff-8005-e66453cd4683", device, keyPes)))
	// fmt.Println("resp del file: ", DelFile("Pes", "1234567890", "storage/users/2026/10/12/18/0b72d726-19db-4eff-8005-e66453cd4683", device, keyPes))
	// fmt.Println(SendMessage("Mongol", "1234567890", "Pes", device, rand.Text()+rand.Text()+rand.Text(), keyMongol))
	// for i := 0; i < 1; i++ {
	// fmt.Println(SendMessage("Mongol", "1234567890", "Pes", device, rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text()+rand.Text(), keyMongol))
	// }
	// fmt.Println(CheckMessages("Pes", "1234567890", device, keyPes))
	// fmt.Println(CheckMessages("Pes", "1234567890", device, keyPes))
	// fmt.Println(DelMessages("Pes", "1234567890", device, keyPes))
	// fmt.Println(CheckMessages("Pes", "1234567890", device, keyPes))

	// fmt.Println(SendMessage("Mongol", "1234567890", "psina2", device, "obichnoe eche ras", key))
	// fmt.Println(InitDialog("Mongol", "1234567890", "psina2", device, key))
	// fmt.Println(CheckMessages("TupMes2", "1234567890", device, key))

	// fmt.Println(SendMessage("TupMes", "1234567890", "TupMes2", device, "ok la", key))

	// auth2 := Auth("TupMes", "1234567890", device)
	// fmt.Println(auth2)
	// key2 := strings.Split(auth2, "\n")[0]

	// fmt.Println(CheckMessages("TupMes2", "1234567890", device, key2))

	// // читаем файл в байтовый срез
	// avatar, err := ioutil.ReadFile("192.webp")
	// if err != nil {
	// 	log.Fatalf("Ошибка чтения файла: %v", err)
	// }

	// fmt.Println(avatar)
	// SetAvatar("Psina2", "1234567890", device, key, avatar)
	// ava := GetAvatar("Psina2", "Mongoliseker2", device, key)

	// fmt.Println(ava)
	// os.WriteFile("new.webp", ava, 0644)

	// // mess := CheckMessages("Mongoliseker2", "1234567890", device, key)
	// // fmt.Println(mess)

	// // messageee := SendMessage("Mongoliseker2", "1234567890", "Psina2", device, "Понял тебя 👍", key)
	// // fmt.Println(messageee)

	// // authSlice := strings.Split(auth, "\n")
	// // key := authSlice[0]

	// for {
	// 	devicee := device + rand.Text()
	// 	reg := Reg("1234567"+rand.Text(), rand.Text(), "1234567890", devicee)
	// 	fmt.Println(reg)
	// }
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

//	func AudioDialog(UUID, UserID, key string, message []byte) map[int][]byte {
//		result := make(map[int][]byte)
//		messageEncrypted, err := crypto.StringEncryptByte(message, key)
//		if err != nil {
//			result[1] = []byte("$ ошибка шифрования сообщения")
//			return result
//		}
//		payload := RequestAudio{
//			UUID:    UUID,
//			UserID:  UserID,
//			Message: messageEncrypted,
//		}
//
//		body, err := json.Marshal(payload)
//		if err != nil {
//			result[1] = []byte("$ Ошибка при маршалинге JSON:" + err.Error())
//			return result
//		}
//
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		defer cancel()
//
//		req, err := http.NewRequestWithContext(ctx, http.MethodPost, forServerReq+"/audiodialog", bytes.NewReader(body))
//		if err != nil {
//			result[1] = []byte("$ Ошибка при создании запроса:" + err.Error())
//			return result
//		}
//		req.Header.Set("Content-Type", "application/json")
//
//		client := &http.Client{}
//
//		resp, err := client.Do(req)
//		if err != nil {
//			if ctx.Err() == context.DeadlineExceeded {
//				result[1] = []byte("$ Сервер не ответил в течение 10 секунд. Завершение программы.")
//				return result
//			}
//			result[1] = []byte("$ Ошибка при выполнении запроса:" + err.Error())
//			return result
//		}
//		defer resp.Body.Close()
//
//		respBody, err := io.ReadAll(resp.Body)
//		if err != nil {
//			result[1] = []byte("$ ошибка чтения ответа")
//			return result
//		}
//
//		switch resp.StatusCode {
//		case http.StatusOK:
//			// Распарсим ответ в map[string]map[time.Time][]byte
//			var data map[string]map[string][]byte
//			if err := json.Unmarshal(respBody, &data); err != nil {
//				result[1] = []byte("$ Ошибка при парсинге ответа: " + err.Error())
//				return result
//			}
//
//			var count = 1
//			for user, messages := range data {
//				if user == UserID {
//					continue
//				}
//
//				for _, msg := range messages {
//					// ts приходит как строка, можно преобразовать в time.Time при необходимости
//
//					// start := strings.Index(ts, "T") + 1
//					// end := strings.Index(ts, "+")
//					// time := ts[start:end]
//
//					msgDecrypt, err := crypto.StringDecryptByte(msg, key)
//					if err != nil {
//						continue
//					}
//
//					result[count] = msgDecrypt
//				}
//			}
//			return result
//		case http.StatusBadRequest:
//			result[1] = []byte("$ Проверьте правильность ввода данных")
//			return result
//		case http.StatusForbidden:
//			result[1] = []byte("$ Не правильный логин или пароль")
//			return result
//		case http.StatusInternalServerError:
//			result[1] = []byte("$ Внутренняя ошибка сервера")
//			return result
//		default:
//			result[1] = []byte("$ Неизвестная ошибка")
//			return result
//		}
//	}
//
// map[int][]byte
//
//export AudioDialog

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
