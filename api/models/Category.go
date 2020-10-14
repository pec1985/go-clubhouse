package models

import (
	"encoding/json"
	"time"
)

// Category A Category can be used to associate Milestones.
type Category struct {
	// Archived A true/false boolean indicating if the Category has been archived.
	Archived bool `json:"archived"`
	// Color The hex color to be displayed with the Category (for example, "#ff0000").
	Color *string `json:"color"`
	// CreatedAt The time/date that the Category was created.
	CreatedAt time.Time `json:"created_at"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID This field can be set to another unique ID. In the case that the Category has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// ID The unique ID of the Category.
	ID int64 `json:"id"`
	// Name The name of the Category.
	Name string `json:"name"`
	// Type The type of entity this Category is associated with; currently Milestone is the only type of Category.
	Type string `json:"type"`
	// UpdatedAt The time/date that the Category was updated.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Category) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Category) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
