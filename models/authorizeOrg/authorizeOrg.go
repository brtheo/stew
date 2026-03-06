package authorizeOrg

import (
	"fmt"
	"os/exec"
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
	SET_CUSTOM_URL viewState = "SET_CUSTOM_URL"
	ERR  viewState = "ERR"
)

type authCommandFinishedMsg  error

type Model struct {
	list list.Model
	input textinput.Model
	alias string
	customUrl string
	orgType OrgType
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
	list.Title = "Select Org Type"
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

func runAuthCommand(args []string) tea.Cmd {
	return func() tea.Msg {
		done := make(chan error, 1)
		go func() {
			err := exec.Command("sf", args...).Run()
			done <- err
		}()
		err := <-done
		return authCommandFinishedMsg(err)
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.list.SetSize(msg.Width - 1, msg.Height - 2)
		case authCommandFinishedMsg:
			return m, tea.Quit
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyEnter:
					switch m.state {
					case PICK_ORG_TYPE:
						m.orgType = OrgType(m.list.SelectedItem().(OrgTypeItem).title)
						if m.orgType == CustomURL {
							m.input.Placeholder = "Enter Custom URL"
							m.state = SET_CUSTOM_URL
						} else {
							m.input.Placeholder = "Enter Org Alias"
							m.state = SET_ALIAS
						}
						return m, nil
					case SET_CUSTOM_URL:
						value := strings.TrimSpace(m.input.Value())
				 		if value == "" {
							m.state = ERR
						} else {
							m.customUrl = value
							m.input.SetValue("")
							m.input.Placeholder = "Enter Org Alias"
							m.state = SET_ALIAS
						}
						return m, nil
					case SET_ALIAS:
					  value := strings.TrimSpace(m.input.Value())
				 		if value == "" {
							m.state = ERR
							return m, nil
						} else {
							m.alias = value
							instanceUrl := ""
							switch m.orgType {
								case Sandbox:
									instanceUrl = " --instance-url https://test.salesforce.com"
								case CustomURL:
									instanceUrl = fmt.Sprintf(" --instance-url %s", m.customUrl)
							}
							raw := fmt.Sprintf("org login web%s --alias %s", instanceUrl, m.alias)
							cmd := strings.Split(raw, " ")
							return m, runAuthCommand(cmd)
						}
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
		Render(" Name cannot be empty")

	switch m.state {
	case PICK_ORG_TYPE:
		return m.list.View()
	case SET_ALIAS, SET_CUSTOM_URL:
		return input
	case ERR:
		return lipgloss.JoinVertical(lipgloss.Left, input, errMsg)
	}
	return m.input.View()
}
