package models

import (
	"encoding/json"
	"time"
)

// icons are used to attach images to Organizations, Members, and Loading screens in the Clubhouse web application.
type Icon struct {
	// CreatedAt the time/date that the Icon was created.
	CreatedAt time.Time `json:"created_at"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ID the unique ID of the Icon.
	ID string `json:"id"`
	// UpdatedAt the time/date that the Icon was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// URL the URL of the Icon.
	URL string `json:"url"`
}

func (m *Icon) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Icon) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
