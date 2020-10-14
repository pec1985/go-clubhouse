package models

import (
	"encoding/json"
	"time"
)

// Member Details about individual Clubhouse user within the Clubhouse organization that has issued the token.
type Member struct {
	// CreatedAt The time/date the Member was created.
	CreatedAt *time.Time `json:"created_at"`
	// CreatedWithoutInvite Whether this member was created as a placeholder entity.
	CreatedWithoutInvite bool `json:"created_without_invite"`
	// Disabled True/false boolean indicating whether the Member has been disabled within this Organization.
	Disabled bool `json:"disabled"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// GroupIDs The Member's group ids
	GroupIDs []string `json:"group_ids"`
	// ID The Member's ID in Clubhouse.
	ID      string  `json:"id"`
	Profile Profile `json:"profile"`
	// ReplacedBy The id of the member that replaces this one when merged.
	ReplacedBy string `json:"replaced_by"`
	// Role The Member's role in the Clubhouse organization.
	Role string `json:"role"`
	// UpdatedAt The time/date the Member was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

func (m *Member) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Member) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
