package models

import "encoding/json"

type GetEpicStories struct {
	// IncludesDescription a true/false boolean indicating whether to return Stories with their descriptions.
	IncludesDescription bool `json:"includes_description,omitempty"`
}

func (m *GetEpicStories) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *GetEpicStories) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
