package genProject

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
type genProjFinishedMsg error
type viewState string
const (
	IDLE viewState = "IDLE"
	PICK_LOCATION viewState = "PICK_LOCATION"
	ERR  viewState = "ERR"
)

type Model struct {
	editor string
	input textinput.Model
	output, name, path  string
	state viewState
}

func New(output, editor string) Model {
	input := textinput.New()
	input.Width = 50
	input.Placeholder = "Enter Project Name"
	input.Focus()

	return Model{
		editor: editor,
		input: input,
		output: output,
		state: IDLE,
	}
}

func runGenProjCommand(args []string) tea.Cmd {
	return func() tea.Msg {
		done := make(chan error, 1)
		go func() {
			err := exec.Command("sf", args...).Run()
			done <- err
		}()
		err := <-done
		return genProjFinishedMsg(err)
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case genProjFinishedMsg:
			return m, tea.Quit
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyEnter:
					if m.state == IDLE {
						value := strings.TrimSpace(m.input.Value())
						if value == "" {
							m.state = ERR
							return m, nil
						}
						m.name = value
						m.state = PICK_LOCATION
						homeDir, _ := os.UserHomeDir()
						m.input.SetValue(homeDir)
						m.input.SetCursor(len(homeDir))
						return m, nil
					} else {
						m.path = strings.TrimSpace(m.input.Value())
						cmd := strings.Split(fmt.Sprintf("project generate --name %s -d %s --manifest", m.name, m.path), " ")
						return m, tea.Sequence(
							func() tea.Cmd {
								exec.Command("sf", cmd...).Run()
								return nil
							}(),
							tea.Quit,
						)
					}
			}
	}
	var inputCmd tea.Cmd
	m.input, inputCmd = m.input.Update(msg)
	return m, tea.Batch(inputCmd)
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
	case IDLE, PICK_LOCATION:
		return input
	case ERR:
		return lipgloss.JoinVertical(lipgloss.Left, input, errMsg)
	}
	return m.input.View()
}
