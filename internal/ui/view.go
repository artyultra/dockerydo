package ui

import (
	"dockerydo/internal/types"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func View(m types.Model) string {
	if m.Err != nil {
		return fmt.Sprintf("Error: %v\n\nPress q to quit.", m.Err)
	}
	if m.Mode == types.ViewModeDetailed {
		return RenderHeader(m) + m.ViewPort.View() + "\n r: Refresh • q: Quit\n"

	}
	style := lipgloss.NewStyle().
		Width(m.Width).
		Align(lipgloss.Center)

	return style.Render(m.Table.View() + "\n r: Refresh • q: Quit\n")
}
