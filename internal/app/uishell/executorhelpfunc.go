package uishell

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/BillyBones007/pwdm_client/internal/customerror"
	"github.com/BillyBones007/pwdm_client/internal/datatypes"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	"github.com/BillyBones007/pwdm_client/internal/tools/filetools"
	"github.com/BillyBones007/pwdm_client/internal/tools/tablesgen"
	"github.com/BillyBones007/pwdm_client/internal/tools/validcard"
	"github.com/c-bata/go-prompt"
)

// saveDialogue - displays a dialog depending on the data type being saved.
func (s *ShellUI) saveDialogue(command string) error {
	switch command {
	case "lp":
		data := models.LogPwdModel{}
		fmt.Println("***")
		fmt.Println("Enter Title, Comment, Tag, Login and Password to save:")
		fmt.Println(hint)
		fmt.Println("***")
		data.TechData.Title = prompt.Input("Enter title > ", s.Completer.Complete)
		data.TechData.Comment = prompt.Input("Enter comment > ", s.Completer.Complete)
		data.TechData.Tag = prompt.Input("Enter tag > ", s.Completer.Complete)
		data.TechData.Type = datatypes.LoginPasswordDataType
		data.Login = prompt.Input("Enter login > ", s.Completer.Complete)
		data.Password = prompt.Input("Enter password > ", s.Completer.Complete)
		if data.Login == "" || data.Password == "" {
			return customerror.ErrEmptyFields
		}
		resp, err := s.saveLogPwd(data)
		if err != nil {
			return err
		}
		fmt.Println()
		fmt.Printf("Record `%s` is saved\n", resp)
		fmt.Println("***")
		fmt.Println()
		return nil

	case "card":
		data := models.CardModel{}
		fmt.Println("***")
		fmt.Println("Enter Title, Comment, Tag and bank card data to save:")
		fmt.Println(hint)
		fmt.Println("***")
		data.TechData.Title = prompt.Input("Enter title > ", s.Completer.Complete)
		data.TechData.Comment = prompt.Input("Enter comment > ", s.Completer.Complete)
		data.TechData.Tag = prompt.Input("Enter tag > ", s.Completer.Complete)
		data.TechData.Type = datatypes.CardDataType
		data.Num = prompt.Input("Enter the card number > ", s.Completer.Complete)
		data.CVC = prompt.Input("Enter the CVC number > ", s.Completer.Complete)
		data.Date = prompt.Input("Enter the expiration date > ", s.Completer.Complete)
		data.FirstName = prompt.Input("Enter the First name > ", s.Completer.Complete)
		data.LastName = prompt.Input("Enter the Last name > ", s.Completer.Complete)
		if data.Num == "" || data.CVC == "" || data.Date == "" || data.FirstName == "" || data.LastName == "" {
			return customerror.ErrEmptyFields
		}
		err := validcard.CnValidator(data.Num)
		if err != nil {
			return err
		}
		err = validcard.CVCValidator(data.CVC)
		if err != nil {
			return err
		}
		err = validcard.ExpValidator(data.Date)
		if err != nil {
			return err
		}

		resp, err := s.saveCard(data)
		if err != nil {
			return err
		}
		fmt.Println()
		fmt.Printf("Record `%s` is saved\n", resp)
		fmt.Println("***")
		fmt.Println()
		return nil
	case "text":
		data := models.TextDataModel{}
		fmt.Println("***")
		fmt.Println("Enter Title, Comment, Tag and some text to save:")
		fmt.Println(hint)
		fmt.Println("***")
		data.TechData.Title = prompt.Input("Enter title > ", s.Completer.Complete)
		data.TechData.Comment = prompt.Input("Enter comment > ", s.Completer.Complete)
		data.TechData.Tag = prompt.Input("Enter tag > ", s.Completer.Complete)
		data.TechData.Type = datatypes.TextDataType
		data.Data = prompt.Input("Enter some text > ", s.Completer.Complete)
		if data.Data == "" {
			return customerror.ErrEmptyFields
		}
		resp, err := s.saveText(data, false)
		if err != nil {
			return err
		}
		fmt.Println()
		fmt.Printf("Record `%s` is saved\n", resp)
		fmt.Println("***")
		fmt.Println()
		return nil
	case "binary":
		data := models.BinaryDataModel{}
		fmt.Println("***")
		fmt.Println("Enter Title, Comment, Tag and binary file to save:")
		fmt.Println(hint)
		fmt.Println("***")
		data.TechData.Title = prompt.Input("Enter title > ", s.Completer.Complete)
		data.TechData.Comment = prompt.Input("Enter comment > ", s.Completer.Complete)
		data.TechData.Tag = prompt.Input("Enter tag > ", s.Completer.Complete)
		data.TechData.Type = datatypes.BinaryDataType
		path := prompt.Input("Enter path to binary file > ", s.Completer.Complete)
		resp, err := s.saveBinary(data, path)
		if err != nil {
			return err
		}
		fmt.Println()
		fmt.Printf("Record `%s` is saved\n", resp)
		fmt.Println("***")
		fmt.Println()
		return nil
	case "-f":
		data := models.TextDataModel{}
		fmt.Println("***")
		fmt.Println("Enter Title, Comment, Tag and text file to save:")
		fmt.Println(hint)
		fmt.Println("***")
		data.TechData.Title = prompt.Input("Enter title > ", s.Completer.Complete)
		data.TechData.Comment = prompt.Input("Enter comment > ", s.Completer.Complete)
		data.TechData.Tag = prompt.Input("Enter tag > ", s.Completer.Complete)
		data.TechData.Type = datatypes.TextDataType
		data.Data = prompt.Input("Enter path to text file > ", s.Completer.Complete)
		resp, err := s.saveText(data, true)
		if err != nil {
			return err
		}
		fmt.Println()
		fmt.Printf("Record `%s` is saved\n", resp)
		fmt.Println("***")
		fmt.Println()
		return nil

	default:
		return customerror.ErrInvalildCommand
	}
}

// editRecordDialogue - updates the record with new values.
func (s *ShellUI) editRecordDialogue(in string) (string, error) {
	in = strings.TrimSpace(in)
	commands := strings.Split(in, " ")
	if len(commands) < 3 {
		return "", customerror.ErrInvalildCommand
	}
	switch commands[1] {
	case "lp":
		key, err := strconv.Atoi(commands[2])
		if err != nil {
			return "", customerror.ErrInvalildCommand
		}
		data := models.LogPwdModel{}
		fmt.Println("***")
		fmt.Println("Enter a new Title, Comment, Tag, Login and Password to update:")
		fmt.Println(hint)
		fmt.Println("***")
		data.TechData.Title = prompt.Input("Enter a new title > ", s.Completer.Complete)
		data.TechData.Comment = prompt.Input("Enter a new comment > ", s.Completer.Complete)
		data.TechData.Tag = prompt.Input("Enter a new tag > ", s.Completer.Complete)
		data.TechData.Type = datatypes.LoginPasswordDataType
		data.Login = prompt.Input("Enter a new login > ", s.Completer.Complete)
		data.Password = prompt.Input("Enter a new password > ", s.Completer.Complete)
		if data.Login == "" || data.Password == "" {
			return "", customerror.ErrEmptyFields
		}
		resp, err := s.Client.UpdateRecord(key, datatypes.LoginPasswordDataType, data)
		if err != nil {
			return "", err
		}
		return resp, nil

	case "card":
		key, err := strconv.Atoi(commands[2])
		if err != nil {
			return "", customerror.ErrInvalildCommand
		}
		data := models.CardModel{}
		fmt.Println("***")
		fmt.Println("Enter a new Title, Comment, Tag and bank card data to update:")
		fmt.Println(hint)
		fmt.Println("***")
		data.TechData.Title = prompt.Input("Enter a new title > ", s.Completer.Complete)
		data.TechData.Comment = prompt.Input("Enter a new comment > ", s.Completer.Complete)
		data.TechData.Tag = prompt.Input("Enter a new tag > ", s.Completer.Complete)
		data.TechData.Type = datatypes.CardDataType
		data.Num = prompt.Input("Enter a new card number > ", s.Completer.Complete)
		data.CVC = prompt.Input("Enter a new CVC number > ", s.Completer.Complete)
		data.Date = prompt.Input("Enter a new expiration date > ", s.Completer.Complete)
		data.FirstName = prompt.Input("Enter a new First name > ", s.Completer.Complete)
		data.LastName = prompt.Input("Enter a new Last name > ", s.Completer.Complete)
		if data.Num == "" || data.CVC == "" || data.Date == "" || data.FirstName == "" || data.LastName == "" {
			return "", customerror.ErrEmptyFields
		}
		err = validcard.CnValidator(data.Num)
		if err != nil {
			return "", err
		}
		err = validcard.CVCValidator(data.CVC)
		if err != nil {
			return "", err
		}
		err = validcard.ExpValidator(data.Date)
		if err != nil {
			return "", err
		}
		resp, err := s.Client.UpdateRecord(key, datatypes.CardDataType, data)
		if err != nil {
			return "", err
		}
		return resp, nil

	case "text":
		if len(commands) > 3 {
			switch commands[3] {
			case "-f":
				key, err := strconv.Atoi(commands[2])
				if err != nil {
					return "", customerror.ErrInvalildCommand
				}
				data := models.TextDataModel{}
				fmt.Println("***")
				fmt.Println("Enter a new Title, Comment, Tag and text file to update:")
				fmt.Println(hint)
				fmt.Println("***")
				data.TechData.Title = prompt.Input("Enter a new title > ", s.Completer.Complete)
				data.TechData.Comment = prompt.Input("Enter a new comment > ", s.Completer.Complete)
				data.TechData.Tag = prompt.Input("Enter a new tag > ", s.Completer.Complete)
				data.TechData.Type = datatypes.TextDataType
				data.Data = prompt.Input("Enter path to text file > ", s.Completer.Complete)
				resp, err := s.updateText(key, data, true)
				if err != nil {
					return "", err
				}
				return resp, nil
			}
		}
		key, err := strconv.Atoi(commands[2])
		if err != nil {
			return "", customerror.ErrInvalildCommand
		}
		data := models.TextDataModel{}
		fmt.Println("***")
		fmt.Println("Enter a new Title, Comment, Tag and some text to update:")
		fmt.Println(hint)
		fmt.Println("***")
		data.TechData.Title = prompt.Input("Enter a new title > ", s.Completer.Complete)
		data.TechData.Comment = prompt.Input("Enter a new comment > ", s.Completer.Complete)
		data.TechData.Tag = prompt.Input("Enter a new tag > ", s.Completer.Complete)
		data.TechData.Type = datatypes.TextDataType
		data.Data = prompt.Input("Enter some text > ", s.Completer.Complete)
		if data.Data == "" {
			return "", customerror.ErrEmptyFields
		}
		resp, err := s.updateText(key, data, false)
		if err != nil {
			return "", err
		}
		return resp, nil

	case "binary":
		key, err := strconv.Atoi(commands[2])
		if err != nil {
			return "", customerror.ErrInvalildCommand
		}
		data := models.BinaryDataModel{}
		fmt.Println("***")
		fmt.Println("Enter a new Title, Comment, Tag and binary file to save:")
		fmt.Println(hint)
		fmt.Println("***")
		data.TechData.Title = prompt.Input("Enter a new title > ", s.Completer.Complete)
		data.TechData.Comment = prompt.Input("Enter a new comment > ", s.Completer.Complete)
		data.TechData.Tag = prompt.Input("Enter a new tag > ", s.Completer.Complete)
		data.TechData.Type = datatypes.BinaryDataType
		path := prompt.Input("Enter path to binary file > ", s.Completer.Complete)
		resp, err := s.updateBinary(key, data, path)
		if err != nil {
			return "", err
		}
		return resp, nil

	default:
		return "", customerror.ErrInvalildCommand
	}
}

// saveLogPwd - save login/password data on the server.
func (s *ShellUI) saveLogPwd(data models.LogPwdModel) (string, error) {
	resp, err := s.Client.SendLogPwd(data)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// saveCard - save banks card data on the server.
func (s *ShellUI) saveCard(data models.CardModel) (string, error) {
	resp, err := s.Client.SendCard(data)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// saveText - save text data on the server.
// If file true - saved object file.
// If file false - saved object input text.
func (s *ShellUI) saveText(data models.TextDataModel, file bool) (string, error) {
	if !file {
		resp, err := s.Client.SendText(data)
		if err != nil {
			return "", err
		}
		return resp, nil
	}
	path := strings.TrimSpace(data.Data)
	fromFile, err := filetools.ReadTextFile(path)
	if err != nil {
		return "", err
	}
	data.Data = fromFile
	resp, err := s.Client.SendText(data)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// saveBinary - save binary data on the server.
func (s *ShellUI) saveBinary(data models.BinaryDataModel, path string) (string, error) {
	path = strings.TrimSpace(path)
	fromFile, err := filetools.ReadBinaryFile(path)
	if err != nil {
		return "", err
	}
	data.Data = fromFile
	resp, err := s.Client.SendBinary(data)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// updateText - update text data on the server.
// If file true - saved object file.
// If file false - saved object input text.
func (s *ShellUI) updateText(key int, data models.TextDataModel, file bool) (string, error) {
	if !file {
		resp, err := s.Client.UpdateRecord(key, datatypes.TextDataType, data)
		if err != nil {
			return "", err
		}
		return resp, nil
	}
	path := strings.TrimSpace(data.Data)
	fromFile, err := filetools.ReadTextFile(path)
	if err != nil {
		return "", err
	}
	data.Data = fromFile
	resp, err := s.Client.UpdateRecord(key, datatypes.TextDataType, data)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// updateBinary - update binary data on the server.
// If file true - saved object file.
// If file false - saved object input text.
func (s *ShellUI) updateBinary(key int, data models.BinaryDataModel, path string) (string, error) {
	path = strings.TrimSpace(path)
	fromFile, err := filetools.ReadBinaryFile(path)
	if err != nil {
		return "", err
	}
	data.Data = fromFile
	resp, err := s.Client.UpdateRecord(key, datatypes.BinaryDataType, data)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// getDataTable - returns table with information on request
func (s *ShellUI) getDataTable(command string) (string, error) {
	dataAll, err := s.Client.UpdateInfo()
	if err != nil {
		return "", err
	}
	switch command {
	case "lp":
		s.Client.Config.Storage.Clear(datatypes.LoginPasswordDataType)
		s.Client.Config.Storage.UpdateStorage(datatypes.LoginPasswordDataType, dataAll)
		data := s.Client.Config.Storage.GetListRecords(datatypes.LoginPasswordDataType)
		table := tablesgen.NewInfoTable(datatypes.LoginPasswordDataType, data)
		return table, nil
	case "card":
		s.Client.Config.Storage.Clear(datatypes.CardDataType)
		s.Client.Config.Storage.UpdateStorage(datatypes.CardDataType, dataAll)
		data := s.Client.Config.Storage.GetListRecords(datatypes.CardDataType)
		table := tablesgen.NewInfoTable(datatypes.CardDataType, data)
		return table, nil
	case "text":
		s.Client.Config.Storage.Clear(datatypes.TextDataType)
		s.Client.Config.Storage.UpdateStorage(datatypes.TextDataType, dataAll)
		data := s.Client.Config.Storage.GetListRecords(datatypes.TextDataType)
		table := tablesgen.NewInfoTable(datatypes.TextDataType, data)
		return table, nil
	case "binary":
		s.Client.Config.Storage.Clear(datatypes.BinaryDataType)
		s.Client.Config.Storage.UpdateStorage(datatypes.BinaryDataType, dataAll)
		data := s.Client.Config.Storage.GetListRecords(datatypes.BinaryDataType)
		table := tablesgen.NewInfoTable(datatypes.BinaryDataType, data)
		return table, nil

	default:
		return "", customerror.ErrInvalildCommand
	}
}

// getRecord - get current record from server.
// Returns some model current data type or nil.
func (s *ShellUI) getRecord(in string) (interface{}, error) {
	in = strings.TrimSpace(in)
	commands := strings.Split(in, " ")
	if len(commands) == 3 {
		switch commands[0] {
		case "get":
			switch commands[1] {
			case "lp":
				key, err := strconv.Atoi(commands[2])
				if err != nil {
					return models.LogPwdModel{}, err
				}
				data, err := s.Client.GetLogPwd(key)
				if err != nil {
					return models.LogPwdModel{}, err
				}
				return data, nil
			case "card":
				key, err := strconv.Atoi(commands[2])
				if err != nil {
					return models.CardModel{}, err
				}
				data, err := s.Client.GetCard(key)
				if err != nil {
					return models.CardModel{}, err
				}
				return data, nil
			case "text":
				key, err := strconv.Atoi(commands[2])
				if err != nil {
					return models.TextDataModel{}, err
				}
				data, err := s.Client.GetText(key)
				if err != nil {
					return models.TextDataModel{}, err
				}
				return data, nil
			case "binary":
				key, err := strconv.Atoi(commands[2])
				if err != nil {
					return models.BinaryDataModel{}, err
				}
				data, err := s.Client.GetBinary(key)
				if err != nil {
					return models.BinaryDataModel{}, err
				}
				return data, nil

			default:
				return nil, customerror.ErrInvalildCommand
			}

		default:
			return nil, customerror.ErrInvalildCommand
		}
	}
	return nil, customerror.ErrInvalildCommand
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
			return customerror.ErrInvalildCommand
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
