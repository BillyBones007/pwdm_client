package ui

import (
	"fmt"
	"strings"

	"github.com/BillyBones007/pwdm_client/internal/customerror"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type RegistrationPageModel struct {
	focusIndex int
	inputs     []textinput.Model
	err        error
}

func InitialRegistrationPageModel() RegistrationPageModel {
	m := RegistrationPageModel{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Login"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		m.inputs[i] = t
	}

	return m
}

func (m RegistrationPageModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m RegistrationPageModel) Update(msg tea.Msg) (RegistrationPageModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "esc":
			m.Reset()
			return m, nil

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, m.setLogPwd()
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *RegistrationPageModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m RegistrationPageModel) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButtonSubmit
	if m.focusIndex == len(m.inputs) {
		button = &focusedButtonSubmit
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	if m.err != nil {
		if m.err.Error() == customerror.ErrGRPCUserIsExists.Error() {
			b.WriteString(errStyle.Render(customerror.ErrUserIsExists.Error()))
			b.WriteRune('\n')

		} else {
			b.WriteString(errStyle.Render(customerror.ErrUnknown.Error()))
			b.WriteRune('\n')
		}
	}

	b.WriteString(textStyle.Render("Registration:\n "))
	b.WriteString(textStyle.Render("\rEnter your login/password and press [ Submit ]\n\n "))
	b.WriteString(textStyle.Render(" \rCtrl+c or Esc - Quit\n"))
	b.WriteString(textStyle.Render(" \rTAB or \u2193 or enter - down\n"))
	b.WriteString(textStyle.Render(" \rShift+TAB or \u2191 - up\n"))

	return b.String()
}

func (m RegistrationPageModel) setLogPwd() tea.Cmd {
	return func() tea.Msg {
		return userMsg{login: m.inputs[0].Value(), password: m.inputs[1].Value()}
	}
}

func (m *RegistrationPageModel) Reset() {
	m.err = nil
}
