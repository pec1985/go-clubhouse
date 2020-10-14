package models

import "encoding/json"

type ListLabels struct {
	// Slim a true/false boolean indicating if the slim versions of the Label should be returned.
	Slim bool `json:"slim"`
}

func (m *ListLabels) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *ListLabels) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
