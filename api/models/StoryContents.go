package models

import (
	"encoding/json"
	"time"
)

// StoryContents a container entity for the attributes this template should populate.
type StoryContents struct {
	// Deadline the due date of the story.
	Deadline time.Time `json:"deadline"`
	// Description the description of the story.
	Description string `json:"description"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// EpicID the ID of the epic the story belongs to.
	EpicID int64 `json:"epic_id"`
	// Estimate the numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate int64 `json:"estimate"`
	// ExternalLinks an array of external links connected to the story.
	ExternalLinks []string `json:"external_links"`
	// ExternalTickets an array of external tickets connected to the story.
	ExternalTickets []ExternalTicket `json:"external_tickets"`
	// Files an array of files attached to the story.
	Files []File `json:"files"`
	// FollowerIDs an array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids"`
	// IterationID the ID of the iteration the story belongs to.
	IterationID int64 `json:"iteration_id"`
	// Labels an array of labels attached to the story.
	Labels []Label `json:"labels"`
	// LinkedFiles an array of linked files attached to the story.
	LinkedFiles []LinkedFile `json:"linked_files"`
	// Name the name of the story.
	Name string `json:"name"`
	// OwnerIDs an array of UUIDs of the owners of this story.
	OwnerIDs []string `json:"owner_ids"`
	// ProjectID the ID of the project the story belongs to.
	ProjectID int64 `json:"project_id"`
	// StoryType the type of story (feature, bug, chore).
	StoryType string `json:"story_type"`
	// Tasks an array of tasks connected to the story.
	Tasks []StoryContentsTask `json:"tasks"`
	// WorkflowStateID the ID of the workflow state the story is currently in.
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
