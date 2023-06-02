package uishell

import (
	"fmt"
	"os"
	"strings"

	"github.com/BillyBones007/pwdm_client/internal/customerror"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	"github.com/BillyBones007/pwdm_client/internal/tools/filetools"
	"github.com/BillyBones007/pwdm_client/internal/tools/tablesgen"
	"github.com/c-bata/go-prompt"
	"golang.org/x/term"
)

func (s *ShellUI) executor(in string) {
	in = strings.TrimSpace(in)
	commands := strings.Split(in, " ")
	switch commands[0] {
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
			bytePwd, err := term.ReadPassword(0)
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
			bytePwd, err := term.ReadPassword(0)
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
		s.Prefix.User = user + " >>> "
		s.Prefix.IsEnable = true

	case "logout":
		s.logOutUser()
		s.Prefix.User = ">>> "
		s.Prefix.IsEnable = false
		return

	case "delete":
		confirmation := prompt.Input("Confirmation of deletion (y/n) > ", s.Completer.Complete)
		if confirmation == "n" {
			fmt.Println("***")
			fmt.Println("Deletion canceled")
			fmt.Println("***")
			return
		} else if confirmation == "y" {
			err := s.delRecord(in)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("***")
			fmt.Println("Record is deleted")
			fmt.Println("***")
			return
		}
		fmt.Println(customerror.ErrInvalildCommand.Error())
		return

	case "show":
		if len(commands) != 2 {
			fmt.Println(customerror.ErrInvalildCommand.Error())
			return
		}
		table, err := s.getDataTable(commands[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		if table == "" {
			fmt.Println("No information")
			return
		}
		fmt.Println("***")
		fmt.Printf("Information on request:\n\n%s\n\n", table)
		fmt.Println("***")
		fmt.Println()
		return

	case "save":
		if len(commands) == 2 {
			err := s.saveDialogue(commands[1])
			if err != nil {
				fmt.Println(err)
				return
			}
		} else if len(commands) > 2 {
			err := s.saveDialogue(commands[2])
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	case "get":
		model, err := s.getRecord(in)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch model := model.(type) {
		case models.LogPwdModel:
			table := tablesgen.NewLogPwdTable(model)
			fmt.Println("***")
			fmt.Printf("Information on request:\n\n%s\n\n", table)
			fmt.Println("***")
			fmt.Println()
			return
		case models.CardModel:
			table := tablesgen.NewCardTable(model)
			fmt.Println("***")
			fmt.Printf("Information on request:\n\n%s\n\n", table)
			fmt.Println("***")
			fmt.Println()
			return
		case models.TextDataModel:
			fileName := prompt.Input("Enter text file name to save > ", s.Completer.Complete)
			err := filetools.WriteTextFile(model.Data, fileName)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("\n***")
			fmt.Printf("\nFile `%s` is saved\n\n", fileName)
			fmt.Println("***")
			return
		case models.BinaryDataModel:
			fileName := prompt.Input("Enter binary file name to save > ", s.Completer.Complete)
			err := filetools.WriteBinaryFile(model.Data, fileName)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("\n***")
			fmt.Printf("\nFile `%s` is saved\n\n", fileName)
			fmt.Println("***")
			return
		}
		return

	case "edit":
		confirmation := prompt.Input("Edit the selected record? (y/n) > ", s.Completer.Complete)
		if confirmation == "n" {
			fmt.Println("***")
			fmt.Println("Edit canceled")
			fmt.Println("***")
			return
		} else if confirmation == "y" {
			titleRec, err := s.editRecordDialogue(in)
			if err != nil {
				fmt.Println("\n***")
				fmt.Println(err)
				fmt.Println("Error update record!")
				fmt.Println("***")
			}
			fmt.Println("\n***")
			fmt.Printf("\nRecord `%s` has been update\n\n", titleRec)
			fmt.Println("***")
			return
		}

	default:
		return
	}
}
