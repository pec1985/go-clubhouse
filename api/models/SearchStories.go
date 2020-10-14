package models

import (
	"encoding/json"
	"time"
)

type SearchStories struct {
	// Archived a true/false boolean indicating whether the Story is in archived state.
	Archived bool `json:"archived"`
	// CompletedAtEnd stories should have been completed before this date.
	CompletedAtEnd time.Time `json:"completed_at_end"`
	// CompletedAtStart stories should have been competed after this date.
	CompletedAtStart time.Time `json:"completed_at_start"`
	// CreatedAtEnd stories should have been created before this date.
	CreatedAtEnd time.Time `json:"created_at_end"`
	// CreatedAtStart stories should have been created after this date.
	CreatedAtStart time.Time `json:"created_at_start"`
	// DeadlineEnd stories should have a deadline before this date.
	DeadlineEnd time.Time `json:"deadline_end"`
	// DeadlineStart stories should have a deadline after this date.
	DeadlineStart time.Time `json:"deadline_start"`
	// EpicID the Epic IDs that may be associated with the Stories.
	EpicID *int64 `json:"epic_id"`
	// EpicIDs the Epic IDs that may be associated with the Stories.
	EpicIDs []int64 `json:"epic_ids"`
	// Estimate the number of estimate points associate with the Stories.
	Estimate int64 `json:"estimate"`
	// ExternalID an ID or URL that references an external resource. Useful during imports.
	ExternalID string `json:"external_id"`
	// GroupID the Group ID that is associated with the Stories
	GroupID *string `json:"group_id"`
	// GroupIDs the Group IDs that are associated with the Stories
	GroupIDs []string `json:"group_ids"`
	// IncludesDescription whether to include the story description in the response.
	IncludesDescription bool `json:"includes_description"`
	// IterationID the Iteration ID that may be associated with the Stories.
	IterationID *int64 `json:"iteration_id"`
	// IterationIDs the Iteration IDs that may be associated with the Stories.
	IterationIDs []int64 `json:"iteration_ids"`
	// LabelIDs the Label IDs that may be associated with the Stories.
	LabelIDs []int64 `json:"label_ids"`
	// LabelName the name of any associated Labels.
	LabelName string `json:"label_name"`
	// OwnerID an array of UUIDs for any Users who may be Owners of the Stories.
	OwnerID *string `json:"owner_id"`
	// OwnerIDs an array of UUIDs for any Users who may be Owners of the Stories.
	OwnerIDs []string `json:"owner_ids"`
	// ProjectID the IDs for the Projects the Stories may be assigned to.
	ProjectID int64 `json:"project_id"`
	// ProjectIDs the IDs for the Projects the Stories may be assigned to.
	ProjectIDs []int64 `json:"project_ids"`
	// RequestedByID the UUID of any Users who may have requested the Stories.
	RequestedByID string `json:"requested_by_id"`
	// StoryType the type of Stories that you want returned.
	StoryType string `json:"story_type"`
	// UpdatedAtEnd stories should have been updated before this date.
	UpdatedAtEnd time.Time `json:"updated_at_end"`
	// UpdatedAtStart stories should have been updated after this date.
	UpdatedAtStart time.Time `json:"updated_at_start"`
	// WorkflowStateID the unique IDs of the specific Workflow States that the Stories should be in.
	WorkflowStateID int64 `json:"workflow_state_id"`
	// WorkflowStateTypes the type of Workflow State the Stories may be in.
	WorkflowStateTypes []string `json:"workflow_state_types"`
}

func (m *SearchStories) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *SearchStories) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
