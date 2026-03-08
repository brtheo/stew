package genProject

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	// "github.com/charmbracelet/bubbles/list"
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
	// filePicker list.Model
}

func New(output, editor string) Model {
	input := textinput.New()
	input.Width = 50
	input.Placeholder = "Enter Project Name"
	input.Focus()

	// filePicker := list.New(
	// 	fillDirItems(homeDir),
	// 	list.NewDefaultDelegate(),
	// 	0,
	// 	0,
	// )
	return Model{
		// filePicker: filePicker,
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
			// fmt.Print("hello ?")
			// cmd := strings.Split(fmt.Sprintf("%s/%s/", m.path, m.name), " ")
			// exec.Command(m.editor, cmd...).Run()
			return m, tea.Quit
		// case tea.WindowSizeMsg:
			// m.filePicker.SetHeight(msg.Height - 1)
		case tea.KeyMsg:
			switch msg.Type {
				// case tea.KeyLeft:
					// m.filePicker.SetItems(fillDirItems(m.filePicker.SelectedItem().(DirItem).title))
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
						 exec.Command("sf", cmd...).Run()
							// if err ==  nil {
							// 	exec.Command(m.editor, fmt.Sprintf("~/Code/SF/%s", m.name)).Run()
							// 	return m, tea.Quit
							// }
						// return m, runGenProjCommand(cmd)
					}
			}
	}
	var inputCmd tea.Cmd
	m.input, inputCmd = m.input.Update(msg)
	// m.filePicker, fpCmd  = m.filePicker.Update(msg)
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
	// case PICK_LOCATION:
	// 	return m.filePicker.View()
	case ERR:
		return lipgloss.JoinVertical(lipgloss.Left, input, errMsg)
	}
	return m.input.View()
}
