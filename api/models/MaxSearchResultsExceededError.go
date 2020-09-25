package models

import "encoding/json"

type MaxSearchResultsExceededError struct {
	// The name for this type of error, `maximum-results-exceeded`
	Error string `json:"error"`
	// The maximum number of search results supported, `1000`
	MaximumResults int64 `json:"maximum-results"`
	// An explanatory message: "A maximum of 1000 search results are supported."
	Message string `json:"message"`
}

func (m *MaxSearchResultsExceededError) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *MaxSearchResultsExceededError) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
