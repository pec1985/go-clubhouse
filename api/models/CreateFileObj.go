package models

import "encoding/json"

type CreateFileObj struct {
	ContentType string `json:"content-type,omitempty"`
	Filename    string `json:"filename,omitempty"`
	Size        int64  `json:"size,omitempty"`
	Tempfile    File   `json:"tempfile,omitempty"`
}

func (m *CreateFileObj) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateFileObj) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
