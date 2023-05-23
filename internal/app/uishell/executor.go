package uishell

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BillyBones007/pwdm_client/internal/datatypes"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	"github.com/c-bata/go-prompt"
	"golang.org/x/crypto/ssh/terminal"
)

func (s *ShellUI) executor(in string) {
	in = strings.TrimSpace(in)
	setCommand := strings.Split(in, " ")
	switch setCommand[0] {
	case "help":
		fmt.Println(helpText)
		return

	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
		return

	case "login":
		if s.Client.AuthFlag {
			s.logOutUser()
			s.Prefix.User = ">>> "
			s.Prefix.IsEnable = false
		}

		login := prompt.Input("Enter login > ", s.Completer.Complete)
		password := func(label string) string {
			fmt.Print(label)
			bytePwd, err := terminal.ReadPassword(0)
			if err != nil {
				return ""
			}
			pwd := strings.TrimSpace(string(bytePwd))
			fmt.Println()
			return pwd
		}("Enter password > ")

		userModel := models.UserModel{Login: login, Password: password}
		user, err := s.Client.LogIn(userModel)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		// go updateStorageData(s.Client)
		s.Prefix.User = user + " >>> "
		s.Prefix.IsEnable = true
		return

	case "register":
		if s.Client.AuthFlag {
			s.logOutUser()
			s.Prefix.User = ">>> "
			s.Prefix.IsEnable = false
		}

		login := prompt.Input("Enter login > ", s.Completer.Complete)
		password := func(label string) string {
			fmt.Print(label)
			bytePwd, err := terminal.ReadPassword(0)
			if err != nil {
				return ""
			}
			pwd := strings.TrimSpace(string(bytePwd))
			fmt.Println()
			return pwd
		}("Enter password > ")

		userModel := models.UserModel{Login: login, Password: password}
		user, err := s.Client.Registration(userModel)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		// go updateStorageData(s.Client)
		s.Prefix.User = user + " >>> "
		s.Prefix.IsEnable = true

	case "logout":
		s.logOutUser()
		s.Prefix.User = ">>> "
		s.Prefix.IsEnable = false
		return

	case "delete":
		err := s.delRecord(in)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Record is deleted")
		return

	default:
		return
	}
}

// delRecord - delete current record from server.
func (s *ShellUI) delRecord(in string) error {
	in = strings.TrimSpace(in)
	commands := strings.Split(in, " ")
	if len(commands) == 3 {
		switch commands[0] {
		case "delete":
			switch commands[1] {
			case "lp":
				key, err := strconv.Atoi(commands[2])
				if err != nil {
					return err
				}
				err = s.Client.DeleteRecord(key, datatypes.LoginPasswordDataType)
				if err != nil {
					return err
				}
				return nil
			case "card":
				key, err := strconv.Atoi(commands[2])
				if err != nil {
					return err
				}
				err = s.Client.DeleteRecord(key, datatypes.CardDataType)
				if err != nil {
					return err
				}
				return nil
			case "text":
				key, err := strconv.Atoi(commands[2])
				if err != nil {
					return err
				}
				err = s.Client.DeleteRecord(key, datatypes.TextDataType)
				if err != nil {
					return err
				}
				return nil
			case "binary":
				key, err := strconv.Atoi(commands[2])
				if err != nil {
					return err
				}
				err = s.Client.DeleteRecord(key, datatypes.BinaryDataType)
				if err != nil {
					return err
				}
				return nil

			}
		default:
			return errors.New("invalid command")
		}
	}
	return errors.New("the parameters must be at least 3")
}

// logOutUser - clears user data when logging out of the account.
func (s *ShellUI) logOutUser() {
	s.Client.Token = ""
	s.Client.AuthFlag = false
	s.Client.Config.Storage.Clear(datatypes.LoginPasswordDataType)
	s.Client.Config.Storage.Clear(datatypes.CardDataType)
	s.Client.Config.Storage.Clear(datatypes.TextDataType)
	s.Client.Config.Storage.Clear(datatypes.BinaryDataType)
}
