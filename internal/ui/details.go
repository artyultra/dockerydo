package ui

import (
	"dockerydo/internal/theme"
	"dockerydo/internal/types"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// renderContainerDetails renders detailed info for a container
func RenderContainerDetails(c types.Container, width, height int, colors theme.Colors) string {
	var b strings.Builder

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Maroon)).
		Background(lipgloss.Color(colors.Base)).
		Underline(true).
		MarginTop(1).
		MarginBottom(1).
		Bold(true)

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Lavender)).
		Bold(true).
		Width(15).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text)).
		Background(lipgloss.Color(colors.Base))

	// General Info
	b.WriteString(renderSection(titleStyle, "General"))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Name", c.Names))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "ID", c.ID[:12]))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Image", c.Image))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Status", formatStatus(c.State, c.Status, colors)))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Created", c.CreatedAt))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Running For", c.RunningFor))
	b.WriteString("\n")
	if c.Size != "" {
		b.WriteString(renderField(labelStyle, valueStyle, "Size", c.Size))
		b.WriteString("\n")
	}
	b.WriteString("\n")

	// Ports
	if len(c.Ports) > 0 {
		b.WriteString(renderSection(titleStyle, "Ports"))
		b.WriteString("\n")
		for _, port := range c.Ports {
			portStr := fmt.Sprintf("%s:%s → %s/%s",
				getIP(port),
				port.InternalRange,
				port.ExternalRange,
				port.Protocol)
			b.WriteString("  ")
			b.WriteString(valueStyle.Render(portStr))
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	// Docker Compose Info
	if c.Labels != nil && c.Labels.ComposeProject != "" {
		b.WriteString(renderSection(titleStyle, "Docker Compose"))
		b.WriteString("\n")
		b.WriteString(renderField(labelStyle, valueStyle, "Project", c.Labels.ComposeProject))
		b.WriteString("\n")
		if c.Labels.ComposeService != "" {
			b.WriteString(renderField(labelStyle, valueStyle, "Service", c.Labels.ComposeService))
			b.WriteString("\n")
		}
		if c.Labels.ComposeConfigFiles != "" {
			b.WriteString(renderField(labelStyle, valueStyle, "Config", c.Labels.ComposeConfigFiles))
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	// Command
	if c.Command != "" {
		b.WriteString(renderSection(titleStyle, "Command"))
		b.WriteString("\n")
		b.WriteString("  ")
		b.WriteString(valueStyle.Render(c.Command))
		b.WriteString("\n")
	}

	containerStyle := lipgloss.NewStyle().
		Width(width - 2).
		Height(height)

	return containerStyle.Render(b.String())
}

// renderImageDetails renders detailed info for an image
func renderImageDetails(img types.Image, width, height int, colors theme.Colors) string {
	var b strings.Builder

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Maroon)).
		Background(lipgloss.Color(colors.Base)).
		Underline(true).
		MarginTop(1).
		MarginBottom(1).
		Bold(true)

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Lavender)).
		Bold(true).
		Width(15).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text)).
		Background(lipgloss.Color(colors.Base))

	b.WriteString(renderSection(titleStyle, "Image Details"))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Repository", img.Repository))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Tag", img.Tag))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "ID", img.ID))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Created", img.CreatedSince+" ago"))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Size", img.Size))
	b.WriteString("\n")
	if img.Digest != "" {
		b.WriteString(renderField(labelStyle, valueStyle, "Digest", img.Digest))
		b.WriteString("\n")
	}

	containerStyle := lipgloss.NewStyle().
		Width(width - 2).
		Height(height)

	return containerStyle.Render(b.String())
}

// renderVolumeDetails renders detailed info for a volume
func renderVolumeDetails(vol types.Volume, width, height int, colors theme.Colors) string {
	var b strings.Builder

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Maroon)).
		Background(lipgloss.Color(colors.Base)).
		Underline(true).
		MarginTop(1).
		MarginBottom(1).
		Bold(true)

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Lavender)).
		Bold(true).
		Width(15).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text)).
		Background(lipgloss.Color(colors.Base))

	b.WriteString(renderSection(titleStyle, "Volume Details"))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Name", vol.Name))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Driver", vol.Driver))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Mountpoint", vol.Mountpoint))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Scope", vol.Scope))
	b.WriteString("\n")
	if vol.Size != "" {
		b.WriteString(renderField(labelStyle, valueStyle, "Size", vol.Size))
		b.WriteString("\n")
	}

	containerStyle := lipgloss.NewStyle().
		Width(width - 2).
		Height(height)

	return containerStyle.Render(b.String())
}

// renderNetworkDetails renders detailed info for a network
func renderNetworkDetails(net types.Network, width, height int, colors theme.Colors) string {
	var b strings.Builder

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Maroon)).
		Background(lipgloss.Color(colors.Base)).
		Underline(true).
		MarginTop(1).
		MarginBottom(1).
		Bold(true)

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Lavender)).
		Bold(true).
		Width(15).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Text)).
		Background(lipgloss.Color(colors.Base))

	b.WriteString(renderSection(titleStyle, "Network Details"))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Name", net.Name))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "ID", net.ID[:12]))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Driver", net.Driver))
	b.WriteString("\n")
	b.WriteString(renderField(labelStyle, valueStyle, "Scope", net.Scope))
	b.WriteString("\n")
	if net.IPv6 != "" {
		b.WriteString(renderField(labelStyle, valueStyle, "IPv6", net.IPv6))
		b.WriteString("\n")
	}
	if net.Internal != "" {
		b.WriteString(renderField(labelStyle, valueStyle, "Internal", net.Internal))
		b.WriteString("\n")
	}

	containerStyle := lipgloss.NewStyle().
		Width(width - 2).
		Height(height)

	return containerStyle.Render(b.String())
}

// Helper functions
func renderSection(style lipgloss.Style, title string) string {
	return style.Render(title)
}

func renderField(labelStyle, valueStyle lipgloss.Style, label, value string) string {
	return labelStyle.Render(label+":") + valueStyle.Render(" "+value)
}

func formatStatus(state, status string, colors theme.Colors) string {
	stateStyle := lipgloss.NewStyle().Background(lipgloss.Color(colors.Base))

	switch state {
	case "running":
		stateStyle = stateStyle.Foreground(lipgloss.Color(colors.Green)).Bold(true)
	case "exited":
		stateStyle = stateStyle.Foreground(lipgloss.Color(colors.Red))
	case "paused":
		stateStyle = stateStyle.Foreground(lipgloss.Color(colors.Yellow))
	default:
		stateStyle = stateStyle.Foreground(lipgloss.Color(colors.Overlay0))
	}

	statusStyle := lipgloss.NewStyle().Background(lipgloss.Color(colors.Base))

	return stateStyle.Render(strings.ToUpper(state)) + statusStyle.Render(" "+status)
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
