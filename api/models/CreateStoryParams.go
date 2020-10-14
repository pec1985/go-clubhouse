package models

import (
	"encoding/json"
	"time"
)

// CreateStoryParams Used to create multiple stories in a single request.
type CreateStoryParams struct {
	// Archived Controls the story's archived state.
	Archived bool `json:"archived"`
	// Comments An array of comments to add to the story.
	Comments []CreateStoryCommentParams `json:"comments"`
	// CompletedAtOverride A manual override for the time/date the Story was completed.
	CompletedAtOverride time.Time `json:"completed_at_override"`
	// CreatedAt The time/date the Story was created.
	CreatedAt time.Time `json:"created_at"`
	// Deadline The due date of the story.
	Deadline *time.Time `json:"deadline"`
	// Description The description of the story.
	Description string `json:"description"`
	// EpicID The ID of the epic the story belongs to.
	EpicID *int64 `json:"epic_id"`
	// Estimate The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate"`
	// ExternalID This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id"`
	// ExternalLinks An array of External Links associated with this story.
	ExternalLinks []string `json:"external_links"`
	// ExternalTickets An array of External Tickets associated with this story. These External Tickets must have unquie external id. Duplicated External Tickets will be removed.
	ExternalTickets []CreateExternalTicketParams `json:"external_tickets"`
	// FileIDs An array of IDs of files attached to the story.
	FileIDs []int64 `json:"file_ids"`
	// FollowerIDs An array of UUIDs of the followers of this story.
	FollowerIDs []string `json:"follower_ids"`
	// GroupID The id of the group to associate with this story.
	GroupID *string `json:"group_id"`
	// IterationID The ID of the iteration the story belongs to.
	IterationID *int64 `json:"iteration_id"`
	// Labels An array of labels attached to the story.
	Labels []CreateLabelParams `json:"labels"`
	// LinkedFileIDs An array of IDs of linked files attached to the story.
	LinkedFileIDs []int64 `json:"linked_file_ids"`
	// Name The name of the story.
	Name string `json:"name"`
	// OwnerIDs An array of UUIDs of the owners of this story.
	OwnerIDs []string `json:"owner_ids"`
	// ProjectID The ID of the project the story belongs to.
	ProjectID int64 `json:"project_id"`
	// RequestedByID The ID of the member that requested the story.
	RequestedByID string `json:"requested_by_id"`
	// StartedAtOverride A manual override for the time/date the Story was started.
	StartedAtOverride time.Time `json:"started_at_override"`
	// StoryLinks An array of story links attached to the story.
	StoryLinks []CreateStoryLinkParams `json:"story_links"`
	// StoryType The type of story (feature, bug, chore).
	StoryType      string                       `json:"story_type"`
	SupportTickets []CreateExternalTicketParams `json:"support_tickets"`
	// Tasks An array of tasks connected to the story.
	Tasks []CreateTaskParams `json:"tasks"`
	// UpdatedAt The time/date the Story was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// WorkflowStateID The ID of the workflow state the story will be in.
	WorkflowStateID int64 `json:"workflow_state_id"`
}

func (m *CreateStoryParams) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateStoryParams) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
