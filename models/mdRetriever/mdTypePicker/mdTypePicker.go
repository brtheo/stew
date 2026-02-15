package mdTypePicker

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type HasPickedTypeMsg string

type keyMap struct {
	Left  key.Binding
	Right key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Left, k.Right}
}

var NavKeys = keyMap{
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "navigate steps"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "navigate steps"),
	),
}

type Model struct {
	List list.Model
}

func New() Model {
	list := list.New(
		MetadataTypes,
		newMdItemDelegate(),
		0,
		0,
	)
	list.Title = "Select Metadata Type"
	list.SetShowStatusBar(false)
	list.AdditionalShortHelpKeys = NavKeys.ShortHelp
	list.AdditionalFullHelpKeys = NavKeys.ShortHelp
	return Model{List: list}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func toggleCheckboxes(list list.Model) {
	index := list.Index()
	items := list.Items()
  for k, item := range items {
    if o, ok := item.(MdItem); ok {
      o.checked = k == index
      items[k] = o
    }
  }
  list.SetItems(items)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var listCmd tea.Cmd
	m.List, listCmd = m.List.Update(msg)
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyEnter:
					metadataType := m.List.SelectedItem().(MdItem).title
					toggleCheckboxes(m.List)
					return m, func() tea.Cmd {
						return func() tea.Msg {
							return HasPickedTypeMsg(metadataType)
						}
					}()
			}
	}
	return m, listCmd
}

func (m Model) View() string {
	return m.List.View()
}
