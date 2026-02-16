package shared

import "github.com/charmbracelet/bubbles/key"

var leftKeyBind = key.NewBinding(
	key.WithKeys("left"),
	key.WithHelp("←", "navigate steps"),
)

var rightKeyBind = key.NewBinding(
	key.WithKeys("right"),
	key.WithHelp("→", "navigate steps"),
)

type NavKeyMap struct {
	Left  key.Binding
	Right key.Binding
}

func (k NavKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Left, k.Right}
}

var NavKeys = NavKeyMap{
	Left: leftKeyBind,
	Right: rightKeyBind,
}

type MdTableKeyMap struct {
	CtrlG key.Binding
	Tab key.Binding
	Left key.Binding
	Right key.Binding
}

func (k MdTableKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.CtrlG, k.Tab, k.Left, k.Right}
}
func (k MdTableKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.CtrlG, k.Tab, k.Left, k.Right}, // first column
	}
}

var CtrlKeys = MdTableKeyMap{
	CtrlG: key.NewBinding(
		key.WithKeys("Ctrl+G"),
		key.WithHelp("Ctrl+G", "generate package.xml"),
	),
	Tab: key.NewBinding(
		key.WithKeys("Tab"),
		key.WithHelp("Tab", "switch between columns"),
	),
	Left: leftKeyBind,
	Right: rightKeyBind,
}
