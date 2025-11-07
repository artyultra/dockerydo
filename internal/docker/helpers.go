package docker

import (
	"dockerydo/internal/types"
	"regexp"
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

func parsePort(portMapping string) types.Ports {
	const (
		ipv6Pattern          = `\[(?P<ipv6>[:a-fA-F0-9]+)\]`
		ipv4Pattern          = `(?P<ipv4>[\d.]+)`
		addressPattern       = `(?:(?:` + ipv6Pattern + `|` + ipv4Pattern + `):)?`
		hostPortPattern      = `(?P<host_port>\d+(?:-\d+)?)`
		containerPortPattern = `(?:->(?P<container_port>\d+(?:-\d+)?))?`
		protocolPattern      = `/(?P<protocol>\w+)`
	)
	portRegex := regexp.MustCompile(
		`^` +
			addressPattern +
			hostPortPattern +
			containerPortPattern +
			protocolPattern +
			`$`,
	)
	portEntries := strings.Split(portMapping, ",")
	var ports types.Ports

	for _, portEntry := range portEntries {
		var portMap types.PortMap
		portEntry = strings.TrimSpace(portEntry)
		matches := portRegex.FindStringSubmatch(portEntry)
		if matches == nil {
			return types.Ports{}
		}

		names := portRegex.SubexpNames()

		result := make(map[string]string)
		for i, name := range names {
			if i != 0 && name != "" {
				result[name] = matches[i]
			}
		}

		if result["ipv6"] != "" {
			portMap.Ipv6 = result["ipv6"]
		}
		if result["ipv4"] != "" {
			portMap.Ipv4 = result["ipv4"]
		}
		if result["host_port"] != "" {
			portMap.InternalRange = result["host_port"]
		}
		if result["container_port"] != "" {
			portMap.ExternalRange = result["container_port"]
		}
		if result["protocol"] != "" {
			portMap.Protocol = result["protocol"]
		}
		ports = append(ports, portMap)

	}

	return ports
}

func ParseOpResponse(resp string) types.OpFailedMsg {
	var result types.OpFailedMsg

	// Remove the "Error response from daemon: " prefix
	msg := strings.TrimPrefix(resp, "Error response from daemon: ")
	msg = strings.TrimPrefix(msg, "Error: ")

	// Find and truncate container IDs (64 character hex strings)
	// Match pattern: 64 consecutive hex characters
	re := regexp.MustCompile(`\b[0-9a-f]{64}\b`)
	msg = re.ReplaceAllStringFunc(msg, func(id string) string {
		return truncate(id) // Truncate to 12 characters
	})

	result.DaemonError = msg
	result.Error = ""

	return result
}

func truncate(s string) string {
	if len(s) <= 12 {
		return s
	}
	return s[:12] + "..."
}

// splitLines splits a string by newlines and filters out empty lines
func splitLines(s string) []string {
	lines := strings.Split(s, "\n")
	var result []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}
	return result
}
