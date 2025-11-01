package types

import "time"

type ErrMsg error

type ContainersMsg []Container

type InspectMsg Container

type TickMsg time.Time

type ContainerOpMsg struct {
	ContainerID string
	Success     bool
}

type OpFailedMsg struct {
	DaemonError string
	Error       string
}
