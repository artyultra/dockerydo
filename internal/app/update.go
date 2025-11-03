package app

import (
	"dockerydo/internal/app/handlers"
	"dockerydo/internal/docker"
	"dockerydo/internal/types"

	tea "github.com/charmbracelet/bubbletea"
)

func Update(msg tea.Msg, m types.Model) (types.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return handlers.HandleWindowResize(msg, m)
	case tea.KeyMsg:
		return handlers.HandleKeyPress(msg, m)
	case types.ContainersMsg:
		return handlers.HandleContainersUpdate(msg, m)
	case types.InspectMsg:
		return handlers.HandleInspect(msg, m)
	case types.ContainerOpMsg:
		return m, docker.GetContainers
	case types.OpFailedMsg:
		return handlers.HandleFailedOp(msg, m)
	case types.ConfirmMsg:
		return handlers.HandleConfirmation(msg, m)
	case types.ErrMsg:
		return handlers.HandleError(msg, m)
	case types.TickMsg:
		return handlers.HandleTick(m)
	}

	var cmd tea.Cmd
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}
