package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jclasley/termhub/internal"
	"github.com/jclasley/termhub/models"
	"github.com/jclasley/termhub/models/grid"
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

	bottomRight := models.New("bottom right")

	gridModel := grid.New(2, 2, map[grid.Position]tea.Model{
		grid.Position{Row: 0, Col: 0}: sModel,
		grid.Position{Row: 1, Col: 1}: bottomRight,
	})

	// DEBUG:
	tea.LogToFile("log.txt", "DEBUG: ")

	if err := tea.NewProgram(gridModel, tea.WithAltScreen(), tea.WithMouseCellMotion()).Start(); err != nil {
		panic(err)
	}

}
