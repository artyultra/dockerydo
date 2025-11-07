package handlers

import (
	"dockerydo/internal/app/utils"
	"dockerydo/internal/docker"
	"dockerydo/internal/types"

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

	rightWidth := int(float64(m.Width) * 0.6)
	contentHeight := m.Height - 7

	m.DetailsViewPort.Width = rightWidth - 8
	m.DetailsViewPort.Height = contentHeight
	m.LogsViewPort.Width = rightWidth - 8
	m.LogsViewPort.Height = contentHeight

	return m, nil
}
