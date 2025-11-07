package docker

import (
	"dockerydo/internal/types"
	"encoding/json"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func GetImages() tea.Msg {
	cmd := exec.Command("docker", "images", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		return types.ErrMsg(err)
	}
	var images []types.Image
	lines := splitLines(string(output))
	for _, line := range lines {
		if line == "" {
			continue
		}
		var image types.Image
		err := json.Unmarshal([]byte(line), &image)
		if err != nil {
			return types.ErrMsg(err)
		}
		images = append(images, image)
	}
	return types.ImagesMsg(images)
}

func RmImage(imageId string, force bool) tea.Cmd {
	return func() tea.Msg {
		args := []string{"rm"}
		if force {
			args = append(args, "-f")
		}
		args = append(args, imageId)
		cmd := exec.Command("docker", args...)
		out, err := cmd.CombinedOutput()
		resp := strings.TrimSpace(string(out))
		if err != nil {
			return ParseOpResponse(resp)
		}

		return types.DockerOpMsg{
			ID:           imageId,
			Success:      true,
			ResourceType: types.ImageResource,
		}
	}
}
