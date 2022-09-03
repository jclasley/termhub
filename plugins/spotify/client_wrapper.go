package spotify

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickPlaying time.Time

func (m Model) tickCurPlaying() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		return TickPlaying(t)
	})
}

func (m Model) curPlaying() tea.Msg {
	playing, err := m.client.PlayerCurrentlyPlaying()
	if err != nil {
		return err
	}
	return playing
}

func (m Model) nextSong() tea.Msg {
	err := m.client.Next()
	if err != nil {
		return err
	}
	return nil
}

func (m Model) prevSong() tea.Msg {
	err := m.client.Previous()
	if err != nil {
		return err
	}
	return nil
}

func (m Model) togglePlay() tea.Msg {
	if m.isPlaying {
		if err := m.client.Pause(); err != nil {
			return err
		}
		return false
	}
	if err := m.client.Play(); err != nil {
		return err
	}
	return true
}
