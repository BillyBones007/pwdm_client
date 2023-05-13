package main

import (
	"fmt"

	"github.com/BillyBones007/pwdm_client/internal/app/grpcclient"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
)

func main() {
	// TODO: запускается отдельная горутина, которая опрашивает сервер
	// через определенный промежуток времени на наличие новой информации в базе.
	// ui.UserInterface()
	client := grpcclient.InitClient()
	user := models.UserModel{Login: "Pasha", Password: "1234"}

	// mylogin, err := client.Registration(user)
	// if err != nil {
	// 	fmt.Printf("ERROR: %s\n", err)
	// 	return
	// }
	mylogin, err := client.LogIn(user)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	fmt.Printf("Login %s\n", mylogin)
	fmt.Printf("Token %s\n", client.Token)

	// data := models.LogPwdModel{
	// 	Login: "Hren", Password: "1slfkj2",
	// 	TechData: models.ReqTechDataModel{
	// 		Title:   "Second Login/Password pair",
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
	data, err := client.UpdateInfo()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
	for _, v := range data {
		fmt.Println(v)
	}

}
