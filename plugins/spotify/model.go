package spotify

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"

	"net/http"
)

type Model struct {
	title    string
	artist   string
	client   *http.Client
	oauthCfg *oauth2.Config
}

func New(clientID, clientSecret string) Model {
	cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://accounts.spotify.com/api/token",
			AuthURL:  "https://accounts.spotify.com/authorize",
		},
		RedirectURL: "http://localhost:8080/redirect",
		Scopes:      []string{"user-read-playback-state"},
	}

	url := cfg.AuthCodeURL("state", oauth2.AccessTypeOffline)

	// Open the authorization URL in the user's browser
	if err := browser.OpenURL(url); err != nil {
		panic(err)
	}

	return Model{oauthCfg: cfg}
}

type updateTokenMsg struct{}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case updateTokenMsg:
		return m, nil
	}
	return m, nil
}

func (m Model) View() string {
	return ""
}
