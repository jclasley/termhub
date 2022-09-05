package grid

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Position struct {
	Row int
	Col int
}

type Model struct {
	rowCount    int
	colCount    int
	childWidth  int
	childHeight int
	components  map[Position]tea.Model
}

func New(rowCount, colCount int, components map[Position]tea.Model) Model {
	return Model{
		rowCount:   rowCount,
		colCount:   colCount,
		components: components,
	}
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	for _, v := range m.components {
		cmds = append(cmds, v.Init())
	}
	return tea.Batch(cmds...)
}

type ComponentSize struct {
	Width  int
	Height int
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.childWidth = msg.Width / m.colCount
		m.childHeight = msg.Height / m.rowCount

		var cmds []tea.Cmd
		for _, v := range m.components {
			var cmd tea.Cmd
			v, cmd = v.Update(ComponentSize{Width: m.childWidth, Height: m.childHeight})
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd
	for k, v := range m.components {
		var cmd tea.Cmd
		m.components[k], cmd = v.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	// create styles for positioning
	baseStyle := lipgloss.NewStyle()

	var s string
	for i := 0; i < m.rowCount; i++ {
		for j := 0; j < m.colCount; j++ {
			style := baseStyle.Copy().Align(lipgloss.Position(i)).Width(m.childWidth).Height(m.childHeight)

			pos := Position{Row: i, Col: j}
			if v, ok := m.components[pos]; ok {
				s += style.Render(v.View())
			}
		}
	}
	return zone.Scan(s)
}
