package models

import (
	"encoding/json"
	"time"
)

type EpicState struct {
	// The description of what sort of Epics belong in that Epic State.
	Description string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique ID of the Epic State.
	Id int64 `json:"id"`
	// The position that the Epic State is in, starting with 0 at the left.
	Position int64 `json:"position"`
	// When the Epic State was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The hex color for this Epic State.
	Color string `json:"color"`
	// The Epic State's name.
	Name string `json:"name"`
	// The type of Epic State (Unstarted, Started, or Done)
	Type string `json:"type"`
	// The time/date the Epic State was created.
	CreatedAt time.Time `json:"created_at"`
}

func (m *EpicState) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EpicState) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
