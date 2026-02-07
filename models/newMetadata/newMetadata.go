package newMetadata

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type metadataType string
const (
	LWC metadataType = "LWC"
	ApexClass metadataType = "ApexClass"
	ApexTrigger metadataType = "ApexTrigger"
)

type viewState string
const (
	IDLE viewState = "IDLE"
	ERR  viewState = "ERR"
	DONE viewState = "DONE"
)

type model struct {
	input textinput.Model
	metadataType metadataType
	path string
	name string
	state viewState
}
var docStyle = lipgloss.NewStyle().Margin(1, 1)

func generateLWC(name, path string) []string {
	raw := fmt.Sprintf("lightning generate component --name %s --type lwc --output-dir %s/lwc", name, path)
	return strings.Split(raw, " ")
}

func New(metadataType metadataType) model {
	input := textinput.New()
	input.Width = 50
	input.Placeholder = fmt.Sprintf("Enter %s name", metadataType)
	input.Focus()

	return model{
		input: input,
		metadataType: metadataType,
		path: "/home/brtheo/Code/SF/DEVORG/force-app/main/default",
		// TODO : replace hard coded path with args from main + using https://github.com/alexflint/go-arg
		state: IDLE,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyCtrlC, tea.KeyEsc:
					return m, tea.Quit
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
								return exec.Command("sf", generateLWC(value, m.path)...).Run()
							}
						}(), tea.Quit,
					)
			}
	}
	var inputCmd tea.Cmd
	m.input, inputCmd = m.input.Update(msg)
	return m, inputCmd
}

func (m model) View() string {
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
		return docStyle.Render(input)
	case ERR:
		return docStyle.Render(lipgloss.JoinVertical(lipgloss.Left, input, errMsg))
	case DONE:
		return docStyle.Render("Metadata created successfully")
	}
	return docStyle.Render(m.input.View())
}
