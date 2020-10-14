package models

import (
	"encoding/json"
	"time"
)

type UpdateStories struct {
	// AfterID the ID of the story that the stories are to be moved below.
	AfterID int64 `json:"after_id"`
	// Archived if the Stories should be archived or not.
	Archived bool `json:"archived"`
	// BeforeID the ID of the story that the stories are to be moved before.
	BeforeID int64 `json:"before_id"`
	// Deadline the due date of the story.
	Deadline *time.Time `json:"deadline"`
	// EpicID the ID of the epic the story belongs to.
	EpicID *int64 `json:"epic_id"`
	// Estimate the numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate"`
	// ExternalLinks an array of External Links associated with this story.
	ExternalLinks []string `json:"external_links"`
	// FollowerIdsAdd the UUIDs of the new followers to be added.
	FollowerIdsAdd []string `json:"follower_ids_add"`
	// FollowerIdsRemove the UUIDs of the followers to be removed.
	FollowerIdsRemove []string `json:"follower_ids_remove"`
	// GroupID the Id of the Group the Stories should belong to.
	GroupID *string `json:"group_id"`
	// IterationID the ID of the iteration the story belongs to.
	IterationID *int64 `json:"iteration_id"`
	// LabelsAdd an array of labels to be added.
	LabelsAdd []CreateLabelParams `json:"labels_add"`
	// LabelsRemove an array of labels to be removed.
	LabelsRemove []CreateLabelParams `json:"labels_remove"`
	MoveTo       string              `json:"move_to"`
	// OwnerIdsAdd the UUIDs of the new owners to be added.
	OwnerIdsAdd []string `json:"owner_ids_add"`
	// OwnerIdsRemove the UUIDs of the owners to be removed.
	OwnerIdsRemove []string `json:"owner_ids_remove"`
	// ProjectID the ID of the Project the Stories should belong to.
	ProjectID int64 `json:"project_id"`
	// RequestedByID the ID of the member that requested the story.
	RequestedByID string `json:"requested_by_id"`
	// StoryIDs the Ids of the Stories you wish to update.
	StoryIDs []int64 `json:"story_ids"`
	// StoryType the type of story (feature, bug, chore).
	StoryType string `json:"story_type"`
	// WorkflowStateID the ID of the workflow state to put the stories in.
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
