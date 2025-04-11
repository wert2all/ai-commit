package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	highlight  = lipgloss.AdaptiveColor{Light: "#003843", Dark: "#003843"}
	errorColor = lipgloss.AdaptiveColor{Light: "#f04e4f", Dark: "#f04e4f"}

	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(highlight).
			Padding(0, 1)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(highlight).
			Padding(0, 1)

	contentStyle = lipgloss.NewStyle().Padding(1, 2)
)
