package handlers

import (
	"dockerydo/internal/docker"
	"dockerydo/internal/theme"
	"dockerydo/internal/types"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func HandleKeyPress(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	// GLOBAL KEYS
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "ctrl+t":
		m.Theme = theme.ToggleTheme(m.Theme)
	}
	// Error popup takes priority
	if m.ShowErrPopup || m.ShowFailedOpPopup {
		return handleErrorPopupKeys(msg, m)
	}

	if m.ShowConfirmPopup {
		return handleConfirmPopupKeys(msg, m)
	}

	if m.RightPanel == types.PanelLogs {
		return handleLogsPanelKeys(msg, m)
	}

	// MAIN NAVIGATION

	return handleMainNavKeys(msg, m)
}

func handleMainNavKeys(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	switch msg.String() {
	// Tab switching
	case "left", "h":
		if m.ActiveTab == types.TabContainers {
			m.ActiveTab = types.TabNetworks
		} else {
			m.ActiveTab--
		}
	case "right", "l":
		if m.ActiveTab == types.TabNetworks {
			m.ActiveTab = types.TabContainers
		} else {
			m.ActiveTab++
		}
	case "1":
		m.ActiveTab = types.TabContainers
	case "2":
		m.ActiveTab = types.TabImages
	case "3":
		m.ActiveTab = types.TabVolumes
	case "4":
		m.ActiveTab = types.TabNetworks

	// list navigation
	case "up", "k":
		return handleCursorUp(m), nil
	case "down", "j":
		return handleCursorDown(m), nil

		// list selection
	case "enter":
		return handleEnterKey(m)
	case "s":
		return handleStartStop(m)
	case "p":
		return handlePauseUnpause(m)
	case "d":
		return handleRemove(m, false)
	case "D":
		return handleRemove(m, true)
	case "r":
		return handleRefresh(m)
	}
	return m, nil
}

func handleCursorUp(m types.Model) types.Model {
	switch m.ActiveTab {
	case types.TabContainers:
		if m.ContainerCursor > 0 {
			m.ContainerCursor--
		}
	case types.TabImages:
		if m.ImageCursor > 0 {
		}
	case types.TabVolumes:
		if m.VolumeCursor > 0 {
		}
	case types.TabNetworks:
		if m.NetworkCursor > 0 {
		}
	}
	return m
}

func handleCursorDown(m types.Model) types.Model {
	switch m.ActiveTab {
	case types.TabContainers:
		if m.ContainerCursor < len(m.Containers)-1 {
			m.ContainerCursor++
		}
	case types.TabImages:
		if m.ImageCursor < len(m.Images)-1 {
		}
	case types.TabVolumes:
		if m.VolumeCursor < len(m.Volumes)-1 {
		}
	case types.TabNetworks:
		if m.NetworkCursor < len(m.Networks)-1 {
		}
	}
	return m
}

func handleEnterKey(m types.Model) (types.Model, tea.Cmd) {
	if m.ActiveTab == types.TabContainers {
		if len(m.Containers) > 0 && m.ContainerCursor < len(m.Containers) {
			container := m.Containers[m.ContainerCursor]
			return m, docker.GetContainerLogs(container)
		}
	}
	return m, nil
}

func handleStartStop(m types.Model) (types.Model, tea.Cmd) {
	if m.ActiveTab == types.TabContainers {
		if len(m.Containers) > 0 && m.ContainerCursor < len(m.Containers) {
			return m, docker.StartStopContainer(m.Containers[m.ContainerCursor])
		}
	}
	return m, nil
}

func handlePauseUnpause(m types.Model) (types.Model, tea.Cmd) {
	if m.ActiveTab == types.TabContainers {
		if len(m.Containers) > 0 && m.ContainerCursor < len(m.Containers) {
			return m, docker.PauseUnpauseContainer(m.Containers[m.ContainerCursor])
		}
	}
	return m, nil
}

func handleRemove(m types.Model, force bool) (types.Model, tea.Cmd) {
	switch m.ActiveTab {

	case types.TabContainers:
		if len(m.Containers) > 0 && m.ContainerCursor < len(m.Containers) {
			container := m.Containers[m.ContainerCursor]
			if force {
				m.ShowConfirmPopup = true
				m.ConfirmPopUpMsg = fmt.Sprintf("Are you sure you want to delete %s ?", container.Names)
				return m, nil
			}
			return m, docker.RmContainer(container, false)
		}

	case types.TabImages:
		if len(m.Images) > 0 && m.ImageCursor < len(m.Images) {
			image := m.Images[m.ImageCursor]
			if force {
				m.ShowConfirmPopup = true
				m.ConfirmPopUpMsg = fmt.Sprintf("Are you sure you want to delete %s ?", image.ID)
				return m, nil
			}
			return m, docker.RmImage(image.ID, false)
		}

	case types.TabVolumes:
		if len(m.Volumes) > 0 && m.VolumeCursor < len(m.Volumes) {
			volume := m.Volumes[m.VolumeCursor]
			return m, docker.RmVolume(volume.Name, false)
		}

	case types.TabNetworks:
		if len(m.Networks) > 0 && m.NetworkCursor < len(m.Networks) {
			network := m.Networks[m.NetworkCursor]
			return m, docker.RmNetwork(network.ID, false)
		}
	}
	return m, nil
}

func handleRefresh(m types.Model) (types.Model, tea.Cmd) {
	return m, tea.Batch(
		docker.GetContainers,
		docker.GetImages,
		docker.GetVolumes,
		docker.GetNetworks,
	)
}

func handleConfirmPopupKeys(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	switch msg.String() {
	case "y", "enter":
		m.ShowConfirmPopup = false
		m.ConfirmPopUpMsg = ""
		return handleConfimAction(m)
	case "n", "esc":
		m.ShowConfirmPopup = false
		m.ConfirmPopUpMsg = ""
		return m, nil
	}
	return m, nil
}

func handleConfimAction(m types.Model) (types.Model, tea.Cmd) {
	switch m.ActiveTab {
	case types.TabContainers:
		if len(m.Containers) > 0 && m.ContainerCursor < len(m.Containers) {
			container := m.Containers[m.ContainerCursor]
			return m, docker.RmContainer(container, true)
		}
	case types.TabImages:
		if len(m.Images) > 0 && m.ImageCursor < len(m.Images) {
			image := m.Images[m.ImageCursor]
			return m, docker.RmImage(image.ID, true)
		}
	}

	return m, nil
}

func handleLogsPanelKeys(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.RightPanel = types.PanelDetails
		return m, nil
	case "up", "k":
		m.LogsViewPort.ScrollUp(1)
		return m, nil
	case "down", "j":
		m.LogsViewPort.ScrollDown(1)
		return m, nil
	}

	return m, nil
}

func handleErrorPopupKeys(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "esc":
		m.ShowErrPopup = false
		m.ShowFailedOpPopup = false
		return m, nil
	}
	return m, nil
}
