package models

import "encoding/json"

type CreateLabelParams struct {
	// The hex color to be displayed with the Label (for example, "#ff0000").
	Color string `json:"color"`
	// The description of the new Label.
	Description string `json:"description"`
	// This field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// The name of the new Label.
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
