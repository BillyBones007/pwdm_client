package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type StartPageModel struct {
	focusIndex   int
	signinButton int
	regButton    int
	buildInfo    BuildInfo
	text         string
}

func InitialStartPageModel(buildInfo BuildInfo) StartPageModel {
	m := StartPageModel{
		buildInfo:    buildInfo,
		focusIndex:   1,
		signinButton: 1,
		regButton:    2,
		text:         textStyle.Render(helpText),
	}
	return m
}

func (m StartPageModel) Update(msg tea.Msg) (StartPageModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "up", "down", "enter":
			s := msg.String()
			if s == "enter" && m.focusIndex == m.signinButton {
				return m, submitSignInStartPage()

			} else if s == "enter" && m.focusIndex == m.regButton {
				return m, submitRegistrationStartPage()
			}
			if s == "tab" || s == "down" {
				m.focusIndex++
			} else {
				m.focusIndex--
			}
			if m.focusIndex > m.regButton {
				m.focusIndex = m.signinButton
			} else if m.focusIndex < m.signinButton {
				m.focusIndex = m.regButton
			}
			return m, nil
		}
	}

	return m, nil
}

func (m StartPageModel) View() string {
	var b strings.Builder

	sButton := &blurredButtonSignIn
	rButton := &blurredButtonRegistr
	if m.focusIndex == m.signinButton {
		sButton = &focusedButtonSignIn
	} else {
		rButton = &focusedButtonRegistr
	}

	b.WriteString(textStyle.Render(helpText))

	fmt.Fprintf(&b, "\n\n%s\t%s\n\n", *sButton, *rButton)

	return b.String()
}

func submitSignInStartPage() tea.Cmd {
	return func() tea.Msg {
		return startMsg{
			signIn: true,
			reg:    false,
		}
	}
}

func submitRegistrationStartPage() tea.Cmd {
	return func() tea.Msg {
		return startMsg{
			signIn: false,
			reg:    true,
		}
	}
}
