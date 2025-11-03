package handlers

import (
	"dockerydo/internal/app/utils"
	"dockerydo/internal/docker"
	"dockerydo/internal/types"
	"dockerydo/internal/ui"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleContainersUpdate(msg types.ContainersMsg, m types.Model) (types.Model, tea.Cmd) {
	m.Containers = msg
	rows := make([]table.Row, 0, len(msg))

	for _, container := range msg {
		rows = append(rows, table.Row{
			container.Names,
			container.ID,
			container.State,
			container.RunningFor,
			ui.FormatPortsForTable(container),
		})
	}

	m.Table.SetRows(rows)
	return m, nil
}

func HandleTick(m types.Model) (types.Model, tea.Cmd) {
	if m.RefreshEnabled {
		return m, tea.Batch(docker.GetContainers, utils.TickCmd())
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

func HandleKeyPress(msg tea.KeyMsg, m types.Model) (types.Model, tea.Cmd) {
	// Error popup takes priority
	if m.ShowErrPopup {
		return HandleErrorPopupKeys(msg, m)
	}

	if m.ShowFailedOpPopup {
		return HandleFailedOpPopupKeys(msg, m)
	}

	if m.ShowConfirmPopup {
		return HandleConfimPopupKeys(msg, m)
	}

	// Detailed view mode
	if m.Mode == types.ViewModeDetailed {
		return HandleDetailedViewKeys(msg, m)
	}

	// Table view mode
	return HandleTableViewKeys(msg, m)
}
