package panel

import "github.com/charmbracelet/lipgloss"

// ideally want to get parent style and then set self to a size based
// on parent size

var panelStyle = lipgloss.NewStyle().
	BorderForeground(lipgloss.Color("#ffffff")).
	BorderStyle(lipgloss.RoundedBorder()).
	MarginLeft(1).
	PaddingLeft(1)
