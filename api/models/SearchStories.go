package models

import (
	"encoding/json"
	"time"
)

type SearchStories struct {
	// Archived A true/false boolean indicating whether the Story is in archived state.
	Archived bool `json:"archived"`
	// CompletedAtEnd Stories should have been completed before this date.
	CompletedAtEnd time.Time `json:"completed_at_end"`
	// CompletedAtStart Stories should have been competed after this date.
	CompletedAtStart time.Time `json:"completed_at_start"`
	// CreatedAtEnd Stories should have been created before this date.
	CreatedAtEnd time.Time `json:"created_at_end"`
	// CreatedAtStart Stories should have been created after this date.
	CreatedAtStart time.Time `json:"created_at_start"`
	// DeadlineEnd Stories should have a deadline before this date.
	DeadlineEnd time.Time `json:"deadline_end"`
	// DeadlineStart Stories should have a deadline after this date.
	DeadlineStart time.Time `json:"deadline_start"`
	// EpicID The Epic IDs that may be associated with the Stories.
	EpicID *int64 `json:"epic_id"`
	// EpicIDs The Epic IDs that may be associated with the Stories.
	EpicIDs []int64 `json:"epic_ids"`
	// Estimate The number of estimate points associate with the Stories.
	Estimate int64 `json:"estimate"`
	// ExternalID An ID or URL that references an external resource. Useful during imports.
	ExternalID string `json:"external_id"`
	// GroupID The Group ID that is associated with the Stories
	GroupID *string `json:"group_id"`
	// GroupIDs The Group IDs that are associated with the Stories
	GroupIDs []string `json:"group_ids"`
	// IncludesDescription Whether to include the story description in the response.
	IncludesDescription bool `json:"includes_description"`
	// IterationID The Iteration ID that may be associated with the Stories.
	IterationID *int64 `json:"iteration_id"`
	// IterationIDs The Iteration IDs that may be associated with the Stories.
	IterationIDs []int64 `json:"iteration_ids"`
	// LabelIDs The Label IDs that may be associated with the Stories.
	LabelIDs []int64 `json:"label_ids"`
	// LabelName The name of any associated Labels.
	LabelName string `json:"label_name"`
	// OwnerID An array of UUIDs for any Users who may be Owners of the Stories.
	OwnerID *string `json:"owner_id"`
	// OwnerIDs An array of UUIDs for any Users who may be Owners of the Stories.
	OwnerIDs []string `json:"owner_ids"`
	// ProjectID The IDs for the Projects the Stories may be assigned to.
	ProjectID int64 `json:"project_id"`
	// ProjectIDs The IDs for the Projects the Stories may be assigned to.
	ProjectIDs []int64 `json:"project_ids"`
	// RequestedByID The UUID of any Users who may have requested the Stories.
	RequestedByID string `json:"requested_by_id"`
	// StoryType The type of Stories that you want returned.
	StoryType string `json:"story_type"`
	// UpdatedAtEnd Stories should have been updated before this date.
	UpdatedAtEnd time.Time `json:"updated_at_end"`
	// UpdatedAtStart Stories should have been updated after this date.
	UpdatedAtStart time.Time `json:"updated_at_start"`
	// WorkflowStateID The unique IDs of the specific Workflow States that the Stories should be in.
	WorkflowStateID int64 `json:"workflow_state_id"`
	// WorkflowStateTypes The type of Workflow State the Stories may be in.
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
