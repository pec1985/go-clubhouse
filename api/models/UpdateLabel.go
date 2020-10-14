package models

import "encoding/json"

type UpdateLabel struct {
	// Archived a true/false boolean indicating if the Label has been archived.
	Archived bool `json:"archived"`
	// Color the hex color to be displayed with the Label (for example, "#ff0000").
	Color *string `json:"color"`
	// Description the new description of the label.
	Description string `json:"description"`
	// Name the new name of the label.
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
