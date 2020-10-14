package models

import (
	"encoding/json"
	"time"
)

// repository refers to a VCS repository.
type Repository struct {
	// CreatedAt the time/date the Repository was created.
	CreatedAt *time.Time `json:"created_at"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID the VCS unique identifier for the Repository.
	ExternalID *string `json:"external_id"`
	// FullName the full name of the VCS repository.
	FullName *string `json:"full_name"`
	// ID the ID associated to the VCS repository in Clubhouse.
	ID *int64 `json:"id"`
	// Name the shorthand name of the VCS repository.
	Name *string `json:"name"`
	// Type the type of Repository. Currently this can only be "github".
	Type string `json:"type"`
	// UpdatedAt the time/date the Repository was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// URL the URL of the Repository.
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
