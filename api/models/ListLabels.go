package models

type ListLabels struct {
	// Slim a true/false boolean indicating if the slim versions of the Label should be returned.
	Slim bool `json:"slim,omitempty"`
}

func (m *ListLabels) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *ListLabels) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
