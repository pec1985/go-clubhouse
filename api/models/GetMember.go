package models

import "encoding/json"

type GetMember struct {
	// The unique ID of the Organization to limit the lookup to.
	OrgPublicId string `json:"org-public-id"`
}

func (m *GetMember) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *GetMember) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
