package handlers

import (
	"dockerydo/internal/types"

	tea "github.com/charmbracelet/bubbletea"
)

func HandleError(msg types.ErrMsg, m types.Model) (types.Model, tea.Cmd) {
	m.Err = msg
	m.ShowErrPopup = true
	m.ErrPopUpMsg = msg.Error()
	return m, nil
}

func HandleConfirmation(msg types.ConfirmMsg, m types.Model) (types.Model, tea.Cmd) {
	return m, nil
}

func HandleFailedOp(msg types.OpFailedMsg, m types.Model) (types.Model, tea.Cmd) {
	m.ShowFailedOpPopup = true
	m.FailedOpPopUpMsg = msg.DaemonError
	return m, nil
}
