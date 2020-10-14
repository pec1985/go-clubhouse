package models

import (
	"encoding/json"
	"time"
)

// Milestone A Milestone is a collection of Epics that represent a release or some other large initiative that your organization is working on.
type Milestone struct {
	// AppURL The Clubhouse application url for the Milestone.
	AppURL string `json:"app_url"`
	// Categories An array of Categories attached to the Milestone.
	Categories []Category `json:"categories"`
	// Completed A true/false boolean indicating if the Milestone has been completed.
	Completed bool `json:"completed"`
	// CompletedAt The time/date the Milestone was completed.
	CompletedAt *time.Time `json:"completed_at"`
	// CompletedAtOverride A manual override for the time/date the Milestone was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// CreatedAt The time/date the Milestone was created.
	CreatedAt time.Time `json:"created_at"`
	// Description The Milestone's description.
	Description string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ID The unique ID of the Milestone.
	ID int64 `json:"id"`
	// Name The name of the Milestone.
	Name string `json:"name"`
	// Position A number representing the position of the Milestone in relation to every other Milestone within the Organization.
	Position int64 `json:"position"`
	// Started A true/false boolean indicating if the Milestone has been started.
	Started bool `json:"started"`
	// StartedAt The time/date the Milestone was started.
	StartedAt *time.Time `json:"started_at"`
	// StartedAtOverride A manual override for the time/date the Milestone was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	// State The workflow state that the Milestone is in.
	State string         `json:"state"`
	Stats MilestoneStats `json:"stats"`
	// UpdatedAt The time/date the Milestone was updated.
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
