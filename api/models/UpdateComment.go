package models

type UpdateComment struct {
	// Text the updated comment text.
	Text string `json:"text,omitempty"`
}

func (m *UpdateComment) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateComment) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
