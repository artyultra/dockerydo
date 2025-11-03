package utils

import (
	"dockerydo/internal/types"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TickCmd() tea.Cmd {
	return tea.Tick(10*time.Second, func(t time.Time) tea.Msg {
		return types.TickMsg(t)
	})
}

func GetSelectedContainer(m types.Model) *types.Container {
	selectedRow := m.Table.Cursor()
	if selectedRow < 0 || selectedRow >= len(m.Containers) {
		return nil
	}
	return &m.Containers[selectedRow]
}
