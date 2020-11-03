package models

import "time"

// EpicState epic State is any of the at least 3 columns. Epic States correspond to one of 3 types: Unstarted, Started, or Done.
type EpicState struct {
	// Color the hex color for this Epic State.
	Color string `json:"color,omitempty"`
	// CreatedAt the time/date the Epic State was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Description the description of what sort of Epics belong in that Epic State.
	Description string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ID the unique ID of the Epic State.
	ID int64 `json:"id,omitempty"`
	// Name the Epic State's name.
	Name string `json:"name,omitempty"`
	// Position the position that the Epic State is in, starting with 0 at the left.
	Position int64 `json:"position,omitempty"`
	// Type the type of Epic State (Unstarted, Started, or Done)
	Type string `json:"type,omitempty"`
	// UpdatedAt when the Epic State was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *EpicState) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *EpicState) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
