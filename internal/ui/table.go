package ui

import (
	"dockerydo/internal/types"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func NewTable() table.Model {
	columns := []table.Column{
		{Title: "CONTAINER NAME", Width: 20},
		{Title: "ID", Width: 12},
		{Title: "STATUS", Width: 12},
		{Title: "UPTIME", Width: 15},
		{Title: "PORTS", Width: 25},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{}),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()

	// Use theme colors for consistency
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(colors.Surface1)).
		BorderBackground(lipgloss.Color(colors.Base)).
		Foreground(lipgloss.Color(colors.Blue)).
		Background(lipgloss.Color(colors.Base)).
		BorderBottom(true).
		Bold(true).
		MarginTop(2).
		// PaddingTop(3).
		Padding(0, 1)

	s.Selected = lipgloss.NewStyle().
		UnsetBackground().
		Foreground(lipgloss.Color(colors.Base)).
		Background(lipgloss.Color(colors.Mauve)).
		Bold(true)

	s.Cell = s.Cell.
		Padding(0, 1)

	t.SetStyles(s)

	return t
}

func UpdateTable(m *types.Model) {
	if m.Width == 0 || m.Height == 0 {
		return
	}

	numCols := 5 // Name, ID, Status, Uptime, Ports
	margin := 2
	borderOverhead := numCols + 1
	cellPadding := numCols * 2 // 1 padding on each side per column

	availableWidth := m.Width - margin*2 - borderOverhead - cellPadding

	// Minimum widths for each column type
	minName := 15
	minID := 12
	minStatus := 10
	minUptime := 12
	minPorts := 20

	// Calculate proportional widths
	// Priority: Name (25%), Ports (30%), Uptime (20%), ID (15%), Status (10%)
	nameWidth := max(minName, int(float64(availableWidth)*0.25))
	portsWidth := max(minPorts, int(float64(availableWidth)*0.30))
	uptimeWidth := max(minUptime, int(float64(availableWidth)*0.20))
	idWidth := max(minID, int(float64(availableWidth)*0.15))
	statusWidth := max(minStatus, int(float64(availableWidth)*0.10))

	// Adjust if total exceeds available width
	totalWidth := nameWidth + idWidth + statusWidth + uptimeWidth + portsWidth
	if totalWidth > availableWidth {
		// Scale down proportionally
		scale := float64(availableWidth) / float64(totalWidth)
		nameWidth = max(minName, int(float64(nameWidth)*scale))
		portsWidth = max(minPorts, int(float64(portsWidth)*scale))
		uptimeWidth = max(minUptime, int(float64(uptimeWidth)*scale))
		idWidth = max(minID, int(float64(idWidth)*scale))
		statusWidth = max(minStatus, int(float64(statusWidth)*scale))
	}

	columns := []table.Column{
		{Title: "CONTAINER NAME", Width: nameWidth},
		{Title: "ID", Width: idWidth},
		{Title: "STATUS", Width: statusWidth},
		{Title: "UPTIME", Width: uptimeWidth},
		{Title: "PORTS", Width: portsWidth},
	}

	m.Table.SetColumns(columns)

	// Account for header (3 lines), footer (3 lines), and margins
	// The table height is the number of visible rows, not total screen height
	headerFooterSpace := 3
	tableHeight := m.Height - headerFooterSpace
	if tableHeight < 5 {
		tableHeight = 5 // minimum visible rows
	}

	m.Table.SetHeight(tableHeight)
}

// max returns the larger of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func RenderTableFooter(width int) string {
	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Lavender)).
		Align(lipgloss.Center).
		Width(width).
		MarginTop(1)

	return footerStyle.Render("\n↑/↓: Navigate • enter: Details • r: Refresh • q: Quit\n")
}
