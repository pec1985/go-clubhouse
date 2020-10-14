package models

import (
	"encoding/json"
	"time"
)

// Category a Category can be used to associate Milestones.
type Category struct {
	// Archived a true/false boolean indicating if the Category has been archived.
	Archived bool `json:"archived"`
	// Color the hex color to be displayed with the Category (for example, "#ff0000").
	Color *string `json:"color"`
	// CreatedAt the time/date that the Category was created.
	CreatedAt time.Time `json:"created_at"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID this field can be set to another unique ID. In the case that the Category has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// ID the unique ID of the Category.
	ID int64 `json:"id"`
	// Name the name of the Category.
	Name string `json:"name"`
	// Type the type of entity this Category is associated with; currently Milestone is the only type of Category.
	Type string `json:"type"`
	// UpdatedAt the time/date that the Category was updated.
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
