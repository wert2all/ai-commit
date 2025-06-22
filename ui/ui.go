package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/wert2all/ai-commit/ai"
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

func NewProviderInfo(providerInfo ai.ProviderInfo) string {
	var fullResponse strings.Builder

	fullResponse.WriteString("\nUsing ")
	fullResponse.WriteString(providerInfo.Name)
	fullResponse.WriteString(" with ")
	fullResponse.WriteString(providerInfo.Model)
	return fullResponse.String()
}
