package models

import (
	"encoding/json"
	"time"
)

// Label A Label can be used to associate and filter Stories and Epics, and also create new Workspaces.
type Label struct {
	// AppURL The Clubhouse application url for the Label.
	AppURL string `json:"app_url"`
	// Archived A true/false boolean indicating if the Label has been archived.
	Archived bool `json:"archived"`
	// Color The hex color to be displayed with the Label (for example, "#ff0000").
	Color *string `json:"color"`
	// CreatedAt The time/date that the Label was created.
	CreatedAt *time.Time `json:"created_at"`
	// Description The description of the Label.
	Description *string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID This field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// ID The unique ID of the Label.
	ID int64 `json:"id"`
	// Name The name of the Label.
	Name  string     `json:"name"`
	Stats LabelStats `json:"stats"`
	// UpdatedAt The time/date that the Label was updated.
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
