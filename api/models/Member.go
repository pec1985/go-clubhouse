package models

import (
	"encoding/json"
	"time"
)

type Member struct {
	Profile Profile `json:"profile"`
	// The id of the member that replaces this one when merged.
	ReplacedBy string `json:"replaced_by"`
	// The Member's role in the Clubhouse organization.
	Role string `json:"role"`
	// The time/date the Member was created.
	CreatedAt *time.Time `json:"created_at"`
	// The Member's group ids
	GroupIds []string `json:"group_ids"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The Member's ID in Clubhouse.
	Id string `json:"id"`
	// The time/date the Member was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// Whether this member was created as a placeholder entity.
	CreatedWithoutInvite bool `json:"created_without_invite"`
	// True/false boolean indicating whether the Member has been disabled within this Organization.
	Disabled bool `json:"disabled"`
}

func (m *Member) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Member) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
