package ui

import (
	"dockerydo/internal/types"
	// "fmt"

	"github.com/charmbracelet/lipgloss"
)

func View(m types.Model) string {
	// render base layout
	baseView := RenderBaseView(m)

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
