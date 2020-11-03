package models

import "time"

// StoryContents a container entity for the attributes this template should populate.
type StoryContents struct {
	// Deadline the due date of the story.
	Deadline time.Time `json:"deadline,omitempty"`
	// Description the description of the story.
	Description string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// EpicID the ID of the epic the story belongs to.
	EpicID int64 `json:"epic_id,omitempty"`
	// Estimate the numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate int64 `json:"estimate,omitempty"`
	// ExternalLinks an array of external links connected to the story.
	ExternalLinks []string `json:"external_links,omitempty"`
	// ExternalTickets an array of external tickets connected to the story.
	ExternalTickets []ExternalTicket `json:"external_tickets,omitempty"`
	// Files an array of files attached to the story.
	Files []File `json:"files,omitempty"`
	// FollowerIDs an array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	// IterationID the ID of the iteration the story belongs to.
	IterationID int64 `json:"iteration_id,omitempty"`
	// Labels an array of labels attached to the story.
	Labels []Label `json:"labels,omitempty"`
	// LinkedFiles an array of linked files attached to the story.
	LinkedFiles []LinkedFile `json:"linked_files,omitempty"`
	// Name the name of the story.
	Name string `json:"name,omitempty"`
	// OwnerIDs an array of UUIDs of the owners of this story.
	OwnerIDs []string `json:"owner_ids,omitempty"`
	// ProjectID the ID of the project the story belongs to.
	ProjectID int64 `json:"project_id,omitempty"`
	// StoryType the type of story (feature, bug, chore).
	StoryType string `json:"story_type,omitempty"`
	// Tasks an array of tasks connected to the story.
	Tasks []StoryContentsTask `json:"tasks,omitempty"`
	// WorkflowStateID the ID of the workflow state the story is currently in.
	WorkflowStateID int64 `json:"workflow_state_id,omitempty"`
}

func (m *StoryContents) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *StoryContents) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
