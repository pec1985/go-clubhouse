package models

import (
	"encoding/json"
	"time"
)

// Projects typically map to teams (such as Frontend, Backend, Mobile, Devops, etc) but can represent any open-ended product, component, or initiative.
type Project struct {
	// Abbreviation The Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation *string `json:"abbreviation"`
	// AppURL The Clubhouse application url for the Project.
	AppURL string `json:"app_url"`
	// Archived True/false boolean indicating whether the Project is in an Archived state.
	Archived bool `json:"archived"`
	// Color The color associated with the Project in the Clubhouse member interface.
	Color *string `json:"color"`
	// CreatedAt The time/date that the Project was created.
	CreatedAt *time.Time `json:"created_at"`
	// DaysToThermometer The number of days before the thermometer appears in the Story summary.
	DaysToThermometer int64 `json:"days_to_thermometer"`
	// Description The description of the Project.
	Description *string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID This field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// FollowerIDs An array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids"`
	// ID The unique ID of the Project.
	ID int64 `json:"id"`
	// IterationLength The number of weeks per iteration in this Project.
	IterationLength int64 `json:"iteration_length"`
	// Name The name of the Project
	Name string `json:"name"`
	// ShowThermometer Configuration to enable or disable thermometers in the Story summary.
	ShowThermometer bool `json:"show_thermometer"`
	// StartTime The date at which the Project was started.
	StartTime time.Time    `json:"start_time"`
	Stats     ProjectStats `json:"stats"`
	// TeamID The ID of the team the project belongs to.
	TeamID int64 `json:"team_id"`
	// UpdatedAt The time/date that the Project was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

func (m *Project) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Project) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
