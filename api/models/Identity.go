package models

import "encoding/json"

// Identity the Identity of the VCS user that authored the Commit.
type Identity struct {
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// Name this is your login in VCS.
	Name *string `json:"name,omitempty"`
	// Type the type of Identity; currently only type is github.
	Type *string `json:"type,omitempty"`
}

func (m *Identity) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Identity) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
