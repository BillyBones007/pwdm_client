package ui

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type BuildInfo struct {
	Version string
	Date    string
}

func InitbuidInfo(version string, date string) BuildInfo {
	v := fmt.Sprintf("Build version: %s", version)
	d := fmt.Sprintf("Build date: %s", date)
	return BuildInfo{Version: textStyle.Render(v), Date: textStyle.Render(d)}
}

func StartUI(buildInfo BuildInfo) {
	// tea.WithAltScreen()
	p := tea.NewProgram(IitialMainModel(buildInfo))

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
