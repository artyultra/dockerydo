package docker

import (
	"dockerydo/internal/types"
	"encoding/json"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func GetContainers() tea.Msg {
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{json .}}")
	out, err := cmd.Output()
	if err != nil {
		return types.ErrMsg(err)
	}

	var containers []types.Container

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

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
