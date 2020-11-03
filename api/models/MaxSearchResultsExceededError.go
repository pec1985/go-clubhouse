package models

// MaxSearchResultsExceededError error returned when total maximum supported results have been reached.
type MaxSearchResultsExceededError struct {
	// Error the name for this type of error, `maximum-results-exceeded`
	Error string `json:"error,omitempty"`
	// MaximumResults the maximum number of search results supported, `1000`
	MaximumResults int64 `json:"maximum-results,omitempty"`
	// Message an explanatory message: "A maximum of 1000 search results are supported."
	Message string `json:"message,omitempty"`
}

func (m *MaxSearchResultsExceededError) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *MaxSearchResultsExceededError) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
