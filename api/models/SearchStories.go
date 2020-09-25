package models

import (
	"encoding/json"
	"time"
)

type SearchStories struct {
	// The UUID of any Users who may have requested the Stories.
	RequestedById string `json:"requested_by_id"`
	// Stories should have been competed after this date.
	CompletedAtStart time.Time `json:"completed_at_start"`
	// The number of estimate points associate with the Stories.
	Estimate int64 `json:"estimate"`
	// The Group IDs that are associated with the Stories
	GroupIds []string `json:"group_ids"`
	// The Iteration IDs that may be associated with the Stories.
	IterationIds []int64 `json:"iteration_ids"`
	// The name of any associated Labels.
	LabelName string `json:"label_name"`
	// An array of UUIDs for any Users who may be Owners of the Stories.
	OwnerId *string `json:"owner_id"`
	// The IDs for the Projects the Stories may be assigned to.
	ProjectId int64 `json:"project_id"`
	// Stories should have been completed before this date.
	CompletedAtEnd time.Time `json:"completed_at_end"`
	// Stories should have been created before this date.
	CreatedAtEnd time.Time `json:"created_at_end"`
	// An ID or URL that references an external resource. Useful during imports.
	ExternalId string `json:"external_id"`
	// The IDs for the Projects the Stories may be assigned to.
	ProjectIds []int64 `json:"project_ids"`
	// Stories should have been updated after this date.
	UpdatedAtStart time.Time `json:"updated_at_start"`
	// An array of UUIDs for any Users who may be Owners of the Stories.
	OwnerIds []string `json:"owner_ids"`
	// A true/false boolean indicating whether the Story is in archived state.
	Archived bool `json:"archived"`
	// Stories should have been created after this date.
	CreatedAtStart time.Time `json:"created_at_start"`
	// Stories should have a deadline before this date.
	DeadlineEnd time.Time `json:"deadline_end"`
	// The Epic IDs that may be associated with the Stories.
	EpicId *int64 `json:"epic_id"`
	// Stories should have been updated before this date.
	UpdatedAtEnd time.Time `json:"updated_at_end"`
	// Stories should have a deadline after this date.
	DeadlineStart time.Time `json:"deadline_start"`
	// The Group ID that is associated with the Stories
	GroupId *string `json:"group_id"`
	// The unique IDs of the specific Workflow States that the Stories should be in.
	WorkflowStateId int64 `json:"workflow_state_id"`
	// The type of Workflow State the Stories may be in.
	WorkflowStateTypes []string `json:"workflow_state_types"`
	// Whether to include the story description in the response.
	IncludesDescription bool `json:"includes_description"`
	// The type of Stories that you want returned.
	StoryType string `json:"story_type"`
	// The Iteration ID that may be associated with the Stories.
	IterationId *int64 `json:"iteration_id"`
	// The Label IDs that may be associated with the Stories.
	LabelIds []int64 `json:"label_ids"`
	// The Epic IDs that may be associated with the Stories.
	EpicIds []int64 `json:"epic_ids"`
}

func (m *SearchStories) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *SearchStories) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
