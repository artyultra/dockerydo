package ui

import (
	"dockerydo/internal/types"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func NewTable() table.Model {
	columns := []table.Column{
		{Title: "Name", Width: 12},
		{Title: "ID", Width: 12},
		{Title: "Time Up", Width: 12},
		{Title: "State", Width: 12},
		{Title: "External", Width: 12},
		{Title: "Internal", Width: 12},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{}),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.ASCIIBorder()).
		BorderForeground(lipgloss.Color("#cdd6f4")).
		BorderBottom(true).
		PaddingTop(1)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#1e1e2e")).
		Background(lipgloss.Color("#cba6f7"))
	t.SetStyles(s)

	return t
}

func UpdateTable(m *types.Model) {
	numCols := len(m.Table.Columns())
	if m.Width == 0 || m.Height == 0 {
		return
	}

	margin := 2

	borderOverhead := numCols + 1

	availableWidth := m.Width - margin*2 - borderOverhead

	colWidth := availableWidth / numCols

	if colWidth < 8 {
		colWidth = 8
	}

	columns := []table.Column{
		{Title: "Name", Width: colWidth},
		{Title: "ID", Width: colWidth},
		{Title: "Time Up", Width: colWidth},
		{Title: "State", Width: colWidth},
		{Title: "External", Width: colWidth},
		{Title: "Internal", Width: colWidth},
	}

	m.Table.SetColumns(columns)
	m.Table.SetHeight(m.Height - 2)
}
