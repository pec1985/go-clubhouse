package models

import (
	"encoding/json"
	"time"
)

// CreateStoryContents A map of story attributes this template populates.
type CreateStoryContents struct {
	// Deadline The due date of the story.
	Deadline *time.Time `json:"deadline"`
	// Description The description of the story.
	Description string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// EpicID The ID of the epic the to be populated.
	EpicID *int64 `json:"epic_id"`
	// Estimate The numeric point estimate to be populated.
	Estimate *int64 `json:"estimate"`
	// ExternalLinks An array of external links to be populated.
	ExternalLinks []string `json:"external_links"`
	// ExternalTickets An array of the external ticket IDs to be populated.
	ExternalTickets []CreateEntityTemplateExternalTicket `json:"external_tickets"`
	// FileIDs An array of the attached file IDs to be populated.
	FileIDs []int64 `json:"file_ids"`
	// Files An array of files attached to the story.
	Files []File `json:"files"`
	// FollowerIDs An array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids"`
	// IterationID The ID of the iteration the to be populated.
	IterationID *int64 `json:"iteration_id"`
	// Labels An array of labels to be populated by the template.
	Labels []CreateLabelParams `json:"labels"`
	// LinkedFileIDs An array of the linked file IDs to be populated.
	LinkedFileIDs []int64 `json:"linked_file_ids"`
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
	// Tasks An array of tasks to be populated by the template.
	Tasks []EntityTemplateTask `json:"tasks"`
	// WorkflowStateID The ID of the workflow state the story is currently in.
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
