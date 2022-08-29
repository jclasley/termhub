package internal

import tea "github.com/charmbracelet/bubbletea"

type Resizable interface {
	tea.Model
	UpdateSize(w, h int) tea.Cmd
}
