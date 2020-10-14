package models

import (
	"encoding/json"
	"time"
)

// StorySlim represents the same resource as a Story, but is more light-weight. For certain fields it provides ids rather than full resources (e.g., `comment_ids` and `file_ids`) and it also excludes certain aggregate values (e.g., `cycle_time`). The `description` field can be optionally included. Use the [Get Story](#Get-Story) endpoint to fetch the unabridged payload for a Story.
type StorySlim struct {
	// AppURL The Clubhouse application url for the Story.
	AppURL string `json:"app_url"`
	// Archived True if the story has been archived or not.
	Archived bool `json:"archived"`
	// Blocked A true/false boolean indicating if the Story is currently blocked.
	Blocked bool `json:"blocked"`
	// Blocker A true/false boolean indicating if the Story is currently a blocker of another story.
	Blocker bool `json:"blocker"`
	// CommentIDs An array of IDs of Comments attached to the story.
	CommentIDs []int64 `json:"comment_ids"`
	// Completed A true/false boolean indicating if the Story has been completed.
	Completed bool `json:"completed"`
	// CompletedAt The time/date the Story was completed.
	CompletedAt *time.Time `json:"completed_at"`
	// CompletedAtOverride A manual override for the time/date the Story was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// CreatedAt The time/date the Story was created.
	CreatedAt time.Time `json:"created_at"`
	// CycleTime The cycle time (in seconds) of this story when complete.
	CycleTime int64 `json:"cycle_time"`
	// Deadline The due date of the story.
	Deadline *time.Time `json:"deadline"`
	// Description The description of the Story.
	Description string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// EpicID The ID of the epic the story belongs to.
	EpicID *int64 `json:"epic_id"`
	// Estimate The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate"`
	// ExternalID This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// ExternalLinks An array of external link (strings) associated with a Story
	ExternalLinks []string `json:"external_links"`
	// ExternalTickets An array of External Tickets associated with a Story
	ExternalTickets []ExternalTicket `json:"external_tickets"`
	// FileIDs An array of IDs of Files attached to the story.
	FileIDs []int64 `json:"file_ids"`
	// FollowerIDs An array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids"`
	// GroupID The ID of the group associated with the story.
	GroupID *string `json:"group_id"`
	// GroupMentionIDs An array of Group IDs that have been mentioned in the Story description.
	GroupMentionIDs []string `json:"group_mention_ids"`
	// ID The unique ID of the Story.
	ID int64 `json:"id"`
	// IterationID The ID of the iteration the story belongs to.
	IterationID *int64 `json:"iteration_id"`
	// Labels An array of labels attached to the story.
	Labels []Label `json:"labels"`
	// LeadTime The lead time (in seconds) of this story when complete.
	LeadTime int64 `json:"lead_time"`
	// LinkedFileIDs An array of IDs of LinkedFiles attached to the story.
	LinkedFileIDs []int64 `json:"linked_file_ids"`
	// MemberMentionIDs An array of Member IDs that have been mentioned in the Story description.
	MemberMentionIDs []string `json:"member_mention_ids"`
	// MentionIDs Deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids"`
	// MovedAt The time/date the Story was last changed workflow-state.
	MovedAt *time.Time `json:"moved_at"`
	// Name The name of the story.
	Name string `json:"name"`
	// NumTasksCompleted The number of tasks on the story which are complete.
	NumTasksCompleted int64 `json:"num_tasks_completed"`
	// OwnerIDs An array of UUIDs of the owners of this story.
	OwnerIDs []string `json:"owner_ids"`
	// Position A number representing the position of the story in relation to every other story in the current project.
	Position int64 `json:"position"`
	// PreviousIterationIDs The IDs of the iteration the story belongs to.
	PreviousIterationIDs []int64 `json:"previous_iteration_ids"`
	// ProjectID The ID of the project the story belongs to.
	ProjectID int64 `json:"project_id"`
	// RequestedByID The ID of the Member that requested the story.
	RequestedByID string `json:"requested_by_id"`
	// Started A true/false boolean indicating if the Story has been started.
	Started bool `json:"started"`
	// StartedAt The time/date the Story was started.
	StartedAt *time.Time `json:"started_at"`
	// StartedAtOverride A manual override for the time/date the Story was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	Stats             StoryStats `json:"stats"`
	// StoryLinks An array of story links attached to the Story.
	StoryLinks []TypedStoryLink `json:"story_links"`
	// StoryType The type of story (feature, bug, chore).
	StoryType      string          `json:"story_type"`
	SupportTickets []SupportTicket `json:"support_tickets"`
	// TaskIDs An array of IDs of Tasks attached to the story.
	TaskIDs []int64 `json:"task_ids"`
	// UpdatedAt The time/date the Story was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// WorkflowStateID The ID of the workflow state the story is currently in.
	WorkflowStateID int64 `json:"workflow_state_id"`
}

func (m *StorySlim) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *StorySlim) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
