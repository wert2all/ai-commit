package ui

import "github.com/charmbracelet/lipgloss"

type Card struct {
	title   string
	content string
	width   int
}

func NewCard(title, content string, width int) Card {
	if width < 20 {
		width = 20 // minimum reasonable width
	}

	return Card{
		title:   title,
		content: content,
		width:   width,
	}
}

// String renders the card as a string
func (c Card) String() string {
	// Adjust styles based on card width
	cardWidth := c.width
	titleWidth := cardWidth - 2   // account for padding
	contentWidth := cardWidth - 4 // account for padding

	adjustedCardStyle := cardStyle.Width(cardWidth)
	adjustedTitleStyle := titleStyle.Width(titleWidth)
	adjustedContentStyle := contentStyle.Width(contentWidth)

	// Style the title and content
	title := adjustedTitleStyle.Render(c.title)
	content := adjustedContentStyle.Render(c.content)

	// Combine parts
	inner := lipgloss.JoinVertical(lipgloss.Left, title, content)
	return adjustedCardStyle.Render(inner)
}
