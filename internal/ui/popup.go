package ui

import (
	"dockerydo/internal/theme"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// RenderErrPopup renders a beautiful error popup overlay
func RenderErrPopup(errorStr string, width, height int, colors theme.Colors) string {
	// Error icon and title
	errorTitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Red)).
		Bold(true).
		Render("⚠  ERROR")

	// Error message with word wrapping
	msgStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text)).
		Width(56)

	errorMsg := msgStyle.Render(errorStr)

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
		width, height,
		lipgloss.Center, lipgloss.Center,
		popup,
		lipgloss.WithWhitespaceChars("░"),
		lipgloss.WithWhitespaceForeground(lipgloss.Color(colors.Surface0)),
	)
}

func RenderConfirmationPopup(message string, width, height int, colors theme.Colors) string {
	// Warning icon and title
	confirmTitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Peach)).
		Bold(true).
		Render("⚠  CONFIRMATION REQUIRED")

	// Message with word wrapping
	msgStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text)).
		Width(56)
	confirmMsg := msgStyle.Render(message)

	// Separator line
	separator := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Surface1)).
		Render(strings.Repeat("─", 58))

	// Button styles
	proceedBtnStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Crust)).
		Background(lipgloss.Color(colors.Red)).
		Bold(true).
		Padding(0, 3).
		MarginRight(2)

	cancelBtnStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text)).
		Background(lipgloss.Color(colors.Surface1)).
		Padding(0, 3)

	proceedBtn := proceedBtnStyle.Render("PROCEED")
	cancelBtn := cancelBtnStyle.Render("CANCEL")

	// Center buttons
	buttons := lipgloss.JoinHorizontal(lipgloss.Left, proceedBtn, cancelBtn)
	centeredButtons := lipgloss.NewStyle().
		Width(60).
		Align(lipgloss.Center).
		Render(buttons)

	// Instruction text
	instruction := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Subtext0)).
		Italic(true).
		Align(lipgloss.Center).
		Width(60).
		Render("Enter/Y to proceed  •  Esc/N to cancel")

	// Combine all elements with proper spacing
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		confirmTitle,
		"",
		confirmMsg,
		"",
		separator,
		"",
		centeredButtons,
		"",
		instruction,
	)

	// Main popup container with enhanced styling
	popup := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(colors.Peach)).
		Background(lipgloss.Color(colors.Mantle)).
		Foreground(lipgloss.Color(colors.Text)).
		Padding(2, 3).
		Width(64).
		Render(content)

	// Place popup centered over the base view with overlay effect
	return lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		popup,
		lipgloss.WithWhitespaceChars("░"),
		lipgloss.WithWhitespaceForeground(lipgloss.Color(colors.Surface0)),
	)
}
