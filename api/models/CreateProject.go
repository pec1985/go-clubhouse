package models

import (
	"encoding/json"
	"time"
)

type CreateProject struct {
	// The Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation string `json:"abbreviation"`
	// The color you wish to use for the Project in the system.
	Color string `json:"color"`
	// Defaults to the time/date it is created but can be set to reflect another date.
	CreatedAt time.Time `json:"created_at"`
	// The Project description.
	Description string `json:"description"`
	// This field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// An array of UUIDs for any members you want to add as Owners on this new Epic.
	FollowerIds []string `json:"follower_ids"`
	// The number of weeks per iteration in this Project.
	IterationLength int64 `json:"iteration_length"`
	// The name of the Project.
	Name string `json:"name"`
	// The date at which the Project was started.
	StartTime time.Time `json:"start_time"`
	// The ID of the team the project belongs to.
	TeamId int64 `json:"team_id"`
	// Defaults to the time/date it is created but can be set to reflect another date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *CreateProject) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateProject) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
