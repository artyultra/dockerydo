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
		Padding(1, 2).
		Width(width)

	var b strings.Builder

	b.WriteString(sectionHeaderStyle.Render("General Info") + "\n")
	b.WriteString(RenderField(labelStyle, valueStyle, "ID", truncate(container.ID, 12)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Name", truncate(container.Names, 12)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Time up", truncate(container.RunningFor, 12)))

	if container.Labels != "" {
		dl := parseLabels(container.Labels)
		b.WriteString(sectionHeaderStyle.Render("Labels") + "\n")
		b.WriteString(RenderField(labelStyle, valueStyle, "Project", dl.ComposeProject))
	}

	return containerStyle.Render(b.String())

}

func RenderFooter(m types.Model) string {
	return ""
}

func RenderField(labelStyle, valueStyle lipgloss.Style, label, value string) string {
	return labelStyle.Render(label+":") + valueStyle.Render(value) + "\n"
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

func parseLabels(labelsStr string) *types.DockerLabels {
	if labelsStr == "" {
		return &types.DockerLabels{}
	}

	labels := &types.DockerLabels{
		RawLabels: make(map[string]string),
	}

	pairs := strings.Split(labelsStr, ",")

	for _, pair := range pairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		labels.RawLabels[key] = value

		switch key {
		case "com.docker.compose.depends_on":
			labels.ComposeDependsOn = value
		case "com.docker.compose.image":
			labels.ComposeImage = value
		case "com.docker.compose.oneoff":
			labels.ComposeOneoff = value
		case "come.docker.compose.service":
			labels.ComposeService = value
		case "com.docker.compose.config-hash":
			labels.ComposeConfigHash = value
		case "com.docker.compose.container-number":
			labels.ComposeContainerNum = value
		case "com.docker.compose.project.working_dir":
			labels.ComposeWorkingDir = value
		case "com.docker.compose.version":
			labels.ComposeVersion = value
		case "com.docker.compose.project":
			labels.ComposeProject = value
		case "com.docker.compose.config_files":
			labels.ComposeConfigFiles = value
		}

	}

	return labels

}
