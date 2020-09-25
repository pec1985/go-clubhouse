package models

import (
	"encoding/json"
	"time"
)

type UpdateMilestone struct {
	// The ID of the Milestone we want to move this Milestone after.
	AfterId int64 `json:"after_id"`
	// The ID of the Milestone we want to move this Milestone before.
	BeforeId int64 `json:"before_id"`
	// An array of IDs of Categories attached to the Milestone.
	Categories []CreateCategoryParams `json:"categories"`
	// A manual override for the time/date the Milestone was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// The Milestone's description.
	Description string `json:"description"`
	// The name of the Milestone.
	Name string `json:"name"`
	// A manual override for the time/date the Milestone was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	// The workflow state that the Milestone is in.
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
