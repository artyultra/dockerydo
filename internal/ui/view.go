package ui

import (
	"dockerydo/internal/types"
)

func View(m types.Model) string {
	// render base layout
	baseView := RenderBaseView(m)

	// overlay popups
	if m.ShowErrPopup {
		return RenderErrPopup(m.ErrPopUpMsg, m.Width, m.Height, m.Theme)
	}

	if m.ShowFailedOpPopup {
		return RenderErrPopup(m.FailedOpPopUpMsg, m.Width, m.Height, m.Theme)
	}

	if m.ShowConfirmPopup {
		return RenderConfirmationPopup(m.ConfirmPopUpMsg, m.Width, m.Height, m.Theme)
	}

	return baseView
}
