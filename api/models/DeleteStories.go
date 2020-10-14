package models

import "encoding/json"

type DeleteStories struct {
	// StoryIDs An array of IDs of Stories to delete.
	StoryIDs []int64 `json:"story_ids"`
}

func (m *DeleteStories) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *DeleteStories) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
