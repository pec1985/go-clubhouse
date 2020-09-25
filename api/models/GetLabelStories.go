package models

import "encoding/json"

type GetLabelStories struct {
	// A true/false boolean indicating whether to return Stories with their descriptions.
	IncludesDescription bool `json:"includes_description"`
}

func (m *GetLabelStories) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *GetLabelStories) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
