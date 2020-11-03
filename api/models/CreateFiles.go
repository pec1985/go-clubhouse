package models

type CreateFiles struct {
	// StoryID the story ID that this file will be associated with.
	StoryID int64 `json:"story_id,omitempty"`
}

func (m *CreateFiles) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateFiles) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
