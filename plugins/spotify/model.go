package spotify

import (
	"log"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/zmb3/spotify"
)

type Model struct {
	title          string
	artist         string
	client         spotify.Client
	debug          string
	currentPlaying *spotify.CurrentlyPlaying
	err            error
	isPlaying      bool
	w              int
	h              int
}

func New(c *http.Client) Model {
	client := spotify.NewClient(c)
	return Model{client: client}
}

func (m Model) Init() tea.Cmd {
	return m.tickCurPlaying()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.w, m.h = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case error:
		log.Println(msg)
		m.err = msg
		return m, nil

	case string:

		return m, nil

	case tea.MouseMsg:
		log.Println("clicked")
		if msg.Type != tea.MouseLeft {
			return m, nil
		}

		if zone.Get("prev").InBounds(msg) {
			log.Println("clicked prev")
			return m, tea.Batch(m.prevSong, m.curPlaying)
		}

		if zone.Get("next").InBounds(msg) {
			log.Println("clicked next")
			return m, tea.Batch(m.nextSong, m.curPlaying)
		}

		if zone.Get("playPause").InBounds(msg) {
			return m, m.togglePlay
		}
	case bool: // used to toggle play button state
		m.isPlaying = msg

	case *spotify.CurrentlyPlaying:
		m.currentPlaying = msg
		m.isPlaying = msg.Playing

	case TickPlaying:
		return m, tea.Batch(m.curPlaying, m.tickCurPlaying())

	}
	return m, nil
}

func (m Model) View() string {
	if m.currentPlaying == nil {
		return "Loading...."
	}

	var playButton string
	if m.isPlaying {
		playButton = "❚❚"
	} else {
		playButton = "▶"
	}

	if m.err != nil {
		return "An unexpected error occurred\n\n" + m.err.Error() + "\n\nPress 'q' or 'esc' to quit"
	}

	var song string
	if playing := m.currentPlaying; playing != nil {
		if item := playing.Item; item != nil {
			song = item.Name
			if len(item.Artists) > 0 {
				song += " by " + item.Artists[0].Name
			}
		}
	}

	prev := zone.Mark("prev", "←")
	next := zone.Mark("next", "→")
	playPause := zone.Mark("playPause", playButton)
	buttonZone := buttonStyle.Render(leftButtonStyle.Render(prev) + "    " + centerButtonStyle.Render(playPause) +
		"    " + rightButtonStyle.Render(next))

	// zone.Scan should be in global manager
	return zone.Scan(playerStyle.Width(m.w).Height(m.h).Render(
		strings.Repeat("\n", (m.h/2)-3) + songStyle.Render(song) + "\n" + buttonZone))
}

var playerStyle = lipgloss.NewStyle().
	Align(lipgloss.Center)

var songStyle = lipgloss.NewStyle().
	Margin(1)
