package handlers

import (
	"dockerydo/internal/app/utils"
	"dockerydo/internal/docker"
	"dockerydo/internal/types"
	"dockerydo/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func HandleTick(m types.Model) (types.Model, tea.Cmd) {
	if m.RefreshEnabled {
		return m, tea.Batch(
			docker.GetContainers,
			docker.GetImages,
			docker.GetVolumes,
			docker.GetNetworks,
			utils.TickCmd(),
		)
	}
	return m, nil
}

func HandleWindowResize(msg tea.WindowSizeMsg, m types.Model) (types.Model, tea.Cmd) {
	m.Width = msg.Width
	m.Height = msg.Height
	ui.UpdateTable(&m)

	if m.Mode == types.ViewModeDetailed && m.SelectedContainer != nil {
		content := ui.RenderDetailedView(m)
		m.ViewPort.Width = m.Width
		m.ViewPort.Height = m.Height
		m.ViewPort.SetContent(content)
	}

	return m, tea.Batch(docker.GetContainers, utils.TickCmd())
}
