package models

import "encoding/json"

type CreateStories struct {
	// Stories An array of stories to be created.
	Stories []CreateStoryParams `json:"stories"`
}

func (m *CreateStories) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateStories) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
