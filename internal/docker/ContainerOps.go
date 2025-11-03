package docker

import (
	"dockerydo/internal/types"
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

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

		return types.ContainerOpMsg{ContainerID: c.ID, Success: true}
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

		return types.ContainerOpMsg{ContainerID: c.ID, Success: true}
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

		return types.ContainerOpMsg{ContainerID: c.ID, Success: true}
	}
}
