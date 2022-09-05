package grid

import tea "github.com/charmbracelet/bubbletea"

type Component interface {
	tea.Model
	Position() (int, int) // Returns the (row, col) in which the component should be located
}
