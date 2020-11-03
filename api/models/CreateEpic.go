package models

import (
	"encoding/json"
	"time"
)

type CreateEpic struct {
	// CompletedAtOverride a manual override for the time/date the Epic was completed.
	CompletedAtOverride time.Time `json:"completed_at_override,omitempty"`
	// CreatedAt defaults to the time/date it is created but can be set to reflect another date.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Deadline the Epic's deadline.
	Deadline *time.Time `json:"deadline,omitempty"`
	// Description the Epic's description.
	Description string `json:"description,omitempty"`
	// EpicStateID the ID of the Epic State.
	EpicStateID int64 `json:"epic_state_id,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`
	// FollowerIDs an array of UUIDs for any Members you want to add as Followers on this new Epic.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	// GroupID the ID of the group to associate with the epic.
	GroupID *string `json:"group_id,omitempty"`
	// Labels an array of Labels attached to the Epic.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// MilestoneID the ID of the Milestone this Epic is related to.
	MilestoneID *int64 `json:"milestone_id,omitempty"`
	// Name the Epic's name.
	Name string `json:"name,omitempty"`
	// OwnerIDs an array of UUIDs for any members you want to add as Owners on this new Epic.
	OwnerIDs []string `json:"owner_ids,omitempty"`
	// PlannedStartDate the Epic's planned start date.
	PlannedStartDate *time.Time `json:"planned_start_date,omitempty"`
	// RequestedByID the ID of the member that requested the epic.
	RequestedByID string `json:"requested_by_id,omitempty"`
	// StartedAtOverride a manual override for the time/date the Epic was started.
	StartedAtOverride time.Time `json:"started_at_override,omitempty"`
	// State `deprecated` The Epic's state (to do, in progress, or done); will be ignored when `epic_state_id` is set.
	State string `json:"state,omitempty"`
	// UpdatedAt defaults to the time/date it is created but can be set to reflect another date.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *CreateEpic) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateEpic) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
