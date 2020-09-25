package models

import (
	"encoding/json"
	"time"
)

type Icon struct {
	// The time/date that the Icon was created.
	CreatedAt time.Time `json:"created_at"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique ID of the Icon.
	Id string `json:"id"`
	// The time/date that the Icon was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The URL of the Icon.
	Url string `json:"url"`
}

func (m *Icon) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Icon) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
