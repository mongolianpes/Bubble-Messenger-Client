package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"client/auth"
	"client/crypto"
	"client/files"
)

type SearchUserResponse struct {
	Users []string
}

type ForDelMessages struct {
	Login    string
	Password string
	Device   string
	Messages []string
}

type FindUser struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

type Message struct {
	Sender   string    `json:"Sender"`
	Message  string    `json:"Message"`
	SendTime time.Time `json:"SendTime"`
}

const (
	// forServerReq string = "http://185.121.15.223:"
	forServerReq              = "http://localhost:"
	mainServicePort           = "23099"
	mediumSizeDataServicePort = "23098"
	audioDialogServicePort    = "23097"
	designationOfInitDialog   = "i"
	designationOfAnswerDialog = "a"
)

type Avatar struct {
	Avatar               []byte `json:"avatar"`
	Login                string `json:"login"`
	Password             string `json:"password"`
	DeviceInfo           string `json:"deviceinfo"`
	KeyForServerDataBase string `json:"forserver"`
}

type RequestAudio struct {
	UUID    string
	UserID  string
	Message []byte
}

type RequestCheckMessages struct {
	Gzip     bool   `json:"Gzip"`
	Messages []byte `json:"Messages"` // или string, если хочешь base64
}

func main() {
	// // dialogKey := "345ee540ab4c6d32e733a6dd8fe3a9b393f52cf95bc89821209a2989f1c94b32"
	device := "g35c47d72f1c8c8d.SM-A155F.mt6789"
	// resp := Reg("Mongol", "mongoliskii", "1234567890", device)
	// fmt.Println("new reg", resp)
	// resp = Reg("Pes", "pesik", "1234567890", device)
	// fmt.Println(resp)
	// testECDHAlogoritm()
	authMongol := Auth("Mongol", "1234567890", device)
	fmt.Println(authMongol)
	authPes := Auth("Pes", "1234567890", device)
	fmt.Println(authPes)
	// keyMongol := strings.Split(authMongol, "\n")[0]
	// keyPes := strings.Split(authPes, "\n")[0]

	// fmt.Println("check messages: ", CheckMessages("Mongol", "1234567890", device, keyMongol))
	// fmt.Println("del: ", DelMessages("Mongol", "1234567890", device, keyMongol))
	// fmt.Println(SearchUser("Mongol", device, "Pes", keyMongol, 1))
	// fmt.Println(SearchUser("Mongol", device, "Mongoliseker", keyMongol, 1))
	// fmt.Println(SearchUser("Mongol", device, "Mongol", keyMongol, 1))

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
	fmt.Println(init)
	initSlice := strings.Split(init, "\\")
	firstAudioDialog := AudioDialog(initSlice[0], initSlice[1], keyMongol, []byte("messeeafge"))
	fmt.Println(firstAudioDialog)

	fmt.Println("........")
	init2 := CheckMessages("Pes", "1234567890", device, keyPes)
	fmt.Println(init2)
	time.Sleep(time.Second * 2)
	initSlice2 := strings.Split(init2, "\\\\")
	initSlice2 = strings.Split(initSlice2[1], "\\")
	fmt.Println(initSlice2)

	go func() {
		i := 1
		for {
			text := fmt.Sprint(i)
			i++
			r := AudioDialog(initSlice[0], initSlice[1], keyMongol, []byte(text))
			fmt.Println("уещ ", string(r))
		}
	}()
	go func() {
		for {
			text := rand.Text()
			r := AudioDialog(initSlice[0], initSlice2[3], keyMongol, []byte(text))
			s := strings.Split(r, "\\")
			i := 1
			for _, a := range s {
				if i == 0 || i == 2 || i == 4 || i == 6 {
					continue
				}
				b, _ := hex.DecodeString(a)
				fmt.Println("уещooooo ", string(b))
				i += 2
			}
			time.Sleep(time.Microsecond * 20)
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
	keyPes := Auth("Mongoliseker", passPes, devicePes)
	keyPes = strings.Split(keyPes, "\n")[0]
	fmt.Println("auth Mongoliseker: ", keyPes)

	privKeyMongol := InitDialog("Mongol", passMongol, "Mongoliseker", deviceMongol, keyMongol)
	fmt.Println("priv key Mongol: ", privKeyMongol)

	messagesPes := CheckMessages("Mongoliseker", passPes, devicePes, keyPes)
	fmt.Println("Messages PES: " + messagesPes)

	messagesPesSplitted := strings.Split(messagesPes, "\\\\")
	fmt.Println("mes splitted pes: ", messagesPesSplitted[1][2:])
	answerPesResult := AnswerInitDialog("Mongoliseker", passPes, "Mongol", devicePes, messagesPesSplitted[1][2:], keyPes)
	fmt.Println(answerPesResult)

	messagesMongol := CheckMessages("Mongol", passMongol, deviceMongol, keyMongol)
	fmt.Println("MEssages Mongol: " + messagesMongol)

	messagesMongolSplitted := strings.Split(messagesMongol, "\\\\")
	fmt.Println("mes splitted mongol: ", messagesMongolSplitted[1][2:])
	endInit := EndInitDialog(privKeyMongol, messagesMongolSplitted[1][2:], "Mongoliseker")
	fmt.Println(endInit)
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
func SearchUser(login, deviceInfo, loginForSearch, key string, numSeaerchRequest int) string {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	loginForSearchEncrypt, err := crypto.StringEncrypt([]byte(loginForSearch), key)
	if err != nil {
		return "$ Необходимо заново авторизоваться"
	}

	formData := url.Values{}
	formData.Add("device", hashedDeviceInfo)
	formData.Add("loginforsearch", loginForSearchEncrypt)
	formData.Add("forserver", crypto.GenKey(deviceInfo))

	resp, err := client.PostForm(forServerReq+mainServicePort+"/searchuser", formData)
	if err != nil {
		if os.IsTimeout(err) {
			return "$ Сервер не отвечает по причине долгого ответа"
		}
		return "$ Сервер не отвечает по какому то дерьму"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "$ Не удалось прочитать ответ от сервера"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		body, err = crypto.StringDecryptByte(body, key)
		if err != nil {
			fmt.Println(err)
		}

		var findUsers []FindUser
		var out string
		if err := json.Unmarshal(body, &findUsers); err != nil {
			return "$ ошибка обработки ответа, попробуй еще раз " + err.Error()
		}

		// out += strconv.Itoa(resp.NumSeaerchRequest)

		fmt.Println(findUsers)
		for findUser := range findUsers {
			if findUsers[findUser].Login == login {
				continue
			}

			if out != "" {
				out += "\n"
			}

			out += findUsers[findUser].Login + " "

			out += findUsers[findUser].Name
		}

		return out
	case http.StatusInternalServerError:
		decryptedResp, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ ошибка чтения ответа" + err.Error()
		}
		return "$ " + decryptedResp
	case http.StatusBadRequest:
		return "$ " + string(body)
	default:
		return "$ " + string(body)
	}
}

//export SendMessage
func SendMessage(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key string) string {
	resp := sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key)
	return resp
}

func sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, message, key string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	senderLoginEncrypted, err := crypto.StringEncrypt([]byte(senderLogin), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	receiverLoginEncrypted, err := crypto.StringEncrypt([]byte(receiverLogin), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	messageEncrypted, err := crypto.StringEncrypt([]byte(message), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	senderPasswordEncrypted, err := crypto.StringEncrypt([]byte(senderPassword), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, senderLogin)

	formData := url.Values{}
	formData.Add("senderlogin", senderLoginEncrypted)
	formData.Add("senderpassword", senderPasswordEncrypted)
	formData.Add("receiverlogin", receiverLoginEncrypted)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("message", messageEncrypted)
	formData.Add("forserver", crypto.GenKey(deviceInfo))
	// if !(message[0] == designationOfAnswerDialog[0] || message[0] == designationOfInitDialog[0]) {
	// 	formData.Add("senderLoginHashed", crypto.HashLogin(senderLogin, dialogKey))
	// }

	resp, err := client.PostForm(forServerReq+mainServicePort+"/sendmessage", formData)
	if err != nil {
		if os.IsTimeout(err) {
			return "$ Сервер не отвечает"
		}
		return "$ Не удается установить соединение с сервером"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "$ Не удается обработать ответ сервера"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		if string(body) == "" {
			return "Сообщение отправлено"
		}
		respDecrypted, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ ошибка расшифровки"
		}
		return respDecrypted
	case http.StatusBadRequest:
		return "$ Проверьте правильность ввода данных"
	case http.StatusForbidden:
		return "$ Не правильный логин или пароль"
	case http.StatusInternalServerError:
		resp, _ := crypto.StringDecrypt(string(body), key)
		return "$ Внутренняя ошибка сервера:" + resp
	default:
		return "$ " + string(body)
	}
}

//export CheckMessages
func CheckMessages(login, password, deviceInfo, key string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	loginEcnrypted, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	passwordEncrypted, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	formData := url.Values{}
	formData.Add("login", loginEcnrypted)
	formData.Add("password", passwordEncrypted)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("forserver", crypto.GenKey(deviceInfo))

	resp, err := client.PostForm(forServerReq+mainServicePort+"/checkmessage", formData)
	if err != nil {
		if os.IsTimeout(err) {
			return "$ Сервер не отвечает"
		}
		return "$ Не удается установить соединение с сервером"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "$ Не удается обработать ответ сервера"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		var response RequestCheckMessages
		if err := json.Unmarshal(body, &response); err != nil {
			return "$ Некорректный формат ответа от сервера"
		}
		decryptedResp, err := crypto.StringDecryptByte(response.Messages, key)
		if err != nil {
			return "$ Ошибка расшифровки ответа"
		}

		if response.Gzip {
			decryptedResp, err = DecompressGzip(decryptedResp)
			if err != nil {
				return "$ не удалось распаковать сообщения"
			}
		}

		lines := strings.Split(string(decryptedResp), "\n")
		var result strings.Builder

		for li, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			var msg Message
			if err := json.Unmarshal([]byte(line), &msg); err != nil {
				return "$ ошибка обработки ответа, попробуй еще раз " + err.Error()
			}

			date := msg.SendTime.Format("02.01.2006 15:04")
			_, offsetSeconds := msg.SendTime.Zone()
			offsetHours := offsetSeconds / 3600
			offsetShort := fmt.Sprintf("%+03d", offsetHours)

			result.WriteString(msg.Sender)
			result.WriteString("\\\\")
			result.WriteString(msg.Message)
			result.WriteString("\\\\")
			result.WriteString(date)
			result.WriteString(" ")
			result.WriteString(offsetShort)

			// добавляем перевод строки между записями
			if li < len(lines)-1 {
				result.WriteString("\n")
			}
		}

		out := result.String()

		return out
	case http.StatusBadRequest:
		return "$ Проверьте правильность ввода данных"
	case http.StatusForbidden:
		return "$ Не правильный логин или пароль"
	case http.StatusInternalServerError:
		resp, _ := crypto.StringDecrypt(string(body), key)
		return "$ Внутренняя ошибка сервера" + resp
	default:
		return "$ Неизвестная ошибка"
	}
}

//export DelMessages
func DelMessages(login, password, deviceInfo, key string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	loginEcnrypted, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	passwordEncrypted, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	formData := url.Values{}
	formData.Add("login", loginEcnrypted)
	formData.Add("password", passwordEncrypted)
	formData.Add("device", hashedDeviceInfo)
	formData.Add("forserver", crypto.GenKey(deviceInfo))

	resp, err := client.PostForm(forServerReq+mediumSizeDataServicePort+"/delmessages", formData)
	if err != nil {
		if os.IsTimeout(err) {
			return "$ Сервер не отвечает"
		}
		return "$ Не удается установить соединение с сервером"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "$ Не удается обработать ответ сервера"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return "Успешно удалено"
	case http.StatusBadRequest:
		resp, err := crypto.StringDecrypt(string(body), key)
		if err != nil {
			return "$ ошибка при расшифровке " + err.Error()
		}
		return "$ Проверьте правильность ввода данных " + resp
	case http.StatusForbidden:
		return "$ Не правильный логин или пароль" + string(body)
	case http.StatusInternalServerError:
		return "$ Внутренняя ошибка сервера" + string(body)
	default:
		return "$ Неизвестная ошибка" + string(body)
	}
}

func initiatorKeyGeneration() (string, string, error) {
	curve := ecdh.P256()

	privKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}
	pubKey := privKey.PublicKey()

	// сериализуем приватный ключ в base64
	privKeyBase64 := base64.StdEncoding.EncodeToString(privKey.Bytes())
	// сериализуем публичный ключ в base64
	pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKey.Bytes())

	return privKeyBase64, pubKeyBase64, nil
}

//export InitDialog
func InitDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, key string) string {
	privKeyBase64, pubKeyBase64, err := initiatorKeyGeneration()
	if err != nil {
		return "Произошла ошибка при генерации ключей"
	}

	resp := sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, designationOfInitDialog+"\\"+pubKeyBase64, key)
	if resp == "Сообщение отправлено" {
		return privKeyBase64
	}
	return "$ " + resp
}

func answerKeyGeneration(user1PubBase64 string) (string, string) {
	curve := ecdh.P256()

	// генерируем ключи юзера2
	privKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("User2: ошибка генерации ключа: %v", err)
	}
	pubKey := privKey.PublicKey()

	// декодируем публичный ключ юзера1
	user1PubBytes, err := base64.StdEncoding.DecodeString(user1PubBase64)
	if err != nil {
		log.Fatalf("User2: ошибка декодирования публичного ключа юзера1: %v", err)
	}
	user1PubKey, err := curve.NewPublicKey(user1PubBytes)
	if err != nil {
		log.Fatalf("User2: ошибка восстановления публичного ключа юзера1: %v", err)
	}

	// вычисляем общий секрет
	sharedSecret, err := privKey.ECDH(user1PubKey)
	if err != nil {
		log.Fatalf("User2: ошибка вычисления общего секрета: %v", err)
	}

	// сериализуем публичный ключ юзера2
	pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKey.Bytes())
	return pubKeyBase64, hex.EncodeToString(sharedSecret)
}

//export AnswerInitDialog
func AnswerInitDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, pubKeyBase64, key string) string {
	pubKeyBase64ForSend, sharedSecret := answerKeyGeneration(pubKeyBase64)

	resp := sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, designationOfAnswerDialog+"\\"+pubKeyBase64ForSend, key)
	if resp == "Сообщение отправлено" {
		return sharedSecret + "\\" + crypto.HashLogin(receiverLogin, sharedSecret)
	}
	return "$ " + resp

}

//export EndInitDialog
func EndInitDialog(privBase64, user2PubBase64, loginCommunicating string) string {
	curve := ecdh.P256()

	// восстанавливаем приватный ключ юзера1
	privBytes, err := base64.StdEncoding.DecodeString(privBase64)
	if err != nil {
		log.Fatalf("User1: ошибка декодирования приватного ключа: %v", err)
	}
	privKey, err := curve.NewPrivateKey(privBytes)
	if err != nil {
		log.Fatalf("User1: ошибка восстановления приватного ключа: %v", err)
	}

	// восстанавливаем публичный ключ юзера2
	pubBytes, err := base64.StdEncoding.DecodeString(user2PubBase64)
	if err != nil {
		log.Fatalf("User1: ошибка декодирования публичного ключа юзера2: %v", err)
	}
	pubKey, err := curve.NewPublicKey(pubBytes)
	if err != nil {
		log.Fatalf("User1: ошибка восстановления публичного ключа юзера2: %v", err)
	}

	// вычисляем общий секрет
	sharedSecret, err := privKey.ECDH(pubKey)
	if err != nil {
		log.Fatalf("User1: ошибка вычисления общего секрета: %v", err)
	}
	return hex.EncodeToString(sharedSecret) + "\\" + crypto.HashLogin(loginCommunicating, hex.EncodeToString(sharedSecret))
}

//export SetAvatar
func SetAvatar(login, password, deviceInfo, key string, avatar []byte) string {
	loginEcnrypted, err := crypto.StringEncrypt([]byte(login), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	passwordEncrypted, err := crypto.StringEncrypt([]byte(password), key)
	if err != nil {
		return "$ ошибка зашифровки"
	}
	avatarEncrypted, err := crypto.StringEncryptByte(avatar, key)
	if err != nil {
		return "$ ошибка зашифровки 1"
	}
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)
	fmt.Println("hashedDeviceInfo ", hashedDeviceInfo)

	payload := Avatar{
		Avatar:               avatarEncrypted,
		Login:                loginEcnrypted,
		Password:             passwordEncrypted,
		DeviceInfo:           hashedDeviceInfo,
		KeyForServerDataBase: crypto.GenKey(deviceInfo),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "$ Ошибка при маршалинге JSON:" + err.Error()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, forServerReq+mediumSizeDataServicePort+"/setavatar", bytes.NewReader(body))
	if err != nil {
		return "$ Ошибка при создании запроса:" + err.Error()
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "$ Сервер не ответил в течение 10 секунд. Завершение программы."
		}
		return "$ Ошибка при выполнении запроса:" + err.Error()
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return "Аватар усшпено загружен"
	case http.StatusBadRequest:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "Не удается обработать ответ сервера"
		}
		return "$ " + string(body)
	case http.StatusForbidden:
		return "$ Не правильный логин или пароль"
	case http.StatusInternalServerError:
		resp, _ := crypto.StringDecrypt(string(body), key)
		return "$ Внутренняя ошибка сервера" + resp
	default:
		return "$ Неизвестная ошибка"
	}
}

//export GetAvatar
func GetAvatar(login, loginForSearch, deviceInfo, key string) []byte {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	loginEcnrypted, err := crypto.StringEncrypt([]byte(loginForSearch), key)
	if err != nil {
		return []byte{}
	}
	hashedDeviceInfo := crypto.HashDeviceInfo(deviceInfo, login)

	formData := url.Values{}
	formData.Add("loginforsearch", loginEcnrypted)
	formData.Add("device", hashedDeviceInfo)
	fmt.Println("hashedDeviceInfo ", hashedDeviceInfo)
	formData.Add("forserver", crypto.GenKey(deviceInfo))

	resp, err := client.PostForm(forServerReq+mediumSizeDataServicePort+"/getavatar", formData)
	if err != nil {
		fmt.Println(err)
		if os.IsTimeout(err) {
			return []byte{}
		}
		return []byte{}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	fmt.Println("resp ", string(respBody))
	fmt.Println(resp.StatusCode)
	switch resp.StatusCode {
	case http.StatusOK:
		var encAvatar []byte
		if err := json.Unmarshal(respBody, &encAvatar); err != nil {
			// Если не получилось, возможно сервер вернул JSON-строку или plain string с зашифрованными байтами
			var s string
			if err2 := json.Unmarshal(respBody, &s); err2 == nil {
				encAvatar = []byte(s)
			} else {
				// как запасной вариант — попробовать использовать сырое тело без JSON
				encAvatar = respBody
			}
		}

		// Расшифровываем полученные байты
		decAvatar, err := crypto.StringDecryptByte(encAvatar, key)
		if err != nil {
			fmt.Println("decrypt avatar:", err)
			return nil
		}
		return decAvatar
	case http.StatusNoContent:
		return []byte{}
	case http.StatusBadRequest:
		return []byte{}
	case http.StatusInternalServerError:
		return []byte{}
	default:
		return []byte{}
	}
}

//export InitAudioDialog
func InitAudioDialog(senderLogin, senderPassword, receiverLogin, deviceInfo, key string) string {
	return sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, "a\\", key)
}

//export StopAudioDialog
func StopAudioDialog(senderLogin, senderPassword, receiverLogin, dialogUUID, deviceInfo, key string) string {
	return sendMessageToServer(senderLogin, senderPassword, receiverLogin, deviceInfo, fmt.Sprintf("as\\%v", dialogUUID), key)
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, forServerReq+audioDialogServicePort+"/audiodialog", bytes.NewReader(body))
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
		return "$ Ошибка сервера"
	}

	var data map[string]map[string][]byte
	if err := json.Unmarshal(raw, &data); err != nil {
		return "$ Ошибка парсинга ответа"
	}

	count := 1
	result := ""
	for user, msgs := range data {
		if user == UserID {
			continue
		}
		for _, msg := range msgs {
			dec, err := crypto.StringDecryptByte(msg, key)
			if err == nil {
				// result[count] = dec
				if len(result) == 0 {
					result = hex.EncodeToString(dec)
				} else {
					result += "\\" + hex.EncodeToString(dec)
				}
				count++
			}
		}
	}

	return result
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

func CompressGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecompressGzip(compressed []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r)
}
