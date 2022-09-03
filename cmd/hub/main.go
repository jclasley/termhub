package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jclasley/termhub/internal"
	"github.com/jclasley/termhub/models/layout"
	"github.com/jclasley/termhub/models/panel"
	"github.com/jclasley/termhub/plugins/spotify"
	zone "github.com/lrstanley/bubblezone"
)

func main() {
	getFlags()
	cfg := GetConfig()
	internal.Setup()
	zone.NewGlobal()

	// DEBUG:
	var children []tea.Model
	for i := 0; i < panelCount; i++ {
		children = append(children, panel.Model{})
	}

	spotifyClient := spotify.ListenForCode(cfg.SpotifyClientID, cfg.SpotifySecret)
	sModel := spotify.New(spotifyClient)
	model := layout.New(sModel)

	// DEBUG:
	tea.LogToFile("log.txt", "DEBUG: ")

	if err := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion()).Start(); err != nil {
		panic(err)
	}

}
