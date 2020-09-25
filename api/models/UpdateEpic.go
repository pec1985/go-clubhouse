package models

import (
	"encoding/json"
	"time"
)

type UpdateEpic struct {
	// The ID of the group to associate with the epic.
	GroupId *string `json:"group_id"`
	// An array of Labels attached to the Epic.
	Labels []CreateLabelParams `json:"labels"`
	// The Epic's name.
	Name string `json:"name"`
	// The ID of the member that requested the epic.
	RequestedById string `json:"requested_by_id"`
	// The ID of the Epic we want to move this Epic before.
	BeforeId int64 `json:"before_id"`
	// The Epic's description.
	Description string `json:"description"`
	// The ID of the Milestone this Epic is related to.
	MilestoneId *int64 `json:"milestone_id"`
	// The Epic's planned start date.
	PlannedStartDate *time.Time `json:"planned_start_date"`
	// The ID of the Epic we want to move this Epic after.
	AfterId int64 `json:"after_id"`
	// The Epic's deadline.
	Deadline *time.Time `json:"deadline"`
	// `Deprecated` The Epic's state (to do, in progress, or done); will be ignored when `epic_state_id` is set.
	State string `json:"state"`
	// A manual override for the time/date the Epic was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// A manual override for the time/date the Epic was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	// An array of UUIDs for any Members you want to add as Followers on this Epic.
	FollowerIds []string `json:"follower_ids"`
	// An array of UUIDs for any members you want to add as Owners on this Epic.
	OwnerIds []string `json:"owner_ids"`
	// A true/false boolean indicating whether the Epic is in archived state.
	Archived bool `json:"archived"`
	// The ID of the Epic State.
	EpicStateId int64 `json:"epic_state_id"`
}

func (m *UpdateEpic) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateEpic) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
