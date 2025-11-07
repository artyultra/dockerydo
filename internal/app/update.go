package app

import (
	"dockerydo/internal/app/handlers"
	"dockerydo/internal/docker"
	"dockerydo/internal/types"

	tea "github.com/charmbracelet/bubbletea"
)

func Update(msg tea.Msg, m types.Model) (types.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// update widnow size
	case tea.WindowSizeMsg:
		return handlers.HandleWindowResize(msg, m)
	// update key press
	case tea.KeyMsg:
		return handlers.HandleKeyPress(msg, m)
	// data updates
	case types.ContainersMsg:
		m.Containers = []types.Container(msg)
		return m, nil
	case types.ImagesMsg:
		m.Images = []types.Image(msg)
		return m, nil
	case types.VolumesMsg:
		m.Volumes = []types.Volume(msg)
		return m, nil
	case types.NetworksMsg:
		m.Networks = []types.Network(msg)
		return m, nil
	case types.LogsMsg:
		m.LogsViewPort.SetContent(msg.Log)
		m.RightPanel = types.PanelLogs

	// operation results
	case types.DockerOpMsg:
		switch msg.ResourceType {
		case types.ContainerResource:
			return m, docker.GetContainers
		case types.ImageResource:
			return m, docker.GetImages
		case types.VolumeResource:
			return m, docker.GetVolumes
		case types.NetworkResource:
			return m, docker.GetNetworks
		}
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
	m.DetailsViewPort, cmd = m.DetailsViewPort.Update(msg)
	return m, cmd
}
