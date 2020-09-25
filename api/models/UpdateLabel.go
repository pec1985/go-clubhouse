package models

import "encoding/json"

type UpdateLabel struct {
	// A true/false boolean indicating if the Label has been archived.
	Archived bool `json:"archived"`
	// The hex color to be displayed with the Label (for example, "#ff0000").
	Color *string `json:"color"`
	// The new description of the label.
	Description string `json:"description"`
	// The new name of the label.
	Name string `json:"name"`
}

func (m *UpdateLabel) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateLabel) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
