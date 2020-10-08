package models

import (
	"encoding/json"
	"time"
)

type UpdateStories struct {
	// The ID of the story that the stories are to be moved below.
	AfterId int64 `json:"after_id"`
	// If the Stories should be archived or not.
	Archived bool `json:"archived"`
	// The ID of the story that the stories are to be moved before.
	BeforeId int64 `json:"before_id"`
	// The due date of the story.
	Deadline *time.Time `json:"deadline"`
	// The ID of the epic the story belongs to.
	EpicId *int64 `json:"epic_id"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate"`
	// An array of External Links associated with this story.
	ExternalLinks []string `json:"external_links"`
	// The UUIDs of the new followers to be added.
	FollowerIdsAdd []string `json:"follower_ids_add"`
	// The UUIDs of the followers to be removed.
	FollowerIdsRemove []string `json:"follower_ids_remove"`
	// The Id of the Group the Stories should belong to.
	GroupId *string `json:"group_id"`
	// The ID of the iteration the story belongs to.
	IterationId *int64 `json:"iteration_id"`
	// An array of labels to be added.
	LabelsAdd []CreateLabelParams `json:"labels_add"`
	// An array of labels to be removed.
	LabelsRemove []CreateLabelParams `json:"labels_remove"`
	MoveTo       string              `json:"move_to"`
	// The UUIDs of the new owners to be added.
	OwnerIdsAdd []string `json:"owner_ids_add"`
	// The UUIDs of the owners to be removed.
	OwnerIdsRemove []string `json:"owner_ids_remove"`
	// The ID of the Project the Stories should belong to.
	ProjectId int64 `json:"project_id"`
	// The ID of the member that requested the story.
	RequestedById string `json:"requested_by_id"`
	// The unique IDs of the Stories you wish to update.
	StoryIds []int64 `json:"story_ids"`
	// The type of story (feature, bug, chore).
	StoryType string `json:"story_type"`
	// The ID of the workflow state to put the stories in.
	WorkflowStateId int64 `json:"workflow_state_id"`
}

func (m *UpdateStories) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateStories) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
