package orgPicker

import (
	"io"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type orgItemDelegate struct {
}

var (
    checkboxStyle = lipgloss.
    	NewStyle().
      Bold(true).
      MarginTop(1).
      Foreground(lipgloss.Color("#0EEE91"))

    activeStyle = lipgloss.
    	NewStyle().
    	Border(lipgloss.ThickBorder(), false, false, false, true).
    	BorderForeground(lipgloss.Color("#1AB9FF"))
)

func checkbox(defaultOrg bool) string {
	var checkbox string
	if checkbox = "[ ]"; defaultOrg {
		checkbox = "[]"
	}
	return checkboxStyle.Render(checkbox)
}

func location(orgItem orgItem) string {
	at := lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#1AB9FF")).
		Render("@")
	username := lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#03B4A7")).
		Render(orgItem.username)
	instance := lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#2E844A")).
		Render(orgItem.instanceUrl)
	return fmt.Sprintf("%s %s %s", username, at, instance)
}

func alias(orgItem orgItem) string {
	return lipgloss.
		NewStyle().
		Border(lipgloss.RoundedBorder(), true, true, true, true).
		BorderForeground(lipgloss.Color("#FF538A")).
		Foreground(lipgloss.Color("#FDB6C5")).
		Render(orgItem.alias)
}

func newOrgItemDelegate() *orgItemDelegate {
	return &orgItemDelegate{}
}

func(d orgItemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	orgItem, _ := item.(orgItem)

	verticalStack := lipgloss.JoinVertical(lipgloss.Left, alias(orgItem), location(orgItem))
	listItem := lipgloss.JoinHorizontal(lipgloss.Left, checkbox(orgItem.isDefaultOrg), verticalStack)

	if index == m.Cursor() {
			listItem = activeStyle.Render(listItem)
	} else {
			listItem = lipgloss.NewStyle().PaddingLeft(2).Render(listItem)
	}
  fmt.Fprint(w, listItem)
}

func(d orgItemDelegate) Height() int {
	return 3
}

func(d orgItemDelegate) Spacing() int {
	return 1
}

func(d orgItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}
