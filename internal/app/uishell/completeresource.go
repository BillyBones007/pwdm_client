package uishell

import (
	"fmt"

	"github.com/BillyBones007/pwdm_client/internal/app/grpcclient"
	"github.com/BillyBones007/pwdm_client/internal/datatypes"
	"github.com/c-bata/go-prompt"
)

// Commands and subcommands for completer.
var (
	authUserCommands = []prompt.Suggest{
		{Text: "logout", Description: "Log out of your account"},
		{Text: "exit", Description: "Exit this program"},
		{Text: "show", Description: "Show user saved data"},
		{Text: "save", Description: "Save user data on the server"},
		{Text: "get", Description: "Get user data from server"},
		{Text: "delete", Description: "Delete user data from server"},
	}

	guestUserCommands = []prompt.Suggest{
		{Text: "login", Description: "Log in to your account"},
		{Text: "register", Description: "Add a new user"},
		{Text: "exit", Description: "Exit this program"},
		{Text: "help", Description: "Help information"},
	}

	saveSubcommands = []prompt.Suggest{
		{Text: "lp", Description: "Save login/password data on the server"},
		{Text: "card", Description: "Save bank cards data on the server"},
		{Text: "text", Description: "Save text data on the server"},
		{Text: "binary", Description: "Save binary data on the server"},
	}

	saveTextSubcommands = []prompt.Suggest{
		{Text: "-f", Description: "Save text data from file on the server"},
	}

	getSubcommands = []prompt.Suggest{
		{Text: "lp", Description: "Get login/password data from server"},
		{Text: "card", Description: "Get bank cards data from server"},
		{Text: "text", Description: "Get text data from server"},
		{Text: "binary", Description: "Get binary data from server"},
	}

	showSubcommands = []prompt.Suggest{
		{Text: "lp", Description: "Show saved login/password data"},
		{Text: "card", Description: "Show saved bank cards data"},
		{Text: "text", Description: "Show saved text data"},
		{Text: "binary", Description: "Show saved binary data"},
	}

	deleteSubcommands = []prompt.Suggest{
		{Text: "lp", Description: "Select login/password data to delete"},
		{Text: "card", Description: "Select bank cards data to delete"},
		{Text: "text", Description: "Select text data to delete"},
		{Text: "binary", Description: "Select binary data to delete"},
	}
)

// getLogPwdInfo- returns login/password data list of the current user
func getLogPwdInfo(client *grpcclient.ClientGRPC) []prompt.Suggest {
	if client.AuthFlag {
		dataAll, err := client.UpdateInfo()
		if err != nil {
			return []prompt.Suggest{}
		}
		client.Config.Storage.Clear(datatypes.LoginPasswordDataType)
		client.Config.Storage.UpdateStorage(datatypes.LoginPasswordDataType, dataAll)
		data := client.Config.Storage.GetListRecords(datatypes.LoginPasswordDataType)
		s := make([]prompt.Suggest, len(data))
		i := 0
		for k, v := range data {
			s[i] = prompt.Suggest{
				Text: fmt.Sprint(k), Description: v.Title,
			}
			i++
		}
		return s
	}
	return []prompt.Suggest{}
}

// getCardInfo - returns bank cards data list of the current user
func getCardInfo(client *grpcclient.ClientGRPC) []prompt.Suggest {
	if client.AuthFlag {
		dataAll, err := client.UpdateInfo()
		if err != nil {
			return []prompt.Suggest{}
		}
		client.Config.Storage.Clear(datatypes.CardDataType)
		client.Config.Storage.UpdateStorage(datatypes.CardDataType, dataAll)
		data := client.Config.Storage.GetListRecords(datatypes.CardDataType)
		s := make([]prompt.Suggest, len(data))
		i := 0
		for k, v := range data {
			s[i] = prompt.Suggest{
				Text: fmt.Sprint(k), Description: v.Title,
			}
			i++
		}
		return s
	}
	return []prompt.Suggest{}
}

// getTextInfo - returns text data list of the current user
func getTextInfo(client *grpcclient.ClientGRPC) []prompt.Suggest {
	if client.AuthFlag {
		dataAll, err := client.UpdateInfo()
		if err != nil {
			return []prompt.Suggest{}
		}
		client.Config.Storage.Clear(datatypes.TextDataType)
		client.Config.Storage.UpdateStorage(datatypes.TextDataType, dataAll)
		data := client.Config.Storage.GetListRecords(datatypes.TextDataType)
		s := make([]prompt.Suggest, len(data))
		i := 0
		for k, v := range data {
			s[i] = prompt.Suggest{
				Text: fmt.Sprint(k), Description: v.Title,
			}
			i++
		}
		return s
	}
	return []prompt.Suggest{}
}

// getBinaryInfo - returns binary data list of the current user
func getBinaryInfo(client *grpcclient.ClientGRPC) []prompt.Suggest {
	if client.AuthFlag {
		dataAll, err := client.UpdateInfo()
		if err != nil {
			return []prompt.Suggest{}
		}
		client.Config.Storage.Clear(datatypes.BinaryDataType)
		client.Config.Storage.UpdateStorage(datatypes.BinaryDataType, dataAll)
		data := client.Config.Storage.GetListRecords(datatypes.BinaryDataType)
		s := make([]prompt.Suggest, len(data))
		i := 0
		for k, v := range data {
			s[i] = prompt.Suggest{
				Text: fmt.Sprint(k), Description: v.Title,
			}
			i++
		}
		return s
	}
	return []prompt.Suggest{}
}
