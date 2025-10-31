package docker

import (
	"dockerydo/internal/types"
	"strings"
)

func parseLabels(labelsStr string) *types.DockerLabels {
	if labelsStr == "" {
		return &types.DockerLabels{}
	}

	labels := &types.DockerLabels{
		RawLabels: make(map[string]string),
	}

	pairs := strings.Split(labelsStr, ",")

	for _, pair := range pairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		labels.RawLabels[key] = value

		switch key {
		case "com.docker.compose.depends_on":
			labels.ComposeDependsOn = value
		case "com.docker.compose.image":
			labels.ComposeImage = value
		case "com.docker.compose.oneoff":
			labels.ComposeOneoff = value
		case "com.docker.compose.service":
			labels.ComposeService = value
		case "com.docker.compose.config-hash":
			labels.ComposeConfigHash = value
		case "com.docker.compose.container-number":
			labels.ComposeContainerNum = value
		case "com.docker.compose.project.working_dir":
			labels.ComposeWorkingDir = value
		case "com.docker.compose.version":
			labels.ComposeVersion = value
		case "com.docker.compose.project":
			labels.ComposeProject = value
		case "com.docker.compose.config_files":
			labels.ComposeConfigFiles = value
		}

	}

	return labels

}
