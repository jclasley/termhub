package panel

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	h     int
	w     int
	style lipgloss.Style
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// need to receive resize message from parent
	case tea.WindowSizeMsg:
		m.h = msg.Height
		m.w = msg.Width

	}
	return m, nil
}

func (m Model) View() string {
	info := fmt.Sprintf("height: %d, width: %d, actual: %d x %d",
		panelStyle.GetHeight(), panelStyle.GetWidth(),
		m.h, m.w)
	return m.style.Inherit(panelStyle).Render(info)
}

func (m Model) SetStyle(style lipgloss.Style) tea.Cmd {
	m.style.Inherit(style)
	return nil
}
