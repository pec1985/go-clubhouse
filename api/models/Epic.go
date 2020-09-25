package models

import (
	"encoding/json"
	"time"
)

type Epic struct {
	// The Clubhouse application url for the Epic.
	AppUrl string `json:"app_url"`
	// True/false boolean that indicates whether the Epic is archived or not.
	Archived bool `json:"archived"`
	// A nested array of threaded comments.
	Comments []ThreadedComment `json:"comments"`
	// A true/false boolean indicating if the Epic has been completed.
	Completed bool `json:"completed"`
	// The time/date the Epic was completed.
	CompletedAt *time.Time `json:"completed_at"`
	// A manual override for the time/date the Epic was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// The time/date the Epic was created.
	CreatedAt *time.Time `json:"created_at"`
	// The Epic's deadline.
	Deadline *time.Time `json:"deadline"`
	// The Epic's description.
	Description string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The ID of the Epic State.
	EpicStateId int64 `json:"epic_state_id"`
	// This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId      *string          `json:"external_id"`
	ExternalTickets []ExternalTicket `json:"external_tickets"`
	// An array of UUIDs for any Members you want to add as Followers on this Epic.
	FollowerIds []string `json:"follower_ids"`
	GroupId     *string  `json:"group_id"`
	// An array of Group IDs that have been mentioned in the Epic description.
	GroupMentionIds []string `json:"group_mention_ids"`
	// The unique ID of the Epic.
	Id int64 `json:"id"`
	// An array of Labels attached to the Epic.
	Labels []Label `json:"labels"`
	// An array of Member IDs that have been mentioned in the Epic description.
	MemberMentionIds []string `json:"member_mention_ids"`
	// Deprecated: use member_mention_ids.
	MentionIds []string `json:"mention_ids"`
	// The ID of the Milestone this Epic is related to.
	MilestoneId *int64 `json:"milestone_id"`
	// The name of the Epic.
	Name string `json:"name"`
	// An array of UUIDs for any members you want to add as Owners on this new Epic.
	OwnerIds []string `json:"owner_ids"`
	// The Epic's planned start date.
	PlannedStartDate *time.Time `json:"planned_start_date"`
	// The Epic's relative position in the Epic workflow state.
	Position int64 `json:"position"`
	// The IDs of Projects related to this Epic.
	ProjectIds []int64 `json:"project_ids"`
	// The ID of the Member that requested the epic.
	RequestedById string `json:"requested_by_id"`
	// A true/false boolean indicating if the Epic has been started.
	Started bool `json:"started"`
	// The time/date the Epic was started.
	StartedAt *time.Time `json:"started_at"`
	// A manual override for the time/date the Epic was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	// `Deprecated` The workflow state that the Epic is in.
	State          string          `json:"state"`
	Stats          EpicStats       `json:"stats"`
	SupportTickets []SupportTicket `json:"support_tickets"`
	// The time/date the Epic was updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

func (m *Epic) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Epic) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
