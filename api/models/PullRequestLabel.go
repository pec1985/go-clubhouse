package models

import "encoding/json"

// PullRequestLabel Corresponds to a VCS Label associated with a Pull Request.
type PullRequestLabel struct {
	// Color The color of the VCS label.
	Color string `json:"color"`
	// Description The description of the VCS label.
	Description *string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ID The unique ID of the VCS Label.
	ID int64 `json:"id"`
	// Name The name of the VCS label.
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
