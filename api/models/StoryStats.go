package models

// StoryStats the stats object for Stories
type StoryStats struct {
	// NumRelatedDocuments the number of documents related to this Story.
	NumRelatedDocuments int64 `json:"num_related_documents,omitempty"`
}

func (m *StoryStats) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *StoryStats) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
