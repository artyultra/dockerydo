package handlers

import (
	"dockerydo/internal/types"

	tea "github.com/charmbracelet/bubbletea"
)

func HandleErrorPopupKeys(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "esc":
		m.ShowErrPopup = false
		m.ErrPopUpMsg = ""
		m.Err = nil
	}
	return m, nil
}

func HandleError(msg types.ErrMsg, m types.Model) (types.Model, tea.Cmd) {
	m.Err = msg
	m.ShowErrPopup = true
	m.ErrPopUpMsg = msg.Error()
	return m, nil
}
