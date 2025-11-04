package types

import (
	"dockerydo/internal/theme"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
)

type TabType int

const (
	TabContainers TabType = iota + 1
	TabImages
	TabVolumes
	TabNetworks
)

type PanelType int

const (
	PanelDetails PanelType = iota
	PanelLogs
)

type PanelFocus int

const (
	FocusListPanel PanelFocus = iota
	FocusMainPanel
)

type ViewMode int

const (
	ViewModeTable ViewMode = iota
	ViewModeDetailed
)

type Model struct {
	// Navigation
	ActiveTab   TabType
	RightPanel  PanelType
	PanelFoucus PanelFocus

	// Data
	Containers []Container
	Images     []struct{}
	Volumes    []struct{}
	Networks   []struct{}

	// List Cursors
	ContainersCursor int
	ImageCursor      int
	VolumeCursor     int
	NetworkCursor    int

	// Scroll position for list
	ContainerScrol int
	ImageScroll    int
	VolumeScroll   int
	NetworkScroll  int

	// ViewPorts
	DetailsViewPort viewport.Model
	LogsViewPort    viewport.Model

	// UI State
	Width          int
	Height         int
	RefreshEnabled bool
	Theme          theme.Colors

	// PopUps
	ShowErrPopup      bool
	ShowFailedOpPopup bool
	ShowConfirmPopup  bool
	FailedOpPopUpMsg  string
	ErrPopUpMsg       string
	ConfirmPopUpMsg   string

	// Errors
	Err error

	Table             table.Model
	Mode              ViewMode
	SelectedContainer *Container
	ViewPort          viewport.Model
}
