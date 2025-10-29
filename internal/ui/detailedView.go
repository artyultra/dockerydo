package ui

import (
	"dockerydo/internal/theme"
	"dockerydo/internal/types"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var colors = theme.Mocha

func RenderHeader(m types.Model) string {

	contName := m.SelectedContainer.Names
	contID := m.SelectedContainer.ID

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colors.Mauve)).
		Align(lipgloss.Center).
		Width(m.Width).
		MarginTop(1)

	dividerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Surface1))

	var b strings.Builder

	b.WriteString(titleStyle.Render(contName+" ("+contID+")") + "\n")
	b.WriteString(dividerStyle.Render("  " + strings.Repeat("_", m.Width-4) + "  \n\n"))

	return b.String()
}

func RenderDetailedView(m types.Model) string {
	// container := m.SelectedContainer
	// width := m.Width
	// height := m.Height

	return ""
}

func RenderFooter(m types.Model) string {
	return ""
}
