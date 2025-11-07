package docker

import (
	"dockerydo/internal/types"
	"encoding/json"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func GetVolumes() tea.Msg {
	cmd := exec.Command("docker", "volume", "ls", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		return types.ErrMsg(err)
	}

	var volumes []types.Volume
	lines := splitLines(string(output))
	for _, line := range lines {
		if line == "" {
			continue
		}
		var volume types.Volume
		err := json.Unmarshal([]byte(line), &volume)
		if err != nil {
			return types.ErrMsg(err)
		}
		volumes = append(volumes, volume)
	}
	return types.VolumesMsg(volumes)
}

func RmVolume(volumeId string, force bool) tea.Cmd {
	return func() tea.Msg {
		args := []string{"rm"}
		if force {
			args = append(args, "-f")
		}
		args = append(args, volumeId)
		cmd := exec.Command("docker", args...)
		out, err := cmd.CombinedOutput()
		resp := strings.TrimSpace(string(out))
		if err != nil {
			return ParseOpResponse(resp)
		}
		return types.DockerOpMsg{
			ID:           volumeId,
			Success:      true,
			ResourceType: types.VolumeResource,
		}
	}
}
