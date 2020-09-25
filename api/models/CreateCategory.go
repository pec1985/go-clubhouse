package models

import "encoding/json"

type CreateCategory struct {
	// The hex color to be displayed with the Category (for example, "#ff0000").
	Color string `json:"color"`
	// This field can be set to another unique ID. In the case that the Category has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// The name of the new Category.
	Name string `json:"name"`
	// The type of entity this Category is associated with; currently Milestone is the only type of Category.
	Type string `json:"type"`
}

func (m *CreateCategory) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateCategory) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
