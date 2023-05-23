package main

import (
	// "bytes"
	// "crypto/rand"
	// "fmt"

	// "github.com/BillyBones007/pwdm_client/internal/app/grpcclient"
	// "github.com/BillyBones007/pwdm_client/internal/datatypes"
	// "github.com/BillyBones007/pwdm_client/internal/storage/models"
	// "github.com/BillyBones007/pwdm_client/internal/app/ui"
	"github.com/BillyBones007/pwdm_client/internal/app/uishell"
)

// Global variables for the linker.
var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
)

func main() {
	appData := uishell.AppData{BuildVersion: buildVersion, BuildDate: buildDate}
	app := uishell.NewShellUI(appData)
	app.RunShell()
	// buildInfo := ui.BuildInfo{Version: buildVersion, Date: buildDate}
	// ui.StartUI(buildInfo)
	// TODO: запускается отдельная горутина, которая опрашивает сервер
	// через определенный промежуток времени на наличие новой информации в базе.
	// ui.UserInterface()
	// client := grpcclient.InitClient()
	// user := models.UserModel{Login: "Pasha", Password: "1234"}

	// mylogin, err := client.Registration(user)
	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// 	return
	// }
	// mylogin, err := client.LogIn(user)
	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// 	return
	// }
	// fmt.Printf("Login %s\n", mylogin)
	// fmt.Printf("Token %s\n", client.Token)

	// data := models.LogPwdModel{
	// 	Login: "Bazilio", Password: "1slfkj2",
	// 	TechData: models.TechDataModel{
	// 		Title:   "Bazilio",
	// 		Type:    datatypes.LoginPasswordDataType,
	// 		Tag:     "Test second tag",
	// 		Comment: "Test second comment",
	// 	},
	// }
	// resp, err := client.SendLogPwd(data)
	// if err != nil {
	// 	fmt.Printf("ERROR: Error from err: %s\n", err)
	// 	fmt.Printf("ERROR: Error from resp: %s\n", resp)
	// }
	// fmt.Printf("INFO: Info from resp - Title: %s\n", resp)

	// -----------------------------------------------------------------
	// cardData := models.CardModel{
	// 	Num:       "1234 1234 1234 1234",
	// 	Date:      "04/25",
	// 	CVC:       "123",
	// 	FirstName: "Ivan",
	// 	LastName:  "Ivanov",
	// 	TechData: models.TechDataModel{
	// 		Title:   "Ivanov card",
	// 		Tag:     "ivanovcard",
	// 		Comment: "Some card",
	// 		Type:    datatypes.CardDataType,
	// 	},
	// }
	// respCard, err := client.SendCard(cardData)
	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// }
	// fmt.Printf("INFO: Send card data: %s\n", respCard)

	// -----------------------------------------------------------------
	// 	text := `Берёза

	// Белая берёза
	// Под моим окном
	// Принакрылась снегом,
	// Точно серебром.

	// На пушистых ветках
	// Снежною каймой
	// Распустились кисти
	// Белой бахромой.

	// И стоит берёза
	// В сонной тишине,
	// И горят снежинки
	// В золотом огне.

	// А заря, лениво
	// Обходя кругом,
	// Обсыпает ветки
	// Новым серебром.
	// 1913 г.`
	// 	textData := models.TextDataModel{
	// 		Data: text,
	// 		TechData: models.TechDataModel{
	// 			Title: "Esenin",
	// 			Tag:   "poems",
	// 			Type:  datatypes.TextDataType,
	// 		},
	// 	}
	// 	respText, err := client.SendText(textData)
	// 	if err != nil {
	// 		fmt.Printf("ERROR: %s\n", err)
	// 	}
	// 	fmt.Printf("INFO: Send text data: %s\n", respText)
	// -----------------------------------------------------------------

	// binaryData, _ := helperRandBytes(1024 * 1024 * 3)
	// m := models.BinaryDataModel{
	// 	Data: binaryData,
	// 	TechData: models.TechDataModel{
	// 		Title: "Binary data",
	// 		Tag:   "binarytag",
	// 		Type:  datatypes.BinaryDataType,
	// 	},
	// }

	// respB, err := client.SendBinary(m)
	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// }
	// fmt.Printf("INFO: Send binary data: %s\n", respB)

	// -----------------------------------------------------------------
	// dataAll, err := client.UpdateInfo()
	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// }
	// client.Storage.UpdateStorage(datatypes.LoginPasswordDataType, dataAll)
	// client.Storage.UpdateStorage(datatypes.CardDataType, dataAll)
	// client.Storage.UpdateStorage(datatypes.TextDataType, dataAll)
	// client.Storage.UpdateStorage(datatypes.BinaryDataType, dataAll)
	// fmt.Printf("INFO: Updating data %v+\n", dataAll)
	// -----------------------------------------------------------------
	// resp, err := client.GetLogPwd(1)
	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// }
	// fmt.Printf("Login: %s\n", resp.Login)
	// fmt.Printf("Password: %s\n", resp.Password)
	// fmt.Printf("Title: %s\n", resp.TechData.Title)
	// fmt.Printf("Tag: %s\n", resp.TechData.Tag)
	// fmt.Printf("Comment: %s\n", resp.TechData.Comment)

	// -----------------------------------------------------------------
	// resp, err := client.GetCard(1)
	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// }
	// fmt.Printf("Num card: %s\n", resp.Num)
	// fmt.Printf("Date card: %s\n", resp.Date)
	// fmt.Printf("CVC card: %s\n", resp.CVC)
	// fmt.Printf("First name: %s\n", resp.FirstName)
	// fmt.Printf("Last name: %s\n", resp.LastName)
	// fmt.Printf("Title record: %s\n", resp.TechData.Title)
	// fmt.Printf("Tag record: %s\n", resp.TechData.Tag)
	// fmt.Printf("Comment: %s\n", resp.TechData.Comment)
	// -----------------------------------------------------------------
	// resp, err := client.GetText(1)
	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// }
	// fmt.Printf("Title: %s\n", resp.TechData.Title)
	// fmt.Printf("Tag: %s\n", resp.TechData.Tag)
	// fmt.Printf("Data: %s\n", resp.Data)
	// -----------------------------------------------------------------
	// resp, err := client.GetBinary(1)
	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// }
	// fmt.Printf("Title: %s\n", resp.TechData.Title)
	// fmt.Printf("Tag: %s\n", resp.TechData.Tag)
	// if bytes.Equal(resp.Data, binaryData) {
	// 	fmt.Println("Recieved data is equal to the send data")

	// }

	// -----------------------------------------------------------------

}

// helperRandBytes - helper function.
// Uses in tests.
// func helperRandBytes(size int) ([]byte, error) {
// 	b := make([]byte, size)
// 	_, err := rand.Read(b)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return b, nil
// }
