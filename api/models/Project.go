package models

import (
	"encoding/json"
	"time"
)

// projects typically map to teams (such as Frontend, Backend, Mobile, Devops, etc) but can represent any open-ended product, component, or initiative.
type Project struct {
	// Abbreviation the Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation *string `json:"abbreviation,omitempty"`
	// AppURL the Clubhouse application url for the Project.
	AppURL string `json:"app_url,omitempty"`
	// Archived true/false boolean indicating whether the Project is in an Archived state.
	Archived bool `json:"archived,omitempty"`
	// Color the color associated with the Project in the Clubhouse member interface.
	Color *string `json:"color,omitempty"`
	// CreatedAt the time/date that the Project was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// DaysToThermometer the number of days before the thermometer appears in the Story summary.
	DaysToThermometer int64 `json:"days_to_thermometer,omitempty"`
	// Description the description of the Project.
	Description *string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id,omitempty"`
	// FollowerIDs an array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	// ID the unique ID of the Project.
	ID int64 `json:"id,omitempty"`
	// IterationLength the number of weeks per iteration in this Project.
	IterationLength int64 `json:"iteration_length,omitempty"`
	// Name the name of the Project
	Name string `json:"name,omitempty"`
	// ShowThermometer configuration to enable or disable thermometers in the Story summary.
	ShowThermometer bool `json:"show_thermometer,omitempty"`
	// StartTime the date at which the Project was started.
	StartTime time.Time    `json:"start_time,omitempty"`
	Stats     ProjectStats `json:"stats,omitempty"`
	// TeamID the ID of the team the project belongs to.
	TeamID int64 `json:"team_id,omitempty"`
	// UpdatedAt the time/date that the Project was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (m *Project) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Project) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
