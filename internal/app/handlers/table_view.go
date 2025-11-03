package handlers

import (
	"dockerydo/internal/app/utils"
	"dockerydo/internal/docker"
	"dockerydo/internal/types"
	"dockerydo/internal/ui"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleTableViewKeys(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "r":
		return m, docker.GetContainers
	case "enter":
		if container := utils.GetSelectedContainer(m); container != nil {
			return m, docker.InspectContainer(*container)
		}
	case "s":
		if container := utils.GetSelectedContainer(m); container != nil {
			return m, docker.StartStopContainer(*container)
		}
	case "p":
		if container := utils.GetSelectedContainer(m); container != nil {
			return m, docker.PauseUnpauseContainer(*container)
		}
	case "D":
		m.ShowConfirmPopup = true
		m.ConfirmPopUpMsg = "Are you sure you want to delete this container?"
	case "d":
		if container := utils.GetSelectedContainer(m); container != nil {
			return m, docker.RmContainer(*container, false)
		}
	}
	var cmd tea.Cmd
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func HandleInspect(msg types.InspectMsg, m types.Model) (types.Model, tea.Cmd) {
	selectedContainer := types.Container(msg)
	m.SelectedContainer = &selectedContainer
	m.ViewPort = viewport.New(m.Width, m.Height-8)
	m.RefreshEnabled = false
	m.Mode = types.ViewModeDetailed

	content := ui.RenderDetailedView(m)
	m.ViewPort.SetContent(content)

	return m, nil
}

func HandleFailedOpPopupKeys(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "esc":
		m.ShowFailedOpPopup = false
		m.FailedOpPopUpMsg = ""
	}
	return m, nil
}

func HandleFailedOp(msg types.OpFailedMsg, m types.Model) (types.Model, tea.Cmd) {
	m.ShowErrPopup = true
	m.ErrPopUpMsg = msg.DaemonError
	return m, nil
}

func HandleConfimPopupKeys(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	switch msg.String() {
	case "y", "enter":
		m.ShowConfirmPopup = false
		m.ConfirmPopUpMsg = ""
		if container := utils.GetSelectedContainer(m); container != nil {
			return m, docker.RmContainer(*container, true)
		}
	case "n", "esc":
		m.ShowConfirmPopup = false
		m.ConfirmPopUpMsg = ""
	}
	return m, nil
}

func HandleConfirmation(msg types.ConfirmMsg, m types.Model) (types.Model, tea.Cmd) {
	m.ShowConfirmPopup = true
	m.ConfirmPopUpMsg = string(msg)
	return m, nil
}
