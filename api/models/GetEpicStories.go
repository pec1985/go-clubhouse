package models

type GetEpicStories struct {
	// IncludesDescription a true/false boolean indicating whether to return Stories with their descriptions.
	IncludesDescription bool `json:"includes_description,omitempty"`
}

func (m *GetEpicStories) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *GetEpicStories) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
