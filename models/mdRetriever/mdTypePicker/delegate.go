package mdTypePicker

import (
	"io"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type mdItemDelegate struct {}

var (
    checkboxStyle = lipgloss.
    	NewStyle().
      Bold(true).
      MarginRight(1).
      Foreground(lipgloss.Color("#0EEE91"))

    activeStyle = lipgloss.
    	NewStyle().
    	Border(lipgloss.ThickBorder(), false, false, false, true).
    	BorderForeground(lipgloss.Color("#1AB9FF"))
)

func checkbox(checked bool) string {
	var checkbox string
	if checkbox = "🞅"; checked {
		checkbox = "🞊"
	}
	return checkboxStyle.Render(checkbox)
}

func newMdItemDelegate() *mdItemDelegate {
	return &mdItemDelegate{}
}

func(d mdItemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	listItem, _ := item.(MdItem)

	listItemContent := fmt.Sprintf(
		"%s %s",
		checkbox(listItem.checked),
		listItem.title,
	)

	if index == m.Index() {
		listItemContent = activeStyle.Render(listItemContent)
	} else {
		listItemContent = lipgloss.NewStyle().PaddingLeft(1).Render(listItemContent)
	}
  fmt.Fprint(w, listItemContent)
}

func(d mdItemDelegate) Height() int {
	return 1
}

func(d mdItemDelegate) Spacing() int {
	return 0
}

func(d mdItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}
