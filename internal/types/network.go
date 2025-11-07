package types

type Network struct {
	CreatedAt string `json:"CreatedAt,omitzero"`
	Driver    string `json:"Driver,omitzero"`
	ID        string `json:"ID,omitzero"`
	IPv4      string `json:"IPv4,omitzero"`
	IPv6      string `json:"IPv6,omitzero"`
	Internal  string `json:"Internal,omitzero"`
	Labels    string `json:"Labels"`
	Name      string `json:"Name,omitzero"`
	Scope     string `json:"Scope,omitzero"`
}
