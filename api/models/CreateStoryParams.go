package models

import (
	"encoding/json"
	"time"
)

type CreateStoryParams struct {
	// Controls the story's archived state.
	Archived bool `json:"archived"`
	// An array of comments to add to the story.
	Comments []CreateStoryCommentParams `json:"comments"`
	// A manual override for the time/date the Story was completed.
	CompletedAtOverride time.Time `json:"completed_at_override"`
	// The time/date the Story was created.
	CreatedAt time.Time `json:"created_at"`
	// The due date of the story.
	Deadline *time.Time `json:"deadline"`
	// The description of the story.
	Description string `json:"description"`
	// The ID of the epic the story belongs to.
	EpicId *int64 `json:"epic_id"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate"`
	// This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// An array of External Tickets associated with this story. These External Tickets must have unquie external id. Duplicated External Tickets will be removed.
	ExternalTickets []CreateExternalTicketParams `json:"external_tickets"`
	// An array of IDs of files attached to the story.
	FileIds []int64 `json:"file_ids"`
	// An array of UUIDs of the followers of this story.
	FollowerIds []string `json:"follower_ids"`
	// The id of the group to associate with this story.
	GroupId *string `json:"group_id"`
	// The ID of the iteration the story belongs to.
	IterationId *int64 `json:"iteration_id"`
	// An array of labels attached to the story.
	Labels []CreateLabelParams `json:"labels"`
	// An array of IDs of linked files attached to the story.
	LinkedFileIds []int64 `json:"linked_file_ids"`
	// The name of the story.
	Name string `json:"name"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids"`
	// The ID of the project the story belongs to.
	ProjectId int64 `json:"project_id"`
	// The ID of the member that requested the story.
	RequestedById string `json:"requested_by_id"`
	// A manual override for the time/date the Story was started.
	StartedAtOverride time.Time `json:"started_at_override"`
	// An array of story links attached to the story.
	StoryLinks []CreateStoryLinkParams `json:"story_links"`
	// The type of story (feature, bug, chore).
	StoryType      string                       `json:"story_type"`
	SupportTickets []CreateExternalTicketParams `json:"support_tickets"`
	// An array of tasks connected to the story.
	Tasks []CreateTaskParams `json:"tasks"`
	// The time/date the Story was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The ID of the workflow state the story will be in.
	WorkflowStateId int64 `json:"workflow_state_id"`
}

func (m *CreateStoryParams) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateStoryParams) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
