package models

import (
	"encoding/json"
	"time"
)

type UpdateMilestone struct {
	// AfterID the ID of the Milestone we want to move this Milestone after.
	AfterID int64 `json:"after_id"`
	// BeforeID the ID of the Milestone we want to move this Milestone before.
	BeforeID int64 `json:"before_id"`
	// Categories an array of IDs of Categories attached to the Milestone.
	Categories []CreateCategoryParams `json:"categories"`
	// CompletedAtOverride a manual override for the time/date the Milestone was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// Description the Milestone's description.
	Description string `json:"description"`
	// Name the name of the Milestone.
	Name string `json:"name"`
	// StartedAtOverride a manual override for the time/date the Milestone was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	// State the workflow state that the Milestone is in.
	State string `json:"state"`
}

func (m *UpdateMilestone) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateMilestone) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
