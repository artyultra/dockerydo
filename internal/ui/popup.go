package ui

import (
	"dockerydo/internal/types"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// RenderErrPopup renders a beautiful error popup overlay
func RenderErrPopup(m types.Model) string {
	// Error icon and title
	errorTitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Red)).
		Bold(true).
		Render("⚠  ERROR")

	// Error message with word wrapping
	msgStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text)).
		Width(56)

	errorMsg := msgStyle.Render(m.ErrPopUpMsg)

	// Separator line
	separator := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Surface1)).
		Render(strings.Repeat("─", 54))

	// Instruction text
	instruction := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Overlay0)).
		Italic(true).
		Render("Press Enter or Esc to dismiss")

	// Combine all elements with proper spacing
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		errorTitle,
		"",
		errorMsg,
		"",
		separator,
		"",
		instruction,
	)

	// Main popup container with enhanced styling
	popup := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(colors.Red)).
		Background(lipgloss.Color(colors.Mantle)).
		Foreground(lipgloss.Color(colors.Text)).
		Padding(2, 3).
		Width(60).
		Render(content)

	// Place popup centered over the base view with overlay effect
	return lipgloss.Place(
		m.Width, m.Height,
		lipgloss.Center, lipgloss.Center,
		popup,
		lipgloss.WithWhitespaceChars("░"),
		lipgloss.WithWhitespaceForeground(lipgloss.Color(colors.Surface0)),
	)
}
