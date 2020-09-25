package models

import (
	"encoding/json"
	"time"
)

type StoryContents struct {
	// The due date of the story.
	Deadline time.Time `json:"deadline"`
	// The description of the story.
	Description string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The ID of the epic the story belongs to.
	EpicId int64 `json:"epic_id"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate int64 `json:"estimate"`
	// An array of external tickets connected to the story.
	ExternalTickets []ExternalTicket `json:"external_tickets"`
	// An array of files attached to the story.
	Files []File `json:"files"`
	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []string `json:"follower_ids"`
	// The ID of the iteration the story belongs to.
	IterationId int64 `json:"iteration_id"`
	// An array of labels attached to the story.
	Labels []Label `json:"labels"`
	// An array of linked files attached to the story.
	LinkedFiles []LinkedFile `json:"linked_files"`
	// The name of the story.
	Name string `json:"name"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids"`
	// The ID of the project the story belongs to.
	ProjectId int64 `json:"project_id"`
	// The type of story (feature, bug, chore).
	StoryType string `json:"story_type"`
	// An array of tasks connected to the story.
	Tasks []StoryContentsTask `json:"tasks"`
	// The ID of the workflow state the story is currently in.
	WorkflowStateId int64 `json:"workflow_state_id"`
}

func (m *StoryContents) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *StoryContents) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
