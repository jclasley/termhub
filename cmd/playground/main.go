package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jclasley/termhub/models/grid"
	"github.com/jclasley/termhub/plugins/spotify"
)

func main() {
	sModel := spotify.New()

	model := grid.New(2, 2, map[grid.Position]tea.Model{
		grid.Position{Row: 0, Col: 0}: topLeft,
		grid.Position{Row: 1, Col: 1}: bottomRight,
	})

	if err := tea.NewProgram(model, tea.WithAltScreen()).Start(); err != nil {
		panic(err)
	}
}
