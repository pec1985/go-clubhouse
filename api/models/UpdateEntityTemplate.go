package models

// UpdateEntityTemplate request parameters for changing either a template's name or any of
//   the attributes it is designed to pre-populate.
type UpdateEntityTemplate struct {
	// Name the updated template name.
	Name          string              `json:"name,omitempty"`
	StoryContents CreateStoryContents `json:"story_contents,omitempty"`
}

func (m *UpdateEntityTemplate) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateEntityTemplate) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
