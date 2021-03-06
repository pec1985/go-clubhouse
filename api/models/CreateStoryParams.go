package models

import "time"

// CreateStoryParams used to create multiple stories in a single request.
type CreateStoryParams struct {
	// Archived controls the story's archived state.
	Archived bool `json:"archived,omitempty"`
	// Comments an array of comments to add to the story.
	Comments []CreateStoryCommentParams `json:"comments,omitempty"`
	// CompletedAtOverride a manual override for the time/date the Story was completed.
	CompletedAtOverride time.Time `json:"completed_at_override,omitempty"`
	// CreatedAt the time/date the Story was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Deadline the due date of the story.
	Deadline *time.Time `json:"deadline,omitempty"`
	// Description the description of the story.
	Description string `json:"description,omitempty"`
	// EpicID the ID of the epic the story belongs to.
	EpicID *int64 `json:"epic_id,omitempty"`
	// Estimate the numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`
	// ExternalLinks an array of External Links associated with this story.
	ExternalLinks []string `json:"external_links,omitempty"`
	// ExternalTickets an array of External Tickets associated with this story. These External Tickets must have unquie external id. Duplicated External Tickets will be removed.
	ExternalTickets []CreateExternalTicketParams `json:"external_tickets,omitempty"`
	// FileIDs an array of IDs of files attached to the story.
	FileIDs []int64 `json:"file_ids,omitempty"`
	// FollowerIDs an array of UUIDs of the followers of this story.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	// GroupID the id of the group to associate with this story.
	GroupID *string `json:"group_id,omitempty"`
	// IterationID the ID of the iteration the story belongs to.
	IterationID *int64 `json:"iteration_id,omitempty"`
	// Labels an array of labels attached to the story.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// LinkedFileIDs an array of IDs of linked files attached to the story.
	LinkedFileIDs []int64 `json:"linked_file_ids,omitempty"`
	// Name the name of the story.
	Name string `json:"name,omitempty"`
	// OwnerIDs an array of UUIDs of the owners of this story.
	OwnerIDs []string `json:"owner_ids,omitempty"`
	// ProjectID the ID of the project the story belongs to.
	ProjectID int64 `json:"project_id,omitempty"`
	// RequestedByID the ID of the member that requested the story.
	RequestedByID string `json:"requested_by_id,omitempty"`
	// StartedAtOverride a manual override for the time/date the Story was started.
	StartedAtOverride time.Time `json:"started_at_override,omitempty"`
	// StoryLinks an array of story links attached to the story.
	StoryLinks []CreateStoryLinkParams `json:"story_links,omitempty"`
	// StoryType the type of story (feature, bug, chore).
	StoryType      string                       `json:"story_type,omitempty"`
	SupportTickets []CreateExternalTicketParams `json:"support_tickets,omitempty"`
	// Tasks an array of tasks connected to the story.
	Tasks []CreateTaskParams `json:"tasks,omitempty"`
	// UpdatedAt the time/date the Story was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// WorkflowStateID the ID of the workflow state the story will be in.
	WorkflowStateID int64 `json:"workflow_state_id,omitempty"`
}

func (m *CreateStoryParams) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateStoryParams) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
