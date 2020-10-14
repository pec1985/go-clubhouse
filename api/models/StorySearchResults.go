package models

import "encoding/json"

// StorySearchResults The results of the Story search query.
type StorySearchResults struct {
	Cursors []string `json:"cursors"`
	// Data A list of search results.
	Data []Story `json:"data"`
	// Next The URL path and query string for the next page of search results.
	Next *string `json:"next"`
	// Total The total number of matches for the search query. The first 1000 matches can be paged through via the API.
	Total int64 `json:"total"`
}

func (m *StorySearchResults) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *StorySearchResults) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
