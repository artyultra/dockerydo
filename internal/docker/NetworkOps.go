package docker

import (
	"dockerydo/internal/types"
	"encoding/json"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func GetNetworks() tea.Msg {
	var networks []types.Network
	cmd := exec.Command("docker", "network", "ls", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		return types.ErrMsg(err)
	}
	lines := splitLines(string(output))
	for _, line := range lines {
		if line == "" {
			continue
		}
		var net types.Network
		err := json.Unmarshal([]byte(line), &net)
		if err != nil {
			return types.ErrMsg(err)
		}
		networks = append(networks, net)

	}

	return types.VolumesMsg([]types.Volume{})
}

func RmNetwork(netId string, force bool) tea.Cmd {
	return func() tea.Msg {
		args := []string{"rm"}
		if force {
			args = append(args, "-f")
		}
		args = append(args, netId)
		cmd := exec.Command("docker", args...)
		out, err := cmd.CombinedOutput()
		resp := strings.TrimSpace(string(out))
		if err != nil {
			return ParseOpResponse(resp)
		}
		return types.DockerOpMsg{
			ID:           netId,
			Success:      true,
			ResourceType: types.NetworkResource,
		}
	}
}
