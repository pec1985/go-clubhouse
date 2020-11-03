package models

import "time"

type CreateMilestone struct {
	// Categories an array of IDs of Categories attached to the Milestone.
	Categories []CreateCategoryParams `json:"categories,omitempty"`
	// CompletedAtOverride a manual override for the time/date the Milestone was completed.
	CompletedAtOverride time.Time `json:"completed_at_override,omitempty"`
	// Description the Milestone's description.
	Description string `json:"description,omitempty"`
	// Name the name of the Milestone.
	Name string `json:"name,omitempty"`
	// StartedAtOverride a manual override for the time/date the Milestone was started.
	StartedAtOverride time.Time `json:"started_at_override,omitempty"`
	// State the workflow state that the Milestone is in.
	State string `json:"state,omitempty"`
}

func (m *CreateMilestone) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateMilestone) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
