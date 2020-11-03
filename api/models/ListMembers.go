package models

import "encoding/json"

type ListMembers struct {
	// OrgPublicID the unique ID of the Organization to limit the list to.
	OrgPublicID string `json:"org-public-id,omitempty"`
}

func (m *ListMembers) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *ListMembers) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
