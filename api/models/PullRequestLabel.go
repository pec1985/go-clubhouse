package models

import "encoding/json"

// PullRequestLabel corresponds to a VCS Label associated with a Pull Request.
type PullRequestLabel struct {
	// Color the color of the VCS label.
	Color string `json:"color"`
	// Description the description of the VCS label.
	Description *string `json:"description"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ID the unique ID of the VCS Label.
	ID int64 `json:"id"`
	// Name the name of the VCS label.
	Name string `json:"name"`
}

func (m *PullRequestLabel) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *PullRequestLabel) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
