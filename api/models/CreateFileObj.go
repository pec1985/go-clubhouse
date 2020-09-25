package models

import "encoding/json"

type CreateFileObj struct {
	ContentType string `json:"content-type"`
	Filename    string `json:"filename"`
	Size        int64  `json:"size"`
	Tempfile    File   `json:"tempfile"`
}

func (m *CreateFileObj) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateFileObj) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
