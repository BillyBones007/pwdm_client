package ui

import (
	"fmt"

	"github.com/BillyBones007/pwdm_client/internal/app/grpcclient"
	"github.com/BillyBones007/pwdm_client/internal/datatypes"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	ClientGRPC *grpcclient.ClientGRPC
	Page       string
	User       string
	Err        string
	StartPage  StartPageModel
	RegPage    RegistrationPageModel
	AuthPage   AuthPageModel
	CardPage   CardPageModel
}

func IitialMainModel(buildInfo BuildInfo) MainModel {
	return MainModel{
		ClientGRPC: grpcclient.InitClient(),
		Page:       startPage,
		StartPage:  InitialStartPageModel(buildInfo),
		RegPage:    InitialRegistrationPageModel(),
		AuthPage:   InitialAuthPageModel(),
		CardPage:   InitialCardPageModel(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return m.AuthPage.Init()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.Page {
	case startPage:
		startPage, cmd := m.StartPage.Update(msg)
		m.StartPage = startPage
		if cmd != nil {
			c := cmd()
			mes := c.(startMsg)
			if mes.signIn {
				m.Page = authPage
			} else if mes.reg {
				m.Page = regPage
			}
			return m, cmd
		}
	case authPage:
		authPageModel, cmd := m.AuthPage.Update(msg)
		m.AuthPage = authPageModel
		if cmd != nil {
			c := cmd()
			switch c := c.(type) {
			case userMsg:
				if c.login != "" && c.password != "" {
					um := models.UserModel{Login: c.login, Password: c.password}
					user, err := m.ClientGRPC.LogIn(um)
					if err != nil {
						m.Page = authPage
						m.AuthPage.err = err.Error()
						return m, nil
					}
					m.User = user
					dataAll, err := m.ClientGRPC.UpdateInfo()
					if err != nil {
						fmt.Printf("ERROR: %s\n", err)
					}
					m.ClientGRPC.Storage.UpdateStorage(datatypes.LoginPasswordDataType, dataAll)
					m.ClientGRPC.Storage.UpdateStorage(datatypes.CardDataType, dataAll)
					m.ClientGRPC.Storage.UpdateStorage(datatypes.TextDataType, dataAll)
					m.ClientGRPC.Storage.UpdateStorage(datatypes.BinaryDataType, dataAll)
				}
				return m, cmd
			default:
				return m, cmd
			}
		}
	case regPage:
		regPageModel, cmd := m.RegPage.Update(msg)
		m.RegPage = regPageModel
		if cmd != nil {
			c := cmd()
			switch c := c.(type) {
			case userMsg:
				if c.login != "" && c.password != "" {
					um := models.UserModel{Login: c.login, Password: c.password}
					user, err := m.ClientGRPC.Registration(um)
					if err != nil {
						m.Page = regPage
						m.RegPage.err = err.Error()
						return m, nil
					}
					m.User = user
					dataAll, err := m.ClientGRPC.UpdateInfo()
					if err != nil {
						fmt.Printf("ERROR: %s\n", err)
					}
					m.ClientGRPC.Storage.UpdateStorage(datatypes.LoginPasswordDataType, dataAll)
					m.ClientGRPC.Storage.UpdateStorage(datatypes.CardDataType, dataAll)
					m.ClientGRPC.Storage.UpdateStorage(datatypes.TextDataType, dataAll)
					m.ClientGRPC.Storage.UpdateStorage(datatypes.BinaryDataType, dataAll)
				}
				return m, cmd
			default:
				return m, cmd
			}
		}
	case cardPage:
		cardPageModel, cmd := m.CardPage.Update(msg)
		m.CardPage = cardPageModel
		if cmd != nil {
			return m, cmd
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	switch m.Page {
	case startPage:
		return m.StartPage.View()
	case authPage:
		return m.AuthPage.View()
	case regPage:
		return m.RegPage.View()
	case cardPage:
		return m.CardPage.View()
	default:
		return ""
	}
}
