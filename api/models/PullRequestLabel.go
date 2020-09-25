package models

import "encoding/json"

type PullRequestLabel struct {
	// The color of the VCS label.
	Color string `json:"color"`
	// The description of the VCS label.
	Description *string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique ID of the VCS Label.
	Id int64 `json:"id"`
	// The name of the VCS label.
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
