package models

import (
	"encoding/json"
	"time"
)

// Label a Label can be used to associate and filter Stories and Epics, and also create new Workspaces.
type Label struct {
	// AppURL the Clubhouse application url for the Label.
	AppURL string `json:"app_url"`
	// Archived a true/false boolean indicating if the Label has been archived.
	Archived bool `json:"archived"`
	// Color the hex color to be displayed with the Label (for example, "#ff0000").
	Color *string `json:"color"`
	// CreatedAt the time/date that the Label was created.
	CreatedAt *time.Time `json:"created_at"`
	// Description the description of the Label.
	Description *string `json:"description"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID this field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// ID the unique ID of the Label.
	ID int64 `json:"id"`
	// Name the name of the Label.
	Name  string     `json:"name"`
	Stats LabelStats `json:"stats"`
	// UpdatedAt the time/date that the Label was updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

func (m *Label) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Label) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
