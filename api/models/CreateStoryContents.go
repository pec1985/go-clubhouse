package models

import "time"

// CreateStoryContents a map of story attributes this template populates.
type CreateStoryContents struct {
	// Deadline the due date of the story.
	Deadline *time.Time `json:"deadline,omitempty"`
	// Description the description of the story.
	Description string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// EpicID the ID of the epic the to be populated.
	EpicID *int64 `json:"epic_id,omitempty"`
	// Estimate the numeric point estimate to be populated.
	Estimate *int64 `json:"estimate,omitempty"`
	// ExternalLinks an array of external links to be populated.
	ExternalLinks []string `json:"external_links,omitempty"`
	// ExternalTickets an array of the external ticket IDs to be populated.
	ExternalTickets []CreateEntityTemplateExternalTicket `json:"external_tickets,omitempty"`
	// FileIDs an array of the attached file IDs to be populated.
	FileIDs []int64 `json:"file_ids,omitempty"`
	// Files an array of files attached to the story.
	Files []File `json:"files,omitempty"`
	// FollowerIDs an array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	// IterationID the ID of the iteration the to be populated.
	IterationID *int64 `json:"iteration_id,omitempty"`
	// Labels an array of labels to be populated by the template.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// LinkedFileIDs an array of the linked file IDs to be populated.
	LinkedFileIDs []int64 `json:"linked_file_ids,omitempty"`
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
	// Tasks an array of tasks to be populated by the template.
	Tasks []EntityTemplateTask `json:"tasks,omitempty"`
	// WorkflowStateID the ID of the workflow state the story is currently in.
	WorkflowStateID int64 `json:"workflow_state_id,omitempty"`
}

func (m *CreateStoryContents) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateStoryContents) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
