package app

import (
	"dockerydo/internal/docker"
	"dockerydo/internal/types"
	"dockerydo/internal/ui"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func Update(msg tea.Msg, m types.Model) (types.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		ui.UpdateTable(&m)

		if m.Mode == types.ViewModeDetailed && m.SelectedContainer != nil {
			content := ui.RenderDetailedView(m)
			m.ViewPort.Width = m.Width
			m.ViewPort.Height = m.Height
			m.ViewPort.SetContent(content)
		}

		return m, tea.Batch(docker.GetContainers, tickCmd())

	case tea.KeyMsg:
		// handle inspect mode scrolling
		if m.Mode == types.ViewModeDetailed {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "esc":
				// go back to table view
				m.Mode = types.ViewModeTable
				m.SelectedContainer = nil
				return m, nil
			case "j", "down":
				m.ViewPort.ScrollDown(1)
				return m, nil
			case "k", "up":
				m.ViewPort.ScrollUp(1)
				return m, nil
			}
		}
		// table commands
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "r":
			return m, docker.GetContainers
		case "enter":
			container := getSelectedContainer(m)
			if container != nil {
				return m, docker.InspectContainer(*container)
			}
		}
	case types.ContainersMsg:
		m.Containers = msg
		rows := []table.Row{}
		for _, container := range msg {
			rows = append(rows, table.Row{
				container.Names,
				container.ID,
				container.RunningFor,
				container.State,
				container.ExternalPort,
				container.InternalPort,
			})
		}
		m.Table.SetRows(rows)
		return m, nil
	case types.InspectMsg:
		selectedContainer := types.Container(msg)
		m.SelectedContainer = &selectedContainer
		m.ViewPort = viewport.New(m.Width, m.Height-8)

		content := ui.RenderDetailedView(m)

		if m.Mode == types.ViewModeDetailed {
			m.ViewPort.SetContent(content)
		} else {
			m.Mode = types.ViewModeDetailed

			m.ViewPort.SetContent(content)
		}
		return m, nil

	case types.ErrMsg:
		m.Err = msg
		return m, nil
	case types.TickMsg:
		return m, docker.GetContainers
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}
