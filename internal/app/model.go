package app

import (
	"dockerydo/internal/app/utils"
	"dockerydo/internal/docker"
	"dockerydo/internal/theme"
	"dockerydo/internal/types"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func NewModel() types.Model {
	return types.Model{
		ActiveTab:       types.TabContainers,
		PanelFoucus:     types.FocusListPanel,
		RightPanel:      types.PanelDetails,
		ShowErrPopup:    false,
		RefreshEnabled:  true,
		DetailsViewPort: viewport.New(0, 0),
		LogsViewPort:    viewport.New(0, 0),
		Theme:           theme.Dark,
	}
}

func Init(m types.Model) tea.Cmd {
	return tea.Batch(
		docker.GetContainers,
		docker.GetImages,
		docker.GetVolumes,
		docker.GetNetworks,
		utils.TickCmd(),
	)
}
