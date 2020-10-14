package models

import (
	"encoding/json"
	"time"
)

// StoryContents A container entity for the attributes this template should populate.
type StoryContents struct {
	// Deadline The due date of the story.
	Deadline time.Time `json:"deadline"`
	// Description The description of the story.
	Description string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// EpicID The ID of the epic the story belongs to.
	EpicID int64 `json:"epic_id"`
	// Estimate The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate int64 `json:"estimate"`
	// ExternalLinks An array of external links connected to the story.
	ExternalLinks []string `json:"external_links"`
	// ExternalTickets An array of external tickets connected to the story.
	ExternalTickets []ExternalTicket `json:"external_tickets"`
	// Files An array of files attached to the story.
	Files []File `json:"files"`
	// FollowerIDs An array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids"`
	// IterationID The ID of the iteration the story belongs to.
	IterationID int64 `json:"iteration_id"`
	// Labels An array of labels attached to the story.
	Labels []Label `json:"labels"`
	// LinkedFiles An array of linked files attached to the story.
	LinkedFiles []LinkedFile `json:"linked_files"`
	// Name The name of the story.
	Name string `json:"name"`
	// OwnerIDs An array of UUIDs of the owners of this story.
	OwnerIDs []string `json:"owner_ids"`
	// ProjectID The ID of the project the story belongs to.
	ProjectID int64 `json:"project_id"`
	// StoryType The type of story (feature, bug, chore).
	StoryType string `json:"story_type"`
	// Tasks An array of tasks connected to the story.
	Tasks []StoryContentsTask `json:"tasks"`
	// WorkflowStateID The ID of the workflow state the story is currently in.
	WorkflowStateID int64 `json:"workflow_state_id"`
}

func (m *StoryContents) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *StoryContents) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
