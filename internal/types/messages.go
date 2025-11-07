package types

import "time"

type ErrMsg error

type ContainersMsg []Container

type ImagesMsg []Image

type VolumesMsg []Volume

type NetworksMsg []Network

type InspectMsg Container

type TickMsg time.Time

type DockerOpMsg struct {
	ResourceType int
	ID           string
	Success      bool
}

const (
	ContainerResource = iota
	ImageResource
	VolumeResource
	NetworkResource
)

type OpFailedMsg struct {
	DaemonError string
	Error       string
}

type ConfirmMsg string

type LogsMsg struct {
	ID  string
	Log string
}
