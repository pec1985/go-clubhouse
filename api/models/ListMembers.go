package models

type ListMembers struct {
	// OrgPublicID the unique ID of the Organization to limit the list to.
	OrgPublicID string `json:"org-public-id,omitempty"`
}

func (m *ListMembers) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *ListMembers) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
