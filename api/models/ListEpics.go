package models

import "encoding/json"

type ListEpics struct {
	// A true/false boolean indicating whether to return Epics with their descriptions.
	IncludesDescription bool `json:"includes_description"`
}

func (m *ListEpics) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *ListEpics) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
