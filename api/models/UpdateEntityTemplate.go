package models

import "encoding/json"

type UpdateEntityTemplate struct {
	// The updated template name.
	Name          string              `json:"name"`
	StoryContents CreateStoryContents `json:"story_contents"`
}

func (m *UpdateEntityTemplate) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateEntityTemplate) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
