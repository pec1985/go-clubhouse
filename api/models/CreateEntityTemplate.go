package models

import "encoding/json"

type CreateEntityTemplate struct {
	// The id of the user creating this template.
	AuthorId string `json:"author_id"`
	// The name of the new entity template
	Name          string              `json:"name"`
	StoryContents CreateStoryContents `json:"story_contents"`
}

func (m *CreateEntityTemplate) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateEntityTemplate) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
