package models

import "encoding/json"

// SearchResults the results of the multi-entity search query.
type SearchResults struct {
	Epics   EpicSearchResults  `json:"epics,omitempty"`
	Stories StorySearchResults `json:"stories,omitempty"`
}

func (m *SearchResults) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *SearchResults) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
