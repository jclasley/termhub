package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jclasley/termhub/models/panel"
	"github.com/jclasley/termhub/plugins/spotify"
)

func main() {
	getFlags()

	// DEBUG:
	var children []tea.Model
	for i := 0; i < panelCount; i++ {
		children = append(children, panel.Model{})
	}

	tea.LogToFile("./log.txt", "DEBUG: ")
	model := spotify.New()

	if err := tea.NewProgram(spotify.New()).Start(); err != nil {
		panic(err)
	}
}
