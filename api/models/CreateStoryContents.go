package models

import (
	"encoding/json"
	"time"
)

// CreateStoryContents a map of story attributes this template populates.
type CreateStoryContents struct {
	// Deadline the due date of the story.
	Deadline *time.Time `json:"deadline"`
	// Description the description of the story.
	Description string `json:"description"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// EpicID the ID of the epic the to be populated.
	EpicID *int64 `json:"epic_id"`
	// Estimate the numeric point estimate to be populated.
	Estimate *int64 `json:"estimate"`
	// ExternalLinks an array of external links to be populated.
	ExternalLinks []string `json:"external_links"`
	// ExternalTickets an array of the external ticket IDs to be populated.
	ExternalTickets []CreateEntityTemplateExternalTicket `json:"external_tickets"`
	// FileIDs an array of the attached file IDs to be populated.
	FileIDs []int64 `json:"file_ids"`
	// Files an array of files attached to the story.
	Files []File `json:"files"`
	// FollowerIDs an array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids"`
	// IterationID the ID of the iteration the to be populated.
	IterationID *int64 `json:"iteration_id"`
	// Labels an array of labels to be populated by the template.
	Labels []CreateLabelParams `json:"labels"`
	// LinkedFileIDs an array of the linked file IDs to be populated.
	LinkedFileIDs []int64 `json:"linked_file_ids"`
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
	// Tasks an array of tasks to be populated by the template.
	Tasks []EntityTemplateTask `json:"tasks"`
	// WorkflowStateID the ID of the workflow state the story is currently in.
	WorkflowStateID int64 `json:"workflow_state_id"`
}

func (m *CreateStoryContents) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateStoryContents) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
