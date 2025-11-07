package types

type Image struct {
	Containers   string `json:"Containers,omitzero"`
	CreatedAt    string `json:"CreatedAt,omitzero"`
	CreatedSince string `json:"CreatedSince,omitzero"`
	Digest       string `json:"Digest,omitzero"`
	ID           string `json:"ID,omitzero"`
	Repository   string `json:"Repository,omitzero"`
	SharedSize   string `json:"SharedSize,omitzero"`
	Size         string `json:"Size,omitzero"`
	Tag          string `json:"Tag,omitzero"`
	UniqueSize   string `json:"UniqueSize,omitzero"`
	VirtualSize  string `json:"VirtualSize,omitzero"`
}
