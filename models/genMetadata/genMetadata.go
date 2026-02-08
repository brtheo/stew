package newMetadata

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MetadataType string
const (
	LWC MetadataType = "LWC"
	ApexClass MetadataType = "ApexClass"
	ApexTrigger MetadataType = "ApexTrigger"
)

type viewState string
const (
	IDLE viewState = "IDLE"
	ERR  viewState = "ERR"
	DONE viewState = "DONE"
)

type Model struct {
	input textinput.Model
	metadataType MetadataType
	output string
	name string
	state viewState
}

func generateLWC(name, path string) []string {
	raw := fmt.Sprintf("lightning generate component --name %s --type lwc --output-dir %s/lwc", name, path)
	return strings.Split(raw, " ")
}

func New(metadataType MetadataType, output string) Model {
	input := textinput.New()
	input.Width = 50
	input.Placeholder = fmt.Sprintf("Enter %s name", metadataType)
	input.Focus()

	return Model{
		input: input,
		metadataType: metadataType,
		output: output,
		// "/home/brtheo/Code/SF/DEVORG/force-app/main/default"
		state: IDLE,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyEnter:
				  value := strings.TrimSpace(m.input.Value())
					if value == "" {
						m.state = ERR
						return m, nil
					}
					m.name = m.input.Value()

					return m, tea.Sequence(
						func() tea.Cmd {
							return func() tea.Msg {
								return exec.Command("sf", generateLWC(value, m.output)...).Run()
							}
						}(), tea.Quit,
					)
			}
	}
	var inputCmd tea.Cmd
	m.input, inputCmd = m.input.Update(msg)
	return m, inputCmd
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
	case IDLE:
		return input
	case ERR:
		return lipgloss.JoinVertical(lipgloss.Left, input, errMsg)
	case DONE:
		return "Metadata created successfully"
	}
	return m.input.View()
}
