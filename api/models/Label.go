package models

import (
	"encoding/json"
	"time"
)

type Label struct {
	// The Clubhouse application url for the Label.
	AppUrl string `json:"app_url"`
	// A true/false boolean indicating if the Label has been archived.
	Archived bool `json:"archived"`
	// The hex color to be displayed with the Label (for example, "#ff0000").
	Color *string `json:"color"`
	// The time/date that the Label was created.
	CreatedAt *time.Time `json:"created_at"`
	// The description of the Label.
	Description *string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// This field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id"`
	// The unique ID of the Label.
	Id int64 `json:"id"`
	// The name of the Label.
	Name  string     `json:"name"`
	Stats LabelStats `json:"stats"`
	// The time/date that the Label was updated.
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
