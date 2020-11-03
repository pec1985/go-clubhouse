package models

// CreateLabelParams request parameters for creating a Label on a Clubhouse story.
type CreateLabelParams struct {
	// Color the hex color to be displayed with the Label (for example, "#ff0000").
	Color string `json:"color,omitempty"`
	// Description the description of the new Label.
	Description string `json:"description,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`
	// Name the name of the new Label.
	Name string `json:"name,omitempty"`
}

func (m *CreateLabelParams) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateLabelParams) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
