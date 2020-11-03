package models

import "encoding/json"

type GetIterationStories struct {
	// IncludesDescription a true/false boolean indicating whether to return Stories with their descriptions.
	IncludesDescription bool `json:"includes_description,omitempty"`
}

func (m *GetIterationStories) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *GetIterationStories) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
