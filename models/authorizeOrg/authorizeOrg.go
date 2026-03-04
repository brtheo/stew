package authorizeOrg

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type viewState string
const (
	PICK_ORG_TYPE viewState = "PICK_ORG_TYPE"
	SET_ALIAS viewState = "SET_ALIAS"
	ERR  viewState = "ERR"
)

type Model struct {
	list list.Model
	input textinput.Model
	output, name, sobject string
	state viewState
}

func New() Model {
	input := textinput.New()
	input.Width = 50
	input.Focus()

	list := list.New(
		OrgTypes,
		list.NewDefaultDelegate(),
		0,
		0,
	)
	list.Title = ""
	list.SetShowStatusBar(false)

	return Model{
		list: list,
		input: input,
		state: PICK_ORG_TYPE,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.list.SetSize(msg.Width - 1, msg.Height - 2)
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyEnter:
				  value := strings.TrimSpace(m.input.Value())
			 		if value == "" {
						m.state = ERR
						return m, nil
					}
			}
	}
	var inputCmd, listCmd tea.Cmd
	m.input, inputCmd = m.input.Update(msg)
	m.list, listCmd = m.list.Update(msg)
	return m, tea.Batch(inputCmd, listCmd)
}

func (m Model) View() string {
	input := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#1AB9FF")).
		Render(m.input.View())
	errMsg := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#D83A00")).
		PaddingLeft(1).
		Render(" Name cannot be empty")

	switch m.state {
	case PICK_ORG_TYPE:
		return m.list.View()
	case SET_ALIAS:
		return input
	case ERR:
		return lipgloss.JoinVertical(lipgloss.Left, input, errMsg)
	}
	return m.input.View()
}
