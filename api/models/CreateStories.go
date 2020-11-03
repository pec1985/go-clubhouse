package models

type CreateStories struct {
	// Stories an array of stories to be created.
	Stories []CreateStoryParams `json:"stories,omitempty"`
}

func (m *CreateStories) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateStories) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
