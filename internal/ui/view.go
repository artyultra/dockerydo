package ui

import (
	"dockerydo/internal/types"
	// "fmt"

	"github.com/charmbracelet/lipgloss"
)

func View(m types.Model) string {
	// if m.Err != nil {
	// 	return fmt.Sprintf("Error: %v\n\nPress q to quit.", m.Err)
	// }
	var baseView string
	if m.Mode == types.ViewModeDetailed {
		baseView = RenderHeader(m) + m.ViewPort.View() + RenderDetailViewFooter(m.Width)

	} else {
		style := lipgloss.NewStyle().
			Width(m.Width).
			Background(lipgloss.Color(colors.Base)).
			Align(lipgloss.Center)

		baseView = style.Render(m.Table.View() + RenderTableFooter(m.Width))
	}

	if m.ShowErrPopup {
		return RenderErrPopup(m.ErrPopUpMsg, m.Width, m.Height)
	}

	if m.ShowFailedOpPopup {
		return RenderErrPopup(m.FailedOpPopUpMsg, m.Width, m.Height)
	}

	if m.ShowConfirmPopup {
		return RenderConfirmationPopup(m.ConfirmPopUpMsg, m.Width, m.Height)
	}

	return baseView
}
