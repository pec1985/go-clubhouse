package models

import "encoding/json"

// UpdateEntityTemplate Request parameters for changing either a template's name or any of
//   the attributes it is designed to pre-populate.
type UpdateEntityTemplate struct {
	// Name The updated template name.
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
