package models

import "time"

// repository refers to a VCS repository.
type Repository struct {
	// CreatedAt the time/date the Repository was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ExternalID the VCS unique identifier for the Repository.
	ExternalID *string `json:"external_id,omitempty"`
	// FullName the full name of the VCS repository.
	FullName *string `json:"full_name,omitempty"`
	// ID the ID associated to the VCS repository in Clubhouse.
	ID *int64 `json:"id,omitempty"`
	// Name the shorthand name of the VCS repository.
	Name *string `json:"name,omitempty"`
	// Type the type of Repository. Currently this can only be "github".
	Type string `json:"type,omitempty"`
	// UpdatedAt the time/date the Repository was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// URL the URL of the Repository.
	URL *string `json:"url,omitempty"`
}

func (m *Repository) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Repository) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
