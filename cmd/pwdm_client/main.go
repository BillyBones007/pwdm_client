package main

import (
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
}
