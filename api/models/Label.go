package models

import "time"

// Label a Label can be used to associate and filter Stories and Epics, and also create new Workspaces.
type Label struct {
	// AppURL the Clubhouse application url for the Label.
	AppURL string `json:"app_url,omitempty"`
	// Archived a true/false boolean indicating if the Label has been archived.
	Archived bool `json:"archived,omitempty"`
	// Color the hex color to be displayed with the Label (for example, "#ff0000").
	Color *string `json:"color,omitempty"`
	// CreatedAt the time/date that the Label was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Description the description of the Label.
	Description *string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id,omitempty"`
	// ID the unique ID of the Label.
	ID int64 `json:"id,omitempty"`
	// Name the name of the Label.
	Name  string     `json:"name,omitempty"`
	Stats LabelStats `json:"stats,omitempty"`
	// UpdatedAt the time/date that the Label was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (m *Label) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Label) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
