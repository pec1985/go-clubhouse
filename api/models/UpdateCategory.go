package models

type UpdateCategory struct {
	// Archived a true/false boolean indicating if the Category has been archived.
	Archived bool `json:"archived,omitempty"`
	// Color the hex color to be displayed with the Category (for example, "#ff0000").
	Color *string `json:"color,omitempty"`
	// Name the new name of the Category.
	Name string `json:"name,omitempty"`
}

func (m *UpdateCategory) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateCategory) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
