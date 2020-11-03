package models

import (
	"encoding/json"
	"time"
)

// epicslim represents the same resource as an Epic but is more light-weight, including all Epic fields except the comments array. The description string can be optionally included. Use the [Get Epic](#Get-Epic) endpoint to fetch the unabridged payload for an Epic.
type EpicSlim struct {
	// AppURL the Clubhouse application url for the Epic.
	AppURL string `json:"app_url,omitempty"`
	// Archived true/false boolean that indicates whether the Epic is archived or not.
	Archived bool `json:"archived,omitempty"`
	// Completed a true/false boolean indicating if the Epic has been completed.
	Completed bool `json:"completed,omitempty"`
	// CompletedAt the time/date the Epic was completed.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CompletedAtOverride a manual override for the time/date the Epic was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override,omitempty"`
	// CreatedAt the time/date the Epic was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Deadline the Epic's deadline.
	Deadline *time.Time `json:"deadline,omitempty"`
	// Description the Epic's description.
	Description string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// EpicStateID the ID of the Epic State.
	EpicStateID int64 `json:"epic_state_id,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID      *string          `json:"external_id,omitempty"`
	ExternalTickets []ExternalTicket `json:"external_tickets,omitempty"`
	// FollowerIDs an array of UUIDs for any Members you want to add as Followers on this Epic.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	GroupID     *string  `json:"group_id,omitempty"`
	// GroupMentionIDs an array of Group IDs that have been mentioned in the Epic description.
	GroupMentionIDs []string `json:"group_mention_ids,omitempty"`
	// ID the unique ID of the Epic.
	ID int64 `json:"id,omitempty"`
	// Labels an array of Labels attached to the Epic.
	Labels []Label `json:"labels,omitempty"`
	// MemberMentionIDs an array of Member IDs that have been mentioned in the Epic description.
	MemberMentionIDs []string `json:"member_mention_ids,omitempty"`
	// MentionIDs deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids,omitempty"`
	// MilestoneID the ID of the Milestone this Epic is related to.
	MilestoneID *int64 `json:"milestone_id,omitempty"`
	// Name the name of the Epic.
	Name string `json:"name,omitempty"`
	// OwnerIDs an array of UUIDs for any members you want to add as Owners on this new Epic.
	OwnerIDs []string `json:"owner_ids,omitempty"`
	// PlannedStartDate the Epic's planned start date.
	PlannedStartDate *time.Time `json:"planned_start_date,omitempty"`
	// Position the Epic's relative position in the Epic workflow state.
	Position int64 `json:"position,omitempty"`
	// ProjectIDs the IDs of Projects related to this Epic.
	ProjectIDs []int64 `json:"project_ids,omitempty"`
	// RequestedByID the ID of the Member that requested the epic.
	RequestedByID string `json:"requested_by_id,omitempty"`
	// Started a true/false boolean indicating if the Epic has been started.
	Started bool `json:"started,omitempty"`
	// StartedAt the time/date the Epic was started.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// StartedAtOverride a manual override for the time/date the Epic was started.
	StartedAtOverride *time.Time `json:"started_at_override,omitempty"`
	// State `deprecated` The workflow state that the Epic is in.
	State          string          `json:"state,omitempty"`
	Stats          EpicStats       `json:"stats,omitempty"`
	SupportTickets []SupportTicket `json:"support_tickets,omitempty"`
	// UpdatedAt the time/date the Epic was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (m *EpicSlim) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EpicSlim) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
