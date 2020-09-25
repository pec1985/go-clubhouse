package models

import (
	"encoding/json"
	"time"
)

type Milestone struct {
	// The Clubhouse application url for the Milestone.
	AppUrl string `json:"app_url"`
	// An array of Categories attached to the Milestone.
	Categories []Category `json:"categories"`
	// A true/false boolean indicating if the Milestone has been completed.
	Completed bool `json:"completed"`
	// The time/date the Milestone was completed.
	CompletedAt *time.Time `json:"completed_at"`
	// A manual override for the time/date the Milestone was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// The time/date the Milestone was created.
	CreatedAt time.Time `json:"created_at"`
	// The Milestone's description.
	Description string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique ID of the Milestone.
	Id int64 `json:"id"`
	// The name of the Milestone.
	Name string `json:"name"`
	// A number representing the position of the Milestone in relation to every other Milestone within the Organization.
	Position int64 `json:"position"`
	// A true/false boolean indicating if the Milestone has been started.
	Started bool `json:"started"`
	// The time/date the Milestone was started.
	StartedAt *time.Time `json:"started_at"`
	// A manual override for the time/date the Milestone was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	// The workflow state that the Milestone is in.
	State string         `json:"state"`
	Stats MilestoneStats `json:"stats"`
	// The time/date the Milestone was updated.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Milestone) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Milestone) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
