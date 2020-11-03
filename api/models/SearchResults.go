package models

// SearchResults the results of the multi-entity search query.
type SearchResults struct {
	Epics   EpicSearchResults  `json:"epics,omitempty"`
	Stories StorySearchResults `json:"stories,omitempty"`
}

func (m *SearchResults) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *SearchResults) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
