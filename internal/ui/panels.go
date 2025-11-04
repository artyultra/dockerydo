package ui

import (
	"dockerydo/internal/types"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func renderListPanel(m types.Model, width, height int) string {
	title := ""
	items := []string{}
	cursor := 0

	switch m.ActiveTab {
	case types.TabContainers:
		title = "Containers"
		cursor = m.ContainerCursor
		for i, c := range m.Containers {
			status := getStatusIcon(c.State)
			name := truncateString(c.Names, 25)
			line := fmt.Sprintf("%s %s", status, name)
			if i == cursor {
				line = lipgloss.NewStyle().
					Foreground(lipgloss.Color(m.Theme.Crust)).
					Background(lipgloss.Color(m.Theme.Mauve)).
					Bold(true).
					Width(width - 4).
					Render("▸ " + line)
			} else {
				line = lipgloss.NewStyle().
					Foreground(lipgloss.Color(m.Theme.Text)).
					Width(width - 4).
					Render(" " + line)
			}
			items = append(items, line)
		}
	}

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(m.Theme.Lavender)).
		Bold(true).
		Padding(0, 1).
		Background(lipgloss.Color(m.Theme.Surface0)).
		Width(width - 2)

	header := titleStyle.Render(title)

	// handle empty list
	if len(items) == 0 {
		emptyMsg := lipgloss.NewStyle().
			Foreground(lipgloss.Color(m.Theme.Overlay0)).
			Italic(true).
			Padding(2, 2).
			Render("No items")
		items = append(items, emptyMsg)
	}

	maxVisible := height - 3 // account for title and padding
	if maxVisible < 1 {
		maxVisible = 1
	}

	startIdx := 0
	endIdx := len(items)

	if len(items) > maxVisible {
		startIdx = cursor - maxVisible/2
		if startIdx < 0 {
			startIdx = 0
		}
		endIdx = startIdx + maxVisible
		if endIdx > len(items) {
			endIdx = len(items)
			startIdx = endIdx - maxVisible
			if startIdx < 0 {
				startIdx = 0
			}
		}
	}

	var visibleItems []string
	if startIdx < len(items) && endIdx <= len(items) && startIdx <= endIdx {
		visibleItems = items[startIdx:endIdx]
	} else {
		visibleItems = items
	}

	listContent := strings.Join(visibleItems, "\n")

	renderedLines := len(visibleItems)
	if renderedLines < maxVisible && maxVisible > 0 {
		padding := strings.Repeat("\n", maxVisible-renderedLines)
		listContent += padding
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		listContent,
	)

	containerWidth := width - 2
	containerHeight := height - 2

	if containerWidth < 5 {
		containerWidth = 5
	}

	if containerHeight < 3 {
		containerHeight = 3
	}

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(m.Theme.Surface1)).
		Background(lipgloss.Color(m.Theme.Base)).
		Width(containerWidth).
		Height(containerHeight).
		Padding(1)

	return containerStyle.Render(content)
}

func renderMainPanel(m types.Model, width, height int) string {
	var content string
	title := "Details"

	// ensure minimum demensions
	detailWidth := width - 2
	if detailWidth < 10 {
		detailWidth = 10
	}
	detailHeight := height - 3
	if detailHeight < 5 {
		detailHeight = 5
	}

	if m.RightPanel == types.PanelLogs {
		title = "Logs"
		if len(m.Containers) > 0 && m.ContainerCursor < len(m.Containers) {
			container := m.Containers[m.ContainerCursor]
			title = fmt.Sprintf("Logs - %s", container.Names)
		}
		content = m.LogsViewPort.View()
	} else {
		content = renderDetails(m, detailWidth, detailHeight)
	}
	return ""
}

// Helper functions
func getStatusIcon(state string) string {
	switch state {
	case "running":
		return "●"
	case "exited":
		return "○"
	case "paused":
		return "⏸"
	default:
		return "?"
	}
}

// renderDetails renders entity details based on active tab
func renderDetails(m types.Model, width, height int) string {
	switch m.ActiveTab {
	case types.TabContainers:
		if len(m.Containers) == 0 || m.ContainerCursor >= len(m.Containers) {
			return renderEmpty("No container selected")
		}
		return renderContainerDetails(m.Containers[m.ContainerCursor], width)

	case types.TabImages:
		if len(m.Images) == 0 || m.ImageCursor >= len(m.Images) {
			return renderEmpty("No image selected")
		}
		return renderImageDetails(m.Images[m.ImageCursor], width)

	case types.TabVolumes:
		if len(m.Volumes) == 0 || m.VolumeCursor >= len(m.Volumes) {
			return renderEmpty("No volume selected")
		}
		return renderVolumeDetails(m.Volumes[m.VolumeCursor], width)

	case types.TabNetworks:
		if len(m.Networks) == 0 || m.NetworkCursor >= len(m.Networks) {
			return renderEmpty("No network selected")
		}
		return renderNetworkDetails(m.Networks[m.NetworkCursor], width)
	}

	return renderEmpty("Unknown tab")
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func renderEmpty(msg string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Overlay0)).
		Italic(true).
		Padding(2, 2).
		Render(msg)
}
