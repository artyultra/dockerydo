package ui

import (
	"dockerydo/internal/types"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func RenderBaseView(m types.Model) string {
	headerHight := 3
	footerHight := 2
	contentHight := m.Height - headerHight - footerHight
	if contentHight < 5 {
		contentHight = 5
	}

	leftWidth := int(float64(m.Width) * 0.4)
	if leftWidth < 10 {
		leftWidth = 10
	}
	rightWidth := m.Width - leftWidth
	if rightWidth < 10 {
		rightWidth = 10
	}

	header := renderTabBar(m)
	leftPanel := renderListPanel(m, leftWidth, contentHight)
	rightPanel := renderMainPanel(m, rightWidth, contentHight)
	footer := renderFooter(m)

	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftPanel,
		rightPanel,
	)

	layout := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		content,
		footer,
	)

	return layout
}

func renderTabBar(m types.Model) string {
	tabs := []struct {
		names  string
		tab    types.TabType
		number int
	}{
		{"Containers", types.TabContainer, 1},
	}

	var renderedTabs []string
	for _, tab := range tabs {
		style := lipgloss.NewStyle().
			Padding(2)
		if m.ActiveTab == tab.tab {
			style = style.Foreground(lipgloss.Color(m.Theme.Crust)).
				Background(lipgloss.Color(m.Theme.Mauve)).
				Bold(true)
		} else {
			style = style.
				Foreground(lipgloss.Color(m.Theme.Subtext0)).
				Background(lipgloss.Color(m.Theme.Surface0))
		}

		label := fmt.Sprintf("[%d] %s", tab.number, tab.names)
		renderedTabs = append(renderedTabs, style.Render(label))
	}

	tabBar := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	containerStyle := lipgloss.NewStyle().
		Width(m.Width).
		Background(lipgloss.Color(m.Theme.Base)).
		Padding(1, 0)

	return containerStyle.Render(tabBar)
}

func renderFooter(m types.Model) string {
	var keys string

	switch m.ActiveTab {

	case types.TabContainers:
		if m.RightPanel == types.PanelLogs {
			keys = "esc: back • ↑↓: scroll • q: quit"
		} else {
			keys = "←→: tabs • ↑↓: navigate • enter: logs • s: start/stop • p: pause • d/D: remove • q: quit"
		}
	case types.TabImages:
		keys = "←→: tabs • ↑↓: navigate • d/D: remove • q: quit"
	case types.TabVolumes:
		keys = "←→: tabs • ↑↓: navigate • d: remove • q: quit"
	case types.TabNetworks:
		keys = "←→: tabs • ↑↓: navigate • d: remove • q: quit"
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(m.Theme.Subtext0)).
		Background(lipgloss.Color(m.Theme.Surface0)).
		Padding(0, 2).
		Width(m.Width)
	return style.Render(keys)
}
