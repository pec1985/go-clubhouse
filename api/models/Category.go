package models

import (
	"encoding/json"
	"time"
)

type Category struct {
	// A true/false boolean indicating if the Category has been archived.
	Archived bool `json:"archived"`
	// The hex color to be displayed with the Category (for example, "#ff0000").
	Color *string `json:"color"`
	// The time/date that the Category was created.
	CreatedAt time.Time `json:"created_at"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// This field can be set to another unique ID. In the case that the Category has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id"`
	// The unique ID of the Category.
	Id int64 `json:"id"`
	// The name of the Category.
	Name string `json:"name"`
	// The type of entity this Category is associated with; currently Milestone is the only type of Category.
	Type string `json:"type"`
	// The time/date that the Category was updated.
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
