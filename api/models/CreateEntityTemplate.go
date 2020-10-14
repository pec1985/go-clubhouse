package models

import "encoding/json"

// CreateEntityTemplate Request paramaters for creating an entirely new entity template.
type CreateEntityTemplate struct {
	// AuthorID The id of the user creating this template.
	AuthorID string `json:"author_id"`
	// Name The name of the new entity template
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
