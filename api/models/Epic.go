package models

import (
	"encoding/json"
	"time"
)

// Epic An Epic is a collection of stories that together might make up a release, a milestone, or some other large initiative that your organization is working on.
type Epic struct {
	// AppURL The Clubhouse application url for the Epic.
	AppURL string `json:"app_url"`
	// Archived True/false boolean that indicates whether the Epic is archived or not.
	Archived bool `json:"archived"`
	// Comments A nested array of threaded comments.
	Comments []ThreadedComment `json:"comments"`
	// Completed A true/false boolean indicating if the Epic has been completed.
	Completed bool `json:"completed"`
	// CompletedAt The time/date the Epic was completed.
	CompletedAt *time.Time `json:"completed_at"`
	// CompletedAtOverride A manual override for the time/date the Epic was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// CreatedAt The time/date the Epic was created.
	CreatedAt *time.Time `json:"created_at"`
	// Deadline The Epic's deadline.
	Deadline *time.Time `json:"deadline"`
	// Description The Epic's description.
	Description string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// EpicStateID The ID of the Epic State.
	EpicStateID int64 `json:"epic_state_id"`
	// ExternalID This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID      *string          `json:"external_id"`
	ExternalTickets []ExternalTicket `json:"external_tickets"`
	// FollowerIDs An array of UUIDs for any Members you want to add as Followers on this Epic.
	FollowerIDs []string `json:"follower_ids"`
	GroupID     *string  `json:"group_id"`
	// GroupMentionIDs An array of Group IDs that have been mentioned in the Epic description.
	GroupMentionIDs []string `json:"group_mention_ids"`
	// ID The unique ID of the Epic.
	ID int64 `json:"id"`
	// Labels An array of Labels attached to the Epic.
	Labels []Label `json:"labels"`
	// MemberMentionIDs An array of Member IDs that have been mentioned in the Epic description.
	MemberMentionIDs []string `json:"member_mention_ids"`
	// MentionIDs Deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids"`
	// MilestoneID The ID of the Milestone this Epic is related to.
	MilestoneID *int64 `json:"milestone_id"`
	// Name The name of the Epic.
	Name string `json:"name"`
	// OwnerIDs An array of UUIDs for any members you want to add as Owners on this new Epic.
	OwnerIDs []string `json:"owner_ids"`
	// PlannedStartDate The Epic's planned start date.
	PlannedStartDate *time.Time `json:"planned_start_date"`
	// Position The Epic's relative position in the Epic workflow state.
	Position int64 `json:"position"`
	// ProjectIDs The IDs of Projects related to this Epic.
	ProjectIDs []int64 `json:"project_ids"`
	// RequestedByID The ID of the Member that requested the epic.
	RequestedByID string `json:"requested_by_id"`
	// Started A true/false boolean indicating if the Epic has been started.
	Started bool `json:"started"`
	// StartedAt The time/date the Epic was started.
	StartedAt *time.Time `json:"started_at"`
	// StartedAtOverride A manual override for the time/date the Epic was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	// State `Deprecated` The workflow state that the Epic is in.
	State          string          `json:"state"`
	Stats          EpicStats       `json:"stats"`
	SupportTickets []SupportTicket `json:"support_tickets"`
	// UpdatedAt The time/date the Epic was updated.
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
