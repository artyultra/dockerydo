package handlers

import (
	"dockerydo/internal/app/utils"
	"dockerydo/internal/docker"
	"dockerydo/internal/types"

	tea "github.com/charmbracelet/bubbletea"
)

func HandleDetailedViewKeys(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "esc":
		m.Mode = types.ViewModeTable
		m.SelectedContainer = nil
		m.RefreshEnabled = true
		return m, tea.Batch(docker.GetContainers, utils.TickCmd())
	case "j", "down":
		m.ViewPort.ScrollDown(1)
	case "k", "up":
		m.ViewPort.ScrollUp(1)
	}
	return m, nil
}
