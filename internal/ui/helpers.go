package ui

import "dockerydo/internal/types"

// FormatPortsForTable formats port information for display in the table
func FormatPortsForTable(container types.Container) string {
	if len(container.Ports) == 0 {
		// Fall back to raw ports string if available
		if container.RawPorts != "" {
			return container.RawPorts
		}
		// Fall back to combining External/Internal if available
		if container.ExternalPort != "" && container.InternalPort != "" {
			return container.ExternalPort + "->" + container.InternalPort
		}
		if container.InternalPort != "" {
			return container.InternalPort
		}
		return "-"
	}

	// Format using the Ports slice for better formatting
	portStrings := []string{}
	for _, pm := range container.Ports {
		if pm.ExternalRange != "" && pm.InternalRange != "" {
			proto := pm.Protocol
			if proto == "" {
				proto = "tcp"
			}
			// Format like Docker: 0.0.0.0:8080->80/tcp
			if pm.Ipv4 != "" {
				portStrings = append(portStrings, pm.Ipv4+":"+pm.ExternalRange+"->"+pm.InternalRange+"/"+proto)
			} else if pm.Ipv6 != "" {
				portStrings = append(portStrings, "["+pm.Ipv6+"]:"+pm.ExternalRange+"->"+pm.InternalRange+"/"+proto)
			} else {
				portStrings = append(portStrings, pm.ExternalRange+"->"+pm.InternalRange+"/"+proto)
			}
		} else if pm.InternalRange != "" {
			// Internal port only (not exposed)
			proto := pm.Protocol
			if proto == "" {
				proto = "tcp"
			}
			portStrings = append(portStrings, pm.InternalRange+"/"+proto)
		}
	}

	if len(portStrings) == 0 {
		return "-"
	}

	// Join multiple ports with comma
	result := ""
	for i, ps := range portStrings {
		if i > 0 {
			result += ", "
		}
		result += ps
		// Limit display to avoid overwhelming the table
		if i >= 1 && len(portStrings) > 2 {
			result += " ..."
			break
		}
	}
	return result
}
