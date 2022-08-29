package layout

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

type keymap struct {
	Quit key.Binding
}

var defaultKeymap = keymap{
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "Quit"),
	),
}

func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Quit,
	}
}

func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}

var helpStyle = lipgloss.NewStyle().Padding(1)
