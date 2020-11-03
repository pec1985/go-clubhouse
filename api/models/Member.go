package models

import (
	"encoding/json"
	"time"
)

// Member details about individual Clubhouse user within the Clubhouse organization that has issued the token.
type Member struct {
	// CreatedAt the time/date the Member was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// CreatedWithoutInvite whether this member was created as a placeholder entity.
	CreatedWithoutInvite bool `json:"created_without_invite,omitempty"`
	// Disabled true/false boolean indicating whether the Member has been disabled within this Organization.
	Disabled bool `json:"disabled,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// GroupIDs the Member's group ids
	GroupIDs []string `json:"group_ids,omitempty"`
	// ID the Member's ID in Clubhouse.
	ID      string  `json:"id,omitempty"`
	Profile Profile `json:"profile,omitempty"`
	// ReplacedBy the id of the member that replaces this one when merged.
	ReplacedBy string `json:"replaced_by,omitempty"`
	// Role the Member's role in the Clubhouse organization.
	Role string `json:"role,omitempty"`
	// UpdatedAt the time/date the Member was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (m *Member) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Member) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
