package types

type Container struct {
	Command      string `json:"Command"`
	CreatedAt    string `json:"CreatedAt"`
	ID           string `json:"ID"`
	Image        string `json:"Image"`
	RawLabels    string `json:"Labels"`
	Labels       *DockerLabels
	LocalVolumes string `json:"LocalVolumes"`
	Mounts       string `json:"Mounts"`
	Names        string `json:"Names"`
	Networks     string `json:"Networks"`
	RawPorts     string `json:"Ports"`
	Ports        Ports
	RunningFor   string `json:"RunningFor"`
	Size         string `json:"Size"`
	State        string `json:"State"`
	Status       string `json:"Status"`
	InternalPort string // parsed from Ports
	ExternalPort string // -------^^^^------
}

type DockerLabels struct {
	ComposeService      string
	ComposeProject      string
	ComposeVersion      string
	ComposeConfigFiles  string
	ComposeWorkingDir   string
	ComposeConfigHash   string
	ComposeContainerNum string
	ComposeDependsOn    string
	ComposeImage        string
	ComposeOneoff       string
	RawLabels           map[string]string // raw key-value pairs
}

type PortMap struct {
	Ipv6          string
	Ipv4          string
	InternalRange string
	ExternalRange string
	Protocol      string
}

type Ports []PortMap
