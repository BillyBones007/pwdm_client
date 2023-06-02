package uishell

import (
	"fmt"

	"github.com/BillyBones007/pwdm_client/internal/app/grpcclient"
	"github.com/c-bata/go-prompt"
)

// AppData - data of build version and date.
type AppData struct {
	BuildVersion string
	BuildDate    string
}

// LivePrefixState - dinamic prefix.
type LivePrefixState struct {
	User     string
	IsEnable bool
}

// ShellUI - main user interface structure.
type ShellUI struct {
	AppData   AppData
	Client    *grpcclient.ClientGRPC
	Completer Completer
	Prefix    LivePrefixState
}

// NewShellUI - constructor ShellUI.
func NewShellUI(appData AppData) *ShellUI {
	s := &ShellUI{AppData: appData, Client: grpcclient.InitClient()}
	s.Completer = *NewCompleter(s.Client)
	return s
}

// RunShell - runs the interactive user interface.
func (s *ShellUI) RunShell() {
	fmt.Println("\n*** Welcom to Data Manager CLI ***")
	fmt.Printf("version %s, build date %s\n", s.AppData.BuildVersion, s.AppData.BuildDate)
	fmt.Println(mainPage)
	defer fmt.Println("Bye!")
	p := prompt.New(
		s.executor,
		s.Completer.Complete,
		prompt.OptionTitle("interactive shell client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(s.changePrefix),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}

// changePrefix - if the user is logged in, the prefix is changed to the user name.
func (s *ShellUI) changePrefix() (string, bool) {
	return s.Prefix.User, s.Prefix.IsEnable
}
