package models

import "encoding/json"

type EpicSearchResults struct {
	Cursors []string `json:"cursors"`
	// A list of search results.
	Data []Epic `json:"data"`
	// The URL path and query string for the next page of search results.
	Next *string `json:"next"`
	// The total number of matches for the search query. The first 1000 matches can be paged through via the API.
	Total int64 `json:"total"`
}

func (m *EpicSearchResults) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EpicSearchResults) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
