package models

type CreateFileObj struct {
	ContentType string `json:"content-type,omitempty"`
	Filename    string `json:"filename,omitempty"`
	Size        int64  `json:"size,omitempty"`
	Tempfile    File   `json:"tempfile,omitempty"`
}

func (m *CreateFileObj) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateFileObj) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
