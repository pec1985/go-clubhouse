package models

import "time"

type UpdateMilestone struct {
	// AfterID the ID of the Milestone we want to move this Milestone after.
	AfterID int64 `json:"after_id,omitempty"`
	// BeforeID the ID of the Milestone we want to move this Milestone before.
	BeforeID int64 `json:"before_id,omitempty"`
	// Categories an array of IDs of Categories attached to the Milestone.
	Categories []CreateCategoryParams `json:"categories,omitempty"`
	// CompletedAtOverride a manual override for the time/date the Milestone was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override,omitempty"`
	// Description the Milestone's description.
	Description string `json:"description,omitempty"`
	// Name the name of the Milestone.
	Name string `json:"name,omitempty"`
	// StartedAtOverride a manual override for the time/date the Milestone was started.
	StartedAtOverride *time.Time `json:"started_at_override,omitempty"`
	// State the workflow state that the Milestone is in.
	State string `json:"state,omitempty"`
}

func (m *UpdateMilestone) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateMilestone) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
