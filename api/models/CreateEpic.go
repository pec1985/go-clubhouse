package models

import (
	"encoding/json"
	"time"
)

type CreateEpic struct {
	// This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// An array of UUIDs for any Members you want to add as Followers on this new Epic.
	FollowerIds []string `json:"follower_ids"`
	// The Epic's planned start date.
	PlannedStartDate *time.Time `json:"planned_start_date"`
	// `Deprecated` The Epic's state (to do, in progress, or done); will be ignored when `epic_state_id` is set.
	State string `json:"state"`
	// The Epic's deadline.
	Deadline *time.Time `json:"deadline"`
	// The Epic's description.
	Description string `json:"description"`
	// The ID of the Epic State.
	EpicStateId int64 `json:"epic_state_id"`
	// The ID of the group to associate with the epic.
	GroupId *string `json:"group_id"`
	// The ID of the Milestone this Epic is related to.
	MilestoneId *int64 `json:"milestone_id"`
	// A manual override for the time/date the Epic was completed.
	CompletedAtOverride time.Time `json:"completed_at_override"`
	// An array of Labels attached to the Epic.
	Labels []CreateLabelParams `json:"labels"`
	// An array of UUIDs for any members you want to add as Owners on this new Epic.
	OwnerIds []string `json:"owner_ids"`
	// The ID of the member that requested the epic.
	RequestedById string `json:"requested_by_id"`
	// Defaults to the time/date it is created but can be set to reflect another date.
	UpdatedAt time.Time `json:"updated_at"`
	// Defaults to the time/date it is created but can be set to reflect another date.
	CreatedAt time.Time `json:"created_at"`
	// The Epic's name.
	Name string `json:"name"`
	// A manual override for the time/date the Epic was started.
	StartedAtOverride time.Time `json:"started_at_override"`
}

func (m *CreateEpic) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateEpic) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
