package models

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
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Identity) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
