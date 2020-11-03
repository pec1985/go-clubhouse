package models

type ListEpics struct {
	// IncludesDescription a true/false boolean indicating whether to return Epics with their descriptions.
	IncludesDescription bool `json:"includes_description,omitempty"`
}

func (m *ListEpics) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *ListEpics) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
