package game

import "github.com/charmbracelet/bubbles/key"

// keyMap defines a set of keybindings. To work for help it must satisfy key.Map.
type keyMap struct {
	ToggleStartPause key.Binding
	Reset            key.Binding
	Quit             key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.ToggleStartPause, k.Reset, k.Quit}
}

// FullHelp returns keybindings for the expanded help view.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ToggleStartPause, k.Reset, k.Quit},
	}
}

var gameKeys = keyMap{
	ToggleStartPause: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("␣", "Start/Pause"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "Reset"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "Quit"),
	),
}
