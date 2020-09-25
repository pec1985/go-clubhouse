package models

import "encoding/json"

type GetProjectStories struct {
	// A true/false boolean indicating whether to return Stories with their descriptions.
	IncludesDescription bool `json:"includes_description"`
}

func (m *GetProjectStories) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *GetProjectStories) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
