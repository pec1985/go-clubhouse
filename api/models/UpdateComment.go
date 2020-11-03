package models

import "encoding/json"

type UpdateComment struct {
	// Text the updated comment text.
	Text string `json:"text,omitempty"`
}

func (m *UpdateComment) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateComment) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
