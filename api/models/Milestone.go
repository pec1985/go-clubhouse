package models

import "time"

// Milestone a Milestone is a collection of Epics that represent a release or some other large initiative that your organization is working on.
type Milestone struct {
	// AppURL the Clubhouse application url for the Milestone.
	AppURL string `json:"app_url,omitempty"`
	// Categories an array of Categories attached to the Milestone.
	Categories []Category `json:"categories,omitempty"`
	// Completed a true/false boolean indicating if the Milestone has been completed.
	Completed bool `json:"completed,omitempty"`
	// CompletedAt the time/date the Milestone was completed.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CompletedAtOverride a manual override for the time/date the Milestone was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override,omitempty"`
	// CreatedAt the time/date the Milestone was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Description the Milestone's description.
	Description string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ID the unique ID of the Milestone.
	ID int64 `json:"id,omitempty"`
	// Name the name of the Milestone.
	Name string `json:"name,omitempty"`
	// Position a number representing the position of the Milestone in relation to every other Milestone within the Organization.
	Position int64 `json:"position,omitempty"`
	// Started a true/false boolean indicating if the Milestone has been started.
	Started bool `json:"started,omitempty"`
	// StartedAt the time/date the Milestone was started.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// StartedAtOverride a manual override for the time/date the Milestone was started.
	StartedAtOverride *time.Time `json:"started_at_override,omitempty"`
	// State the workflow state that the Milestone is in.
	State string         `json:"state,omitempty"`
	Stats MilestoneStats `json:"stats,omitempty"`
	// UpdatedAt the time/date the Milestone was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *Milestone) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Milestone) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
