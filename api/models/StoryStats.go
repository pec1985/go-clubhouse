package models

import "encoding/json"

// StoryStats The stats object for Stories
type StoryStats struct {
	// NumRelatedDocuments The number of documents related to this Story.
	NumRelatedDocuments int64 `json:"num_related_documents"`
}

func (m *StoryStats) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *StoryStats) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
