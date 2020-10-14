package models

import (
	"encoding/json"
	"time"
)

type CreateEpic struct {
	// CompletedAtOverride A manual override for the time/date the Epic was completed.
	CompletedAtOverride time.Time `json:"completed_at_override"`
	// CreatedAt Defaults to the time/date it is created but can be set to reflect another date.
	CreatedAt time.Time `json:"created_at"`
	// Deadline The Epic's deadline.
	Deadline *time.Time `json:"deadline"`
	// Description The Epic's description.
	Description string `json:"description"`
	// EpicStateID The ID of the Epic State.
	EpicStateID int64 `json:"epic_state_id"`
	// ExternalID This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id"`
	// FollowerIDs An array of UUIDs for any Members you want to add as Followers on this new Epic.
	FollowerIDs []string `json:"follower_ids"`
	// GroupID The ID of the group to associate with the epic.
	GroupID *string `json:"group_id"`
	// Labels An array of Labels attached to the Epic.
	Labels []CreateLabelParams `json:"labels"`
	// MilestoneID The ID of the Milestone this Epic is related to.
	MilestoneID *int64 `json:"milestone_id"`
	// Name The Epic's name.
	Name string `json:"name"`
	// OwnerIDs An array of UUIDs for any members you want to add as Owners on this new Epic.
	OwnerIDs []string `json:"owner_ids"`
	// PlannedStartDate The Epic's planned start date.
	PlannedStartDate *time.Time `json:"planned_start_date"`
	// RequestedByID The ID of the member that requested the epic.
	RequestedByID string `json:"requested_by_id"`
	// StartedAtOverride A manual override for the time/date the Epic was started.
	StartedAtOverride time.Time `json:"started_at_override"`
	// State `Deprecated` The Epic's state (to do, in progress, or done); will be ignored when `epic_state_id` is set.
	State string `json:"state"`
	// UpdatedAt Defaults to the time/date it is created but can be set to reflect another date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *CreateEpic) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateEpic) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
