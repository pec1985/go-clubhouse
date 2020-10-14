package models

import "encoding/json"

// CreateLabelParams Request parameters for creating a Label on a Clubhouse story.
type CreateLabelParams struct {
	// Color The hex color to be displayed with the Label (for example, "#ff0000").
	Color string `json:"color"`
	// Description The description of the new Label.
	Description string `json:"description"`
	// ExternalID This field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id"`
	// Name The name of the new Label.
	Name string `json:"name"`
}

func (m *CreateLabelParams) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateLabelParams) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
