package models

import (
	"encoding/json"
	"time"
)

type UpdateStories struct {
	// AfterID The ID of the story that the stories are to be moved below.
	AfterID int64 `json:"after_id"`
	// Archived If the Stories should be archived or not.
	Archived bool `json:"archived"`
	// BeforeID The ID of the story that the stories are to be moved before.
	BeforeID int64 `json:"before_id"`
	// Deadline The due date of the story.
	Deadline *time.Time `json:"deadline"`
	// EpicID The ID of the epic the story belongs to.
	EpicID *int64 `json:"epic_id"`
	// Estimate The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate"`
	// ExternalLinks An array of External Links associated with this story.
	ExternalLinks []string `json:"external_links"`
	// FollowerIdsAdd The UUIDs of the new followers to be added.
	FollowerIdsAdd []string `json:"follower_ids_add"`
	// FollowerIdsRemove The UUIDs of the followers to be removed.
	FollowerIdsRemove []string `json:"follower_ids_remove"`
	// GroupID The Id of the Group the Stories should belong to.
	GroupID *string `json:"group_id"`
	// IterationID The ID of the iteration the story belongs to.
	IterationID *int64 `json:"iteration_id"`
	// LabelsAdd An array of labels to be added.
	LabelsAdd []CreateLabelParams `json:"labels_add"`
	// LabelsRemove An array of labels to be removed.
	LabelsRemove []CreateLabelParams `json:"labels_remove"`
	MoveTo       string              `json:"move_to"`
	// OwnerIdsAdd The UUIDs of the new owners to be added.
	OwnerIdsAdd []string `json:"owner_ids_add"`
	// OwnerIdsRemove The UUIDs of the owners to be removed.
	OwnerIdsRemove []string `json:"owner_ids_remove"`
	// ProjectID The ID of the Project the Stories should belong to.
	ProjectID int64 `json:"project_id"`
	// RequestedByID The ID of the member that requested the story.
	RequestedByID string `json:"requested_by_id"`
	// StoryIDs The Ids of the Stories you wish to update.
	StoryIDs []int64 `json:"story_ids"`
	// StoryType The type of story (feature, bug, chore).
	StoryType string `json:"story_type"`
	// WorkflowStateID The ID of the workflow state to put the stories in.
	WorkflowStateID int64 `json:"workflow_state_id"`
}

func (m *UpdateStories) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateStories) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
