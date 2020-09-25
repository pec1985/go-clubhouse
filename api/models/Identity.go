package models

import "encoding/json"

type Identity struct {
	// The type of Identity; currently only type is github.
	Type *string `json:"type"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// This is your login in VCS.
	Name *string `json:"name"`
}

func (m *Identity) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Identity) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
