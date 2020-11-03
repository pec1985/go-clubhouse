package models

type GetProjectStories struct {
	// IncludesDescription a true/false boolean indicating whether to return Stories with their descriptions.
	IncludesDescription bool `json:"includes_description,omitempty"`
}

func (m *GetProjectStories) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *GetProjectStories) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
