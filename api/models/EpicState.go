package models

import (
	"encoding/json"
	"time"
)

// EpicState Epic State is any of the at least 3 columns. Epic States correspond to one of 3 types: Unstarted, Started, or Done.
type EpicState struct {
	// Color The hex color for this Epic State.
	Color string `json:"color"`
	// CreatedAt The time/date the Epic State was created.
	CreatedAt time.Time `json:"created_at"`
	// Description The description of what sort of Epics belong in that Epic State.
	Description string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ID The unique ID of the Epic State.
	ID int64 `json:"id"`
	// Name The Epic State's name.
	Name string `json:"name"`
	// Position The position that the Epic State is in, starting with 0 at the left.
	Position int64 `json:"position"`
	// Type The type of Epic State (Unstarted, Started, or Done)
	Type string `json:"type"`
	// UpdatedAt When the Epic State was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *EpicState) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EpicState) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
