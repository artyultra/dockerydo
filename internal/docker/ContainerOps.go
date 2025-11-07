package docker

import (
	"dockerydo/internal/types"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func GetContainers() tea.Msg {
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		return types.ErrMsg(err)
	}

	var containers []types.Container

	lines := splitLines(string(output))

	for _, line := range lines {
		if line == "" {
			continue
		}
		var container types.Container
		err := json.Unmarshal([]byte(line), &container)
		if err != nil {
			return types.ErrMsg(err)
		}
		newTimeUp := strings.ReplaceAll(container.RunningFor, " ago", "")
		container.RunningFor = newTimeUp
		if container.RawPorts != "" {
			container.Ports = parsePort(container.RawPorts)
		}
		container.Labels = parseLabels(container.RawLabels)

		containers = append(containers, container)
	}
	return types.ContainersMsg(containers)
}

func InspectContainer(container types.Container) tea.Cmd {
	return func() tea.Msg {
		return types.InspectMsg(container)
	}
}

func StartStopContainer(c types.Container) tea.Cmd {
	return func() tea.Msg {
		var dockerCmd string
		switch c.State {
		case "running":
			dockerCmd = "stop"
		case "exited":
			dockerCmd = "start"
		default:
			return types.ErrMsg(fmt.Errorf("\"Docker %s %s\" not a valid command", dockerCmd, c.ID))
		}

		cmd := exec.Command("docker", dockerCmd, c.ID)
		out, err := cmd.Output()
		resp := strings.TrimSpace(string(out))
		if resp != c.ID {
			return ParseOpResponse(resp)
		}
		if err != nil {
			return types.ErrMsg(err)
		}

		return types.DockerOpMsg{
			ResourceType: types.ContainerResource,
			ID:           c.ID,
			Success:      true,
		}
	}
}

func PauseUnpauseContainer(c types.Container) tea.Cmd {
	return func() tea.Msg {
		var dockerCmd string

		switch c.State {
		case "running":
			dockerCmd = "pause"
		case "paused":
			dockerCmd = "unpause"
		default:
			dockerCmd = "unpause"
		}

		cmd := exec.Command("docker", dockerCmd, c.ID)
		out, err := cmd.CombinedOutput()
		resp := strings.TrimSpace(string(out))
		if err != nil {
			return ParseOpResponse(resp)
		}

		return types.DockerOpMsg{
			ResourceType: types.ContainerResource,
			ID:           c.ID,
			Success:      true,
		}
	}
}

func RmContainer(c types.Container, f bool) tea.Cmd {
	return func() tea.Msg {
		var cmd *exec.Cmd

		switch f {
		case true:
			cmd = exec.Command("docker", "rm", "-f", c.ID)
		case false:
			cmd = exec.Command("docker", "rm", c.ID)
		}

		out, err := cmd.Output()
		if err != nil {
			return types.ErrMsg(err)
		}

		resp := strings.TrimSpace(string(out))
		if resp != c.ID {
			return ParseOpResponse(resp)
		}

		return types.DockerOpMsg{
			ResourceType: types.ContainerResource,
			ID:           c.ID,
			Success:      true,
		}
	}
}

func GetContainerLogs(c types.Container) tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("docker", "logs", "--tail", "100", c.ID)
		output, err := cmd.Output()
		if err != nil {
			return types.ErrMsg(err)
		}

		return types.LogsMsg{
			ID:  c.ID,
			Log: string(output),
		}
	}
}
