package app

import (
	"dockerydo/internal/app/utils"
	"dockerydo/internal/docker"
	"dockerydo/internal/types"
	"dockerydo/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func NewModel() types.Model {
	return types.Model{
		Table:        ui.NewTable(),
		ShowErrPopup: false,
	}
}

func Init(m types.Model) tea.Cmd {
	return tea.Batch(docker.GetContainers, utils.TickCmd())
}
