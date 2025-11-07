package ui

import (
	"dockerydo/internal/theme"
	"dockerydo/internal/types"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// renderContainerDetails renders detailed info for a container
func RenderContainerDetails(c types.Container, width int, colors theme.Colors) string {
	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Lavender)).
		Bold(true).
		Width(15).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text))

	var lines []string

	// General Info
	lines = append(lines, renderSection("General", colors))
	lines = append(lines, renderField(labelStyle, valueStyle, "Name", c.Names))
	lines = append(lines, renderField(labelStyle, valueStyle, "ID", c.ID[:12]))
	lines = append(lines, renderField(labelStyle, valueStyle, "Image", c.Image))
	lines = append(lines, renderField(labelStyle, valueStyle, "Status", formatStatus(c.State, c.Status, colors)))
	lines = append(lines, renderField(labelStyle, valueStyle, "Created", c.CreatedAt))
	lines = append(lines, renderField(labelStyle, valueStyle, "Running For", c.RunningFor))
	if c.Size != "" {
		lines = append(lines, renderField(labelStyle, valueStyle, "Size", c.Size))
	}
	lines = append(lines, "")

	// Ports
	if len(c.Ports) > 0 {
		lines = append(lines, renderSection("Ports", colors))
		for _, port := range c.Ports {
			portStr := fmt.Sprintf("%s:%s → %s/%s",
				getIP(port),
				port.InternalRange,
				port.ExternalRange,
				port.Protocol)
			lines = append(lines, "  "+valueStyle.Render(portStr))
		}
		lines = append(lines, "")
	}

	// Docker Compose Info
	if c.Labels != nil && c.Labels.ComposeProject != "" {
		lines = append(lines, renderSection("Docker Compose", colors))
		lines = append(lines, renderField(labelStyle, valueStyle, "Project", c.Labels.ComposeProject))
		if c.Labels.ComposeService != "" {
			lines = append(lines, renderField(labelStyle, valueStyle, "Service", c.Labels.ComposeService))
		}
		if c.Labels.ComposeConfigFiles != "" {
			lines = append(lines, renderField(labelStyle, valueStyle, "Config", c.Labels.ComposeConfigFiles))
		}
		lines = append(lines, "")
	}

	// Command
	if c.Command != "" {
		lines = append(lines, renderSection("Command", colors))
		lines = append(lines, "  "+valueStyle.Render(c.Command))
	}

	return strings.Join(lines, "\n")
}

// renderImageDetails renders detailed info for an image
func renderImageDetails(img types.Image, width int, colors theme.Colors) string {
	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Lavender)).
		Bold(true).
		Width(15).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text))

	var lines []string

	lines = append(lines, renderSection("Image Details", colors))
	lines = append(lines, renderField(labelStyle, valueStyle, "Repository", img.Repository))
	lines = append(lines, renderField(labelStyle, valueStyle, "Tag", img.Tag))
	lines = append(lines, renderField(labelStyle, valueStyle, "ID", img.ID))
	lines = append(lines, renderField(labelStyle, valueStyle, "Created", img.CreatedSince+" ago"))
	lines = append(lines, renderField(labelStyle, valueStyle, "Size", img.Size))
	if img.Digest != "" {
		lines = append(lines, renderField(labelStyle, valueStyle, "Digest", img.Digest))
	}

	return strings.Join(lines, "\n")
}

// renderVolumeDetails renders detailed info for a volume
func renderVolumeDetails(vol types.Volume, width int, colors theme.Colors) string {
	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Lavender)).
		Bold(true).
		Width(15).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text))

	var lines []string

	lines = append(lines, renderSection("Volume Details", colors))
	lines = append(lines, renderField(labelStyle, valueStyle, "Name", vol.Name))
	lines = append(lines, renderField(labelStyle, valueStyle, "Driver", vol.Driver))
	lines = append(lines, renderField(labelStyle, valueStyle, "Mountpoint", vol.Mountpoint))
	lines = append(lines, renderField(labelStyle, valueStyle, "Scope", vol.Scope))
	if vol.Size != "" {
		lines = append(lines, renderField(labelStyle, valueStyle, "Size", vol.Size))
	}

	return strings.Join(lines, "\n")
}

// renderNetworkDetails renders detailed info for a network
func renderNetworkDetails(net types.Network, width int, colors theme.Colors) string {
	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Lavender)).
		Bold(true).
		Width(15).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text))

	var lines []string

	lines = append(lines, renderSection("Network Details", colors))
	lines = append(lines, renderField(labelStyle, valueStyle, "Name", net.Name))
	lines = append(lines, renderField(labelStyle, valueStyle, "ID", net.ID[:12]))
	lines = append(lines, renderField(labelStyle, valueStyle, "Driver", net.Driver))
	lines = append(lines, renderField(labelStyle, valueStyle, "Scope", net.Scope))
	if net.IPv6 != "" {
		lines = append(lines, renderField(labelStyle, valueStyle, "IPv6", net.IPv6))
	}
	if net.Internal != "" {
		lines = append(lines, renderField(labelStyle, valueStyle, "Internal", net.Internal))
	}

	return strings.Join(lines, "\n")
}

// Helper functions
func renderSection(title string, colors theme.Colors) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Peach)).
		Bold(true).
		Underline(true).
		Render(title)
}

func renderField(labelStyle, valueStyle lipgloss.Style, label, value string) string {
	return labelStyle.Render(label+":") + "  " + valueStyle.Render(value)
}

func formatStatus(state, status string, colors theme.Colors) string {
	statusStyle := lipgloss.NewStyle()

	switch state {
	case "running":
		statusStyle = statusStyle.Foreground(lipgloss.Color(colors.Green)).Bold(true)
	case "exited":
		statusStyle = statusStyle.Foreground(lipgloss.Color(colors.Red))
	case "paused":
		statusStyle = statusStyle.Foreground(lipgloss.Color(colors.Yellow))
	default:
		statusStyle = statusStyle.Foreground(lipgloss.Color(colors.Overlay0))
	}

	return statusStyle.Render(strings.ToUpper(state)) + " " + status
}

func getIP(port types.PortMap) string {
	if port.Ipv4 != "" {
		return port.Ipv4
	}
	if port.Ipv6 != "" {
		return "[" + port.Ipv6 + "]"
	}
	return "0.0.0.0"
}
