package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("120"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("210"))
	cursorStyle  = focusedStyle.Copy()
	noStyle      = lipgloss.NewStyle()
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	errStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	focusedButtonSubmit  = focusedStyle.Copy().Render("[ Submit ]")
	blurredButtonSubmit  = fmt.Sprintf(" %s ", blurredStyle.Render("[ Submit ]"))
	focusedButtonSignIn  = focusedStyle.Copy().Render("[ SignIn ]")
	blurredButtonSignIn  = fmt.Sprintf(" %s ", blurredStyle.Render("[ SignIn ]"))
	focusedButtonRegistr = focusedStyle.Copy().Render("[ Registration ]")
	blurredButtonRegistr = fmt.Sprintf(" %s ", blurredStyle.Render("[ Registration ]"))
)

const (
	helpText string = `	Клиент менеджера паролей.

Шифрует полученные данные и сохраняет их на сервере.

Позволяет работать со следующими данными:
	1. Пары логин/пароль
	2. Данные банковских карт
	3. Произвольные текстовые данные
	4. Произвольные бинарные данные


Навигация:
	ВВЕРХ/ВНИЗ клавиши TAB / Shift+TAB 
	Выход из приложения Ctrl+c
	
Для входа в приложение необходимо авторизоваться/зарегестрироваться:`

	errText   string = "\tОшибка:\n\n"
	startPage string = "start"
	authPage  string = "auth"
	regPage   string = "registration"
	cardPage  string = "card"
	errPage   string = "error"
)
