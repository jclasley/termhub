package spotify

import (
	tea "github.com/charmbracelet/bubbletea"
	"net/http"
)

type Model struct {
	title  string
	artist string
	client *http.Client
}

func New() Model {
	oauthToken()
	return Model{}
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

func (m *Model) SetClient(c *http.Client) {
	m.client = c
}

func (m *Model) OauthHandler(req *http.Request, w http.ResponseWriter) error {
	code := req.URL.Query().Get("code")
	tok, err :=
}
