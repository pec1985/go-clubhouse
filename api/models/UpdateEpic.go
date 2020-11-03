package models

import "time"

type UpdateEpic struct {
	// AfterID the ID of the Epic we want to move this Epic after.
	AfterID int64 `json:"after_id,omitempty"`
	// Archived a true/false boolean indicating whether the Epic is in archived state.
	Archived bool `json:"archived,omitempty"`
	// BeforeID the ID of the Epic we want to move this Epic before.
	BeforeID int64 `json:"before_id,omitempty"`
	// CompletedAtOverride a manual override for the time/date the Epic was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override,omitempty"`
	// Deadline the Epic's deadline.
	Deadline *time.Time `json:"deadline,omitempty"`
	// Description the Epic's description.
	Description string `json:"description,omitempty"`
	// EpicStateID the ID of the Epic State.
	EpicStateID int64 `json:"epic_state_id,omitempty"`
	// FollowerIDs an array of UUIDs for any Members you want to add as Followers on this Epic.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	// GroupID the ID of the group to associate with the epic.
	GroupID *string `json:"group_id,omitempty"`
	// Labels an array of Labels attached to the Epic.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// MilestoneID the ID of the Milestone this Epic is related to.
	MilestoneID *int64 `json:"milestone_id,omitempty"`
	// Name the Epic's name.
	Name string `json:"name,omitempty"`
	// OwnerIDs an array of UUIDs for any members you want to add as Owners on this Epic.
	OwnerIDs []string `json:"owner_ids,omitempty"`
	// PlannedStartDate the Epic's planned start date.
	PlannedStartDate *time.Time `json:"planned_start_date,omitempty"`
	// RequestedByID the ID of the member that requested the epic.
	RequestedByID string `json:"requested_by_id,omitempty"`
	// StartedAtOverride a manual override for the time/date the Epic was started.
	StartedAtOverride *time.Time `json:"started_at_override,omitempty"`
	// State `deprecated` The Epic's state (to do, in progress, or done); will be ignored when `epic_state_id` is set.
	State string `json:"state,omitempty"`
}

func (m *UpdateEpic) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateEpic) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
