package models

type GetMember struct {
	// OrgPublicID the unique ID of the Organization to limit the lookup to.
	OrgPublicID string `json:"org-public-id,omitempty"`
}

func (m *GetMember) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *GetMember) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
