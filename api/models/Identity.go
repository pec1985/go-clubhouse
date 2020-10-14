package models

import "encoding/json"

// Identity The Identity of the VCS user that authored the Commit.
type Identity struct {
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// Name This is your login in VCS.
	Name *string `json:"name"`
	// Type The type of Identity; currently only type is github.
	Type *string `json:"type"`
}

func (m *Identity) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Identity) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
