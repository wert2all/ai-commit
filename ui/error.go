package ui

import "github.com/charmbracelet/lipgloss"

type UIError struct {
	message string
	width   int
}

func NewError(message string, width int) UIError {
	if width < 20 {
		width = 20 // minimum reasonable width
	}
	return UIError{message: message, width: width}
}

// String renders the card as a string
func (e UIError) String() string {
	cardStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(errorColor).
		Width(e.width)

	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(errorColor).
		Width(e.width)

	inner := lipgloss.JoinVertical(lipgloss.Left, titleStyle.Render(e.message))
	return cardStyle.Render(inner)
}
