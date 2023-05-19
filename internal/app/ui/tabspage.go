package ui

import (
	"strings"

	"github.com/BillyBones007/pwdm_client/internal/app/grpcclient"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TabsPageModel struct {
	client     *grpcclient.ClientGRPC
	tabs       []string
	tabContent string
	user       string
	activeTab  int
}

func InitialTabsPageModel(c *grpcclient.ClientGRPC) TabsPageModel {
	tabs := TabsPageModel{client: c, tabs: nameTabs}
	return tabs
}

func (m TabsPageModel) Update(msg tea.Msg) (TabsPageModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit
		case "right", "l", "tab":
			m.activeTab = min(m.activeTab+1, len(m.tabs)-1)
			return m, nil
		case "left", "h", "shift+tab":
			m.activeTab = max(m.activeTab-1, 0)
			return m, nil
		}
	}

	return m, nil
}

func (m TabsPageModel) View() string {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range m.tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == LogPwdTab, i == BinaryTab, i == m.activeTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}
	doc.WriteString(m.user)
	doc.WriteString("\n")
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")

	switch m.activeTab {
	case LogPwdTab:
		m.tabContent = InitialTableModel(LogPwdTab, m.client).View()

	case CardTab:
		m.tabContent = InitialTableModel(CardTab, m.client).View()

	case TextTab:
		m.tabContent = InitialTableModel(TextTab, m.client).View()

	case BinaryTab:
		m.tabContent = InitialTableModel(BinaryTab, m.client).View()

	}
	doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(m.tabContent))
	return docStyle.Render(doc.String())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
