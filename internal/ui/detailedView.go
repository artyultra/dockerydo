package ui

import (
	"dockerydo/internal/theme"
	"dockerydo/internal/types"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var colors = theme.Mocha

func RenderHeader(m types.Model) string {

	contName := m.SelectedContainer.Names
	contID := m.SelectedContainer.ID

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colors.Mauve)).
		Align(lipgloss.Center).
		Width(m.Width).
		MarginTop(1)

	dividerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Surface1))

	var b strings.Builder

	b.WriteString(titleStyle.Render(contName+" ("+contID+")") + "\n")
	b.WriteString(dividerStyle.Render("  " + strings.Repeat("_", m.Width-4) + "  \n\n"))

	return b.String()
}

func RenderDetailedView(m types.Model) string {
	container := m.SelectedContainer
	width := m.Width

	sectionHeaderStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colors.Blue)).
		Align(lipgloss.Left).
		MarginTop(1).
		MarginBottom(1)

	labelStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colors.Yellow)).
		Width(15).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text)).
		MarginLeft(1)

	containerStyle := lipgloss.NewStyle().
		Margin(0, 2).
		Width(width).
		Height(m.Height - 4)

	var b strings.Builder

	b.WriteString(sectionHeaderStyle.Render("General Info") + "\n")
	b.WriteString(RenderField(labelStyle, valueStyle, "ID", truncate(container.ID)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Name", truncate(container.Names)))
	b.WriteString(RenderField(labelStyle, valueStyle, "State", truncate(container.State)))
	b.WriteString(RenderField(labelStyle, valueStyle, "External port", truncate(container.ExternalPort)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Internal port", truncate(container.InternalPort)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Time up", truncate(container.RunningFor)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Image", truncate(container.Image)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Creation date", truncate(container.CreatedAt)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Command", truncate(container.Command)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Size RW", truncate(container.Size)))

	if container.RawLabels != "" {
		dl := container.Labels
		b.WriteString(sectionHeaderStyle.Render(" Compose Labels") + "\n")
		b.WriteString(RenderField(labelStyle, valueStyle, "Project", truncate(dl.ComposeProject)))
		b.WriteString(RenderField(labelStyle, valueStyle, "Service", truncate(dl.ComposeService)))
		b.WriteString(RenderField(labelStyle, valueStyle, "Image", truncate(dl.ComposeImage)))
		b.WriteString(RenderField(labelStyle, valueStyle, "Working_dir", truncate(dl.ComposeOneoff)))
		b.WriteString(RenderField(labelStyle, valueStyle, "Config_files", truncate(dl.ComposeConfigFiles)))
		b.WriteString(RenderField(labelStyle, valueStyle, "Config_hash", truncate(dl.ComposeConfigHash)))
		b.WriteString(RenderField(labelStyle, valueStyle, "Container_num", truncate(dl.ComposeContainerNum)))
		b.WriteString(RenderField(labelStyle, valueStyle, "Depends_on", truncate(dl.ComposeDependsOn)))
		b.WriteString(RenderField(labelStyle, valueStyle, "Version", truncate(dl.ComposeVersion)))
	}

	return containerStyle.Render(b.String())

}

func RenderDetailViewFooter(width int) string {
	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Lavender)).
		Align(lipgloss.Center).
		MarginTop(2).
		Width(width)

	return footerStyle.Render("\n esc: Return • q: Quit\n")
}

func RenderField(labelStyle, valueStyle lipgloss.Style, label, value string) string {
	return labelStyle.Render(label+":") + valueStyle.Render(value) + "\n"
}

func truncate(s string, max ...int) string {
	maxLen := 20 // default
	if len(max) > 0 {
		maxLen = max[0]
	}

	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
