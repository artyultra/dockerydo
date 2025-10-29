package app

import (
	"dockerydo/internal/docker"
	"dockerydo/internal/types"
	"dockerydo/internal/ui"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func Update(msg tea.Msg, m types.Model) (types.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		ui.UpdateTable(&m)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "r":
			return m, docker.GetContainers
		}
	case types.ContainersMsg:
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
	case types.ErrMsg:
		m.Err = msg
		return m, nil
	case types.TickMsg:
		return m, docker.GetContainers
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}
