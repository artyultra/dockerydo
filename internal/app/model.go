package app

import (
	"dockerydo/internal/docker"
	"dockerydo/internal/types"
	"dockerydo/internal/ui"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func tickCmd() tea.Cmd {
	return tea.Tick(10*time.Second, func(t time.Time) tea.Msg {
		return types.TickMsg(t)
	})
}

func NewModel() types.Model {
	return types.Model{
		Table: ui.NewTable(),
	}
}

func getSelectedContainer(m types.Model) *types.Container {
	selectedRow := m.Table.Cursor()
	if selectedRow < 0 || selectedRow >= len(m.Containers) {
		return nil
	}
	return &m.Containers[selectedRow]
}

func Init(m types.Model) tea.Cmd {
	return tea.Batch(docker.GetContainers, tickCmd())
}
