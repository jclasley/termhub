package layout

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jclasley/termhub/models/panel"
	"log"
)

// Model is responsible for controlling the layout of its children
type Model struct {
	curModel tea.Model
	children []tea.Model
	keys     keymap
	help     help.Model
}

func (m Model) Init() tea.Cmd {
	return m.curModel.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Printf("w: %d, h: %d\n", msg.Width, msg.Height)
		layoutStyle = layoutStyle.Width(msg.Width - 4)
		layoutStyle = layoutStyle.Height(msg.Height - 4)
		return m, nil

	// layout model responsible for handling pkg-wide commands
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd
	for k, child := range m.children {
		var cmd tea.Cmd
		m.children[k], cmd = child.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	//v := m.children[0].View()
	childStyle = layoutStyle.Copy()

	if len(m.children) > 1 {
		childStyle = childStyle.Width(layoutStyle.GetWidth() / 2)
	}

	if len(m.children) > 2 {
		childStyle = childStyle.Height(layoutStyle.GetHeight() / 2)
	}

	for _, child := range m.children {
		c := child.(panel.Model)
		c.SetStyle(childStyle)
	}

	v := lipgloss.JoinHorizontal(lipgloss.Right, m.children[0].View(), m.children[1].View())

	return layoutStyle.Render(v) + "\n" + fmt.Sprintf("childStyle %d x %d", childStyle.GetWidth(),
		childStyle.GetHeight())
}

func New(initModel tea.Model, children []tea.Model) Model {
	curModel := initModel
	return Model{
		curModel: curModel,
		children: children,
		keys:     defaultKeymap,
		help:     help.New(),
	}
}

type Size struct {
	w int
	h int
}

func (m Model) setChildSize(w, h int) tea.Cmd {
	for _, child := range m.children {
		c := child.(panel.Model)
	}
	return tea.Batch()
}
