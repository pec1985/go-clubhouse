package models

import "encoding/json"

type Search struct {
	Include string `json:"include"`
	// Next The next page token.
	Next string `json:"next"`
	// PageSize The number of search results to include in a page. Minimum of 1 and maximum of 25.
	PageSize int64 `json:"page_size"`
	// Query See our help center article on [search operators](https://help.clubhouse.io/hc/en-us/articles/360000046646-Search-Operators)
	Query string `json:"query"`
}

func (m *Search) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Search) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
