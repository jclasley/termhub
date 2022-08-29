package main

import (
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jclasley/termhub/models/panel"
	"github.com/jclasley/termhub/plugins/spotify"
)

func main() {
	getFlags()
	cfg := GetConfig()

	// DEBUG:
	var children []tea.Model
	for i := 0; i < panelCount; i++ {
		children = append(children, panel.Model{})
	}

	tea.LogToFile("./log.txt", "DEBUG: ")
	sModel := spotify.New(cfg.SpotifyClientID, cfg.SpotifySecret)

	http.HandleFunc("/spotify/redirect", sModel.OauthHandler)

	if err := tea.NewProgram(spotify.New()).Start(); err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupOauthRoutes() {

}
