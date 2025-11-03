package types

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
)

type ViewMode int

const (
	ViewModeTable ViewMode = iota
	ViewModeDetailed
)

type Model struct {
	Table             table.Model
	Containers        []Container
	Err               error
	Width             int
	Height            int
	Mode              ViewMode
	SelectedContainer *Container
	ViewPort          viewport.Model
	RefreshEnabled    bool
	ShowErrPopup      bool
	ShowFailedOpPopup bool
	ShowConfirmPopup  bool
	FailedOpPopUpMsg  string
	ErrPopUpMsg       string
	ConfirmPopUpMsg   string
}
