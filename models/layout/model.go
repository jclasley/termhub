package layout

import (
	"context"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jclasley/termhub/plugins/spotify"
	zone "github.com/lrstanley/bubblezone"
)

// Model is responsible for controlling the layout of its children
type Model struct {
	curModel     tea.Model
	children     []tea.Model
	input        textinput.Model
	enteringText bool
	w            int
	h            int
	err          error
	viewShell    bool
}

func (m Model) Init() tea.Cmd {
	return m.curModel.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.w, m.h = msg.Width, msg.Height-1 // minus 1 for input height

	// layout model responsible for handling pkg-wide commands
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			if m.enteringText {
				m.input.Reset()
				m.enteringText = false
				return m, nil
			}
			return m, tea.Quit
		case ":":
			m.input = textinput.New()
			m.input.SetCursorMode(textinput.CursorStatic)
			m.input.Width = m.w
			m.enteringText = true
			m.input.Focus()
			return m, nil

		case "enter":
			if m.enteringText {
				m.enteringText = false
				m.viewShell = true
				return m, tea.ExitAltScreen
				return m, tea.Sequentially(tea.ExitAltScreen, m.runCommand(m.input.Value()))
			}
			//return m, tea.EnterAltScreen
		}
	case error:
		m.err = msg

	case spotify.TickPlaying:
		if m.viewShell {

		}
	}

	var cCmd tea.Cmd
	m.curModel, cCmd = m.curModel.Update(msg)

	var iCmd tea.Cmd
	m.input, iCmd = m.input.Update(msg)
	return m, tea.Batch(cCmd, iCmd)
}

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	if m.viewShell {
		return ""
	}

	if m.enteringText {
		return zone.Scan(m.curModel.View() + "\n" + m.input.View())
	}
	return zone.Scan(m.curModel.View())
}

func New(initModel tea.Model) Model {
	curModel := initModel
	return Model{
		curModel: curModel,
	}
}

type Size struct {
	w int
	h int
}

func (m Model) runCommand(input string) tea.Cmd {
	return func() tea.Msg {
		input = strings.TrimRight(input, ";")
		input += "; echo 'Press any key to continue'; read -n 1 -s"
		cmd := exec.CommandContext(context.Background(), "sh", "-c", input,
			"echo 'Press ENTER to continue'", "read -n 1 -s")

		return tea.ExecProcess(cmd, func(err error) tea.Msg {
			return err
		})
	}
}
