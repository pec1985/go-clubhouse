package models

import "time"

// storyslim represents the same resource as a Story, but is more light-weight. For certain fields it provides ids rather than full resources (e.g., `comment_ids` and `file_ids`) and it also excludes certain aggregate values (e.g., `cycle_time`). The `description` field can be optionally included. Use the [Get Story](#Get-Story) endpoint to fetch the unabridged payload for a Story.
type StorySlim struct {
	// AppURL the Clubhouse application url for the Story.
	AppURL string `json:"app_url,omitempty"`
	// Archived true if the story has been archived or not.
	Archived bool `json:"archived,omitempty"`
	// Blocked a true/false boolean indicating if the Story is currently blocked.
	Blocked bool `json:"blocked,omitempty"`
	// Blocker a true/false boolean indicating if the Story is currently a blocker of another story.
	Blocker bool `json:"blocker,omitempty"`
	// CommentIDs an array of IDs of Comments attached to the story.
	CommentIDs []int64 `json:"comment_ids,omitempty"`
	// Completed a true/false boolean indicating if the Story has been completed.
	Completed bool `json:"completed,omitempty"`
	// CompletedAt the time/date the Story was completed.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CompletedAtOverride a manual override for the time/date the Story was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override,omitempty"`
	// CreatedAt the time/date the Story was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CycleTime the cycle time (in seconds) of this story when complete.
	CycleTime int64 `json:"cycle_time,omitempty"`
	// Deadline the due date of the story.
	Deadline *time.Time `json:"deadline,omitempty"`
	// Description the description of the Story.
	Description string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// EpicID the ID of the epic the story belongs to.
	EpicID *int64 `json:"epic_id,omitempty"`
	// Estimate the numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id,omitempty"`
	// ExternalLinks an array of external link (strings) associated with a Story
	ExternalLinks []string `json:"external_links,omitempty"`
	// ExternalTickets an array of External Tickets associated with a Story
	ExternalTickets []ExternalTicket `json:"external_tickets,omitempty"`
	// FileIDs an array of IDs of Files attached to the story.
	FileIDs []int64 `json:"file_ids,omitempty"`
	// FollowerIDs an array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	// GroupID the ID of the group associated with the story.
	GroupID *string `json:"group_id,omitempty"`
	// GroupMentionIDs an array of Group IDs that have been mentioned in the Story description.
	GroupMentionIDs []string `json:"group_mention_ids,omitempty"`
	// ID the unique ID of the Story.
	ID int64 `json:"id,omitempty"`
	// IterationID the ID of the iteration the story belongs to.
	IterationID *int64 `json:"iteration_id,omitempty"`
	// Labels an array of labels attached to the story.
	Labels []Label `json:"labels,omitempty"`
	// LeadTime the lead time (in seconds) of this story when complete.
	LeadTime int64 `json:"lead_time,omitempty"`
	// LinkedFileIDs an array of IDs of LinkedFiles attached to the story.
	LinkedFileIDs []int64 `json:"linked_file_ids,omitempty"`
	// MemberMentionIDs an array of Member IDs that have been mentioned in the Story description.
	MemberMentionIDs []string `json:"member_mention_ids,omitempty"`
	// MentionIDs deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids,omitempty"`
	// MovedAt the time/date the Story was last changed workflow-state.
	MovedAt *time.Time `json:"moved_at,omitempty"`
	// Name the name of the story.
	Name string `json:"name,omitempty"`
	// NumTasksCompleted the number of tasks on the story which are complete.
	NumTasksCompleted int64 `json:"num_tasks_completed,omitempty"`
	// OwnerIDs an array of UUIDs of the owners of this story.
	OwnerIDs []string `json:"owner_ids,omitempty"`
	// Position a number representing the position of the story in relation to every other story in the current project.
	Position int64 `json:"position,omitempty"`
	// PreviousIterationIDs the IDs of the iteration the story belongs to.
	PreviousIterationIDs []int64 `json:"previous_iteration_ids,omitempty"`
	// ProjectID the ID of the project the story belongs to.
	ProjectID int64 `json:"project_id,omitempty"`
	// RequestedByID the ID of the Member that requested the story.
	RequestedByID string `json:"requested_by_id,omitempty"`
	// Started a true/false boolean indicating if the Story has been started.
	Started bool `json:"started,omitempty"`
	// StartedAt the time/date the Story was started.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// StartedAtOverride a manual override for the time/date the Story was started.
	StartedAtOverride *time.Time `json:"started_at_override,omitempty"`
	Stats             StoryStats `json:"stats,omitempty"`
	// StoryLinks an array of story links attached to the Story.
	StoryLinks []TypedStoryLink `json:"story_links,omitempty"`
	// StoryType the type of story (feature, bug, chore).
	StoryType      string          `json:"story_type,omitempty"`
	SupportTickets []SupportTicket `json:"support_tickets,omitempty"`
	// TaskIDs an array of IDs of Tasks attached to the story.
	TaskIDs []int64 `json:"task_ids,omitempty"`
	// UpdatedAt the time/date the Story was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// WorkflowStateID the ID of the workflow state the story is currently in.
	WorkflowStateID int64 `json:"workflow_state_id,omitempty"`
}

func (m *StorySlim) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *StorySlim) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
