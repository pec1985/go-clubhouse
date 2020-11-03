package models

// CreateEntityTemplate request paramaters for creating an entirely new entity template.
type CreateEntityTemplate struct {
	// AuthorID the id of the user creating this template.
	AuthorID string `json:"author_id,omitempty"`
	// Name the name of the new entity template
	Name          string              `json:"name,omitempty"`
	StoryContents CreateStoryContents `json:"story_contents,omitempty"`
}

func (m *CreateEntityTemplate) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateEntityTemplate) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
