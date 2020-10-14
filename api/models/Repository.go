package models

import (
	"encoding/json"
	"time"
)

// Repository refers to a VCS repository.
type Repository struct {
	// CreatedAt The time/date the Repository was created.
	CreatedAt *time.Time `json:"created_at"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID The VCS unique identifier for the Repository.
	ExternalID *string `json:"external_id"`
	// FullName The full name of the VCS repository.
	FullName *string `json:"full_name"`
	// ID The ID associated to the VCS repository in Clubhouse.
	ID *int64 `json:"id"`
	// Name The shorthand name of the VCS repository.
	Name *string `json:"name"`
	// Type The type of Repository. Currently this can only be "github".
	Type string `json:"type"`
	// UpdatedAt The time/date the Repository was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// URL The URL of the Repository.
	URL *string `json:"url"`
}

func (m *Repository) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Repository) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
