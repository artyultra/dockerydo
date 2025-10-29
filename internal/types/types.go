package types

import (
	"time"

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
}

type Container struct {
	Command      string `json:"Command"`
	CreatedAt    string `json:"CreatedAt"`
	ID           string `json:"ID"`
	Image        string `json:"Image"`
	Labels       string `json:"Labels"`
	LocalVolumes string `json:"LocalVolumes"`
	Mounts       string `json:"Mounts"`
	Names        string `json:"Names"`
	Networks     string `json:"Networks"`
	Ports        string `json:"Ports"`
	RunningFor   string `json:"RunningFor"`
	Size         string `json:"Size"`
	State        string `json:"State"`
	Status       string `json:"Status"`
	InternalPort string // parsed from Ports
	ExternalPort string // -------^^^^------
}

type DockerLabels struct {
	ComposeService      string
	ComposeProject      string
	ComposeVersion      string
	ComposeConfigFiles  string
	ComposeWorkingDir   string
	ComposeConfigHash   string
	ComposeContainerNum string
	ComposeDependsOn    string
	ComposeImage        string
	ComposeOneoff       string
	RawLabels           map[string]string // raw key-value pairs
}

type ErrMsg error

type ContainersMsg []Container

type InspectMsg Container

type TickMsg time.Time
