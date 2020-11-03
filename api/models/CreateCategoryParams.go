package models

import "encoding/json"

// CreateCategoryParams request parameters for creating a Category with a Milestone.
type CreateCategoryParams struct {
	// Color the hex color to be displayed with the Category (for example, "#ff0000").
	Color string `json:"color,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Category has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`
	// Name the name of the new Category.
	Name string `json:"name,omitempty"`
}

func (m *CreateCategoryParams) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateCategoryParams) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
