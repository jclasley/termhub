package spotify

import "github.com/charmbracelet/lipgloss"

var buttonStyle = lipgloss.NewStyle().
	Padding(0, 1).
	Margin(0, 1).
	Background(lipgloss.Color("#009105")).
	Foreground(lipgloss.Color("#ffffff"))

var leftButtonStyle = lipgloss.NewStyle().Align(lipgloss.Left)
var rightButtonStyle = lipgloss.NewStyle().Align(lipgloss.Right)
var centerButtonStyle = lipgloss.NewStyle().Align(lipgloss.Center)
