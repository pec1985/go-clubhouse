package models

import "encoding/json"

type SearchResults struct {
	Epics   EpicSearchResults  `json:"epics"`
	Stories StorySearchResults `json:"stories"`
}

func (m *SearchResults) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *SearchResults) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
