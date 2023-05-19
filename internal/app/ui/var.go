package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// Base variables
var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#4BFD87"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FD874B"))
	cursorStyle  = focusedStyle.Copy()
	noStyle      = lipgloss.NewStyle()
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	errStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#C1FD4B"))

	focusedButtonSubmit  = focusedStyle.Copy().Render("[ Submit ]")
	blurredButtonSubmit  = fmt.Sprintf(" %s ", blurredStyle.Render("[ Submit ]"))
	focusedButtonSignIn  = focusedStyle.Copy().Render("[ LogIn ]")
	blurredButtonSignIn  = fmt.Sprintf(" %s ", blurredStyle.Render("[ LogIn ]"))
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

	// Page names
	startPage string = "start"
	authPage  string = "auth"
	regPage   string = "registration"
	cardPage  string = "card"
	tabsPage  string = "tabs"
)

// ------------------------------------------------------------------

// Tabs variables
var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	// highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	highlightColor   = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#9864FD"}
	inactiveTabStyle = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle   = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle      = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()

	nameTabs = []string{"Login and Password", "Bank cards", "Text data", "Binary data"}
)

const (
	// Tabs id
	LogPwdTab = iota
	CardTab
	TextTab
	BinaryTab
)

// helper function
func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

// ------------------------------------------------------------------

// Table variables
var (
	baseTableStyle    = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#9864FD"))
	defaultTableStyle = defaultTableStyles()
)

const (
	// Column names
	numRec     string = "Record №"
	titleRec   string = "Title"
	tagRec     string = "Tag"
	commentRec string = "Comment"
)

// helper function
func defaultTableStyles() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#9864FD")).BorderBottom(true).Bold(false)
	s.Selected = s.Selected.Foreground(lipgloss.Color("#9864FD")).Background(lipgloss.Color("#9864FD")).Bold(false)
	return s
}

// ------------------------------------------------------------------
