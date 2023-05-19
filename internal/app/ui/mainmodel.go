package ui

import (
	"time"

	"github.com/BillyBones007/pwdm_client/internal/app/grpcclient"
	"github.com/BillyBones007/pwdm_client/internal/datatypes"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	ClientGRPC *grpcclient.ClientGRPC
	StartPage  StartPageModel
	RegPage    RegistrationPageModel
	AuthPage   AuthPageModel
	CardPage   CardPageModel
	TabsPage   TabsPageModel
	Page       string
	User       string
	Err        error
}

func IitialMainModel(buildInfo BuildInfo) MainModel {
	model := MainModel{
		ClientGRPC: grpcclient.InitClient(),
		Page:       startPage,
		StartPage:  InitialStartPageModel(buildInfo),
		RegPage:    InitialRegistrationPageModel(),
		AuthPage:   InitialAuthPageModel(),
		CardPage:   InitialCardPageModel(),
	}
	model.TabsPage = InitialTabsPageModel(model.ClientGRPC)
	return model
}

func (m MainModel) Init() tea.Cmd {
	return m.AuthPage.Init()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		if m.ClientGRPC.AuthFlag {
			m.UpdateStorageData()
		}
	}
	switch m.Page {
	case startPage:
		startPage, cmd := m.StartPage.Update(msg)
		m.StartPage = startPage
		if cmd != nil {
			c := cmd()
			switch c := c.(type) {
			case tea.QuitMsg:
				return m, tea.Quit
			case startMsg:
				if c.signIn {
					m.Page = authPage
				} else if c.reg {
					m.Page = regPage
				}
				return m, cmd
			default:
				return m, cmd
			}
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
						m.AuthPage.err = err
						return m, nil
					}
					m.User = user
					if err = m.UpdateStorageData(); err != nil {
						m.Page = tabsPage
						m.Err = err
						return m, nil
					}
					m.Page = tabsPage
					return m, nil
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
						m.RegPage.err = err
						return m, nil
					}
					m.User = user
					if err = m.UpdateStorageData(); err != nil {
						m.Page = tabsPage
						m.Err = err
						return m, nil
					}
					m.Page = tabsPage
					return m, nil
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

	case tabsPage:
		tabsPageModel, cmd := m.TabsPage.Update(msg)
		m.TabsPage = tabsPageModel
		if cmd != nil {
			return m, cmd
		}
	}
	return m, tickEvery(time.Second * 3)
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
	case tabsPage:
		return m.TabsPage.View()
	default:
		return ""
	}
}

// tickEvery - returns tick every some time
func tickEvery(d time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(d)
		return tickMsg(time.Now())
	}
}

func (m *MainModel) UpdateStorageData() error {
	dataAll, err := m.ClientGRPC.UpdateInfo()
	if err != nil {
		return err
	}
	m.ClientGRPC.Storage.UpdateStorage(datatypes.LoginPasswordDataType, dataAll)
	m.ClientGRPC.Storage.UpdateStorage(datatypes.CardDataType, dataAll)
	m.ClientGRPC.Storage.UpdateStorage(datatypes.TextDataType, dataAll)
	m.ClientGRPC.Storage.UpdateStorage(datatypes.BinaryDataType, dataAll)
	return nil
}
