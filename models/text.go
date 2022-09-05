package models

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Debug struct {
	text string
}

func (d Debug) Init() tea.Cmd {
	return nil
}

func (d Debug) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Printf("model w: %d\nmodel h: %d\n", msg.Width, msg.Height)
	}
	return d, nil
}

func (d Debug) View() string {
	return d.text
}

func New(s string) Debug {
	return Debug{text: s}
}
