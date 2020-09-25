package models

import (
	"encoding/json"
	"time"
)

type CreateStoryContents struct {
	// The due date of the story.
	Deadline *time.Time `json:"deadline"`
	// The description of the story.
	Description string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The ID of the epic the to be populated.
	EpicId *int64 `json:"epic_id"`
	// The numeric point estimate to be populated.
	Estimate *int64 `json:"estimate"`
	// An array of the external ticket IDs to be populated.
	ExternalTickets []CreateEntityTemplateExternalTicket `json:"external_tickets"`
	// An array of the attached file IDs to be populated.
	FileIds []int64 `json:"file_ids"`
	// An array of files attached to the story.
	Files []File `json:"files"`
	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []string `json:"follower_ids"`
	// The ID of the iteration the to be populated.
	IterationId *int64 `json:"iteration_id"`
	// An array of labels to be populated by the template.
	Labels []CreateLabelParams `json:"labels"`
	// An array of the linked file IDs to be populated.
	LinkedFileIds []int64 `json:"linked_file_ids"`
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
	// An array of tasks to be populated by the template.
	Tasks []EntityTemplateTask `json:"tasks"`
	// The ID of the workflow state the story is currently in.
	WorkflowStateId int64 `json:"workflow_state_id"`
}

func (m *CreateStoryContents) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateStoryContents) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
