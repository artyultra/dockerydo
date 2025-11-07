package types

type Volume struct {
	Availability string `json:"Availability,omitzero"`
	Driver       string `json:"Driver,omitzero"`
	Group        string `json:"Group,omitzero"`
	Labels       string `json:"Labels"`
	Links        string `json:"Links,omitzero"`
	Mountpoint   string `json:"Mountpoint,omitzero"`
	Name         string `json:"Name,omitzero"`
	Scope        string `json:"Scope,omitzero"`
	Size         string `json:"Size,omitzero"`
	Status       string `json:"Status,omitzero"`
}
