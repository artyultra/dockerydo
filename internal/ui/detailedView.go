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
	maxValueWidth := width - 25 // Reserve space for labels and margins

	sectionHeaderStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colors.Blue)).
		Align(lipgloss.Left).
		MarginTop(1).
		MarginBottom(1)

	labelStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colors.Yellow)).
		Width(18).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text)).
		MarginLeft(1)

	containerStyle := lipgloss.NewStyle().
		Margin(0, 2).
		Width(width).
		Height(m.Height - 4)

	var b strings.Builder

	// General Info Section
	b.WriteString(sectionHeaderStyle.Render("┌─ General Info") + "\n")
	b.WriteString(RenderField(labelStyle, valueStyle, "ID", smartTruncate(container.ID, 64)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Name", smartTruncate(container.Names, maxValueWidth)))
	b.WriteString(RenderField(labelStyle, valueStyle, "State", container.State))
	b.WriteString(RenderField(labelStyle, valueStyle, "Running For", smartTruncate(container.RunningFor, maxValueWidth)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Image", smartTruncate(container.Image, maxValueWidth)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Created At", smartTruncate(container.CreatedAt, maxValueWidth)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Command", smartTruncate(container.Command, maxValueWidth)))
	b.WriteString(RenderField(labelStyle, valueStyle, "Size", container.Size))

	// Ports Section
	if len(container.Ports) > 0 {
		b.WriteString("\n")
		b.WriteString(sectionHeaderStyle.Render("┌─ Network Ports") + "\n")
		for _, portMap := range container.Ports {
			portDisplay := formatPortMapping(portMap)
			if portDisplay != "" {
				b.WriteString(labelStyle.Render("→") + valueStyle.Render(portDisplay) + "\n")
			}
		}
	}

	// Compose Labels Section
	if container.RawLabels != "" {
		dl := container.Labels
		// Only show section if at least one label has content
		hasContent := dl.ComposeProject != "" || dl.ComposeService != "" || dl.ComposeImage != ""

		if hasContent {
			b.WriteString("\n")
			b.WriteString(sectionHeaderStyle.Render("┌─ Docker Compose") + "\n")
			if dl.ComposeProject != "" {
				b.WriteString(RenderField(labelStyle, valueStyle, "Project", smartTruncate(dl.ComposeProject, maxValueWidth)))
			}
			if dl.ComposeService != "" {
				b.WriteString(RenderField(labelStyle, valueStyle, "Service", smartTruncate(dl.ComposeService, maxValueWidth)))
			}
			if dl.ComposeImage != "" {
				b.WriteString(RenderField(labelStyle, valueStyle, "Image", smartTruncate(dl.ComposeImage, maxValueWidth)))
			}
			if dl.ComposeConfigFiles != "" {
				b.WriteString(RenderField(labelStyle, valueStyle, "Config Files", smartTruncate(dl.ComposeConfigFiles, maxValueWidth)))
			}
			if dl.ComposeConfigHash != "" {
				b.WriteString(RenderField(labelStyle, valueStyle, "Config Hash", smartTruncate(dl.ComposeConfigHash, 16)))
			}
			if dl.ComposeContainerNum != "" {
				b.WriteString(RenderField(labelStyle, valueStyle, "Container Number", dl.ComposeContainerNum))
			}
			if dl.ComposeDependsOn != "" {
				b.WriteString(RenderField(labelStyle, valueStyle, "Depends On", smartTruncate(dl.ComposeDependsOn, maxValueWidth)))
			}
			if dl.ComposeVersion != "" {
				b.WriteString(RenderField(labelStyle, valueStyle, "Version", dl.ComposeVersion))
			}
		}
	}

	return containerStyle.Render(b.String())

}

// formatPortMapping creates a Docker-like port display: "0.0.0.0:8080->80/tcp"
func formatPortMapping(pm types.PortMap) string {
	var parts []string

	// IPv4 mapping
	if pm.Ipv4 != "" && pm.ExternalRange != "" && pm.InternalRange != "" {
		proto := pm.Protocol
		if proto == "" {
			proto = "tcp"
		}
		parts = append(parts, pm.Ipv4+":"+pm.ExternalRange+"->"+pm.InternalRange+"/"+proto)
	}

	// IPv6 mapping
	if pm.Ipv6 != "" && pm.ExternalRange != "" && pm.InternalRange != "" {
		proto := pm.Protocol
		if proto == "" {
			proto = "tcp"
		}
		parts = append(parts, "["+pm.Ipv6+"]:"+pm.ExternalRange+"->"+pm.InternalRange+"/"+proto)
	}

	// Internal only (no external mapping)
	if pm.InternalRange != "" && pm.ExternalRange == "" {
		proto := pm.Protocol
		if proto == "" {
			proto = "tcp"
		}
		parts = append(parts, pm.InternalRange+"/"+proto)
	}

	return strings.Join(parts, ", ")
}

// smartTruncate only truncates if the string exceeds the max length
func smartTruncate(s string, maxLen int) string {
	if maxLen <= 0 || len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
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
	if strings.Contains(label, "-") {
		return labelStyle.Render(label) + valueStyle.Render(value) + "\n"
	}
	return labelStyle.Render(label+":") + valueStyle.Render(value) + "\n"
}

