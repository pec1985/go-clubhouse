package models

type DeleteStories struct {
	// StoryIDs an array of IDs of Stories to delete.
	StoryIDs []int64 `json:"story_ids,omitempty"`
}

func (m *DeleteStories) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *DeleteStories) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
