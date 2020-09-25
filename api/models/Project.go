package models

import (
	"encoding/json"
	"time"
)

type Project struct {
	// The number of weeks per iteration in this Project.
	IterationLength int64 `json:"iteration_length"`
	// The date at which the Project was started.
	StartTime time.Time `json:"start_time"`
	// The time/date that the Project was created.
	CreatedAt *time.Time `json:"created_at"`
	// The number of days before the thermometer appears in the Story summary.
	DaysToThermometer int64 `json:"days_to_thermometer"`
	// The unique ID of the Project.
	Id int64 `json:"id"`
	// This field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id"`
	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []string `json:"follower_ids"`
	// Configuration to enable or disable thermometers in the Story summary.
	ShowThermometer bool         `json:"show_thermometer"`
	Stats           ProjectStats `json:"stats"`
	// The ID of the team the project belongs to.
	TeamId int64 `json:"team_id"`
	// True/false boolean indicating whether the Project is in an Archived state.
	Archived bool `json:"archived"`
	// The color associated with the Project in the Clubhouse member interface.
	Color *string `json:"color"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The time/date that the Project was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// The Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation *string `json:"abbreviation"`
	// The description of the Project.
	Description *string `json:"description"`
	// The name of the Project
	Name string `json:"name"`
	// The Clubhouse application url for the Project.
	AppUrl string `json:"app_url"`
}

func (m *Project) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Project) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
