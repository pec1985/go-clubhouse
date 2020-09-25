package models

import (
	"encoding/json"
	"time"
)

type UpdateStory struct {
	// The type of story (feature, bug, chore).
	StoryType string `json:"story_type"`
	// The due date of the story.
	Deadline *time.Time `json:"deadline"`
	// The description of the story.
	Description string `json:"description"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate"`
	// An array of IDs of files attached to the story.
	FileIds []int64 `json:"file_ids"`
	// An array of UUIDs of the followers of this story.
	FollowerIds []string `json:"follower_ids"`
	// The ID of the group to associate with this story
	GroupId *string `json:"group_id"`
	// The title of the story.
	Name string `json:"name"`
	// The ID of the member that requested the story.
	RequestedById string `json:"requested_by_id"`
	// An array of IDs of Branches attached to the story.
	BranchIds []int64 `json:"branch_ids"`
	// The ID of the iteration the story belongs to.
	IterationId *int64 `json:"iteration_id"`
	// An array of IDs of linked files attached to the story.
	LinkedFileIds []int64 `json:"linked_file_ids"`
	// The ID of the project the story belongs to.
	ProjectId int64 `json:"project_id"`
	// The ID of the story we want to move this story after.
	AfterId int64 `json:"after_id"`
	// The ID of the story we want to move this story before.
	BeforeId int64 `json:"before_id"`
	// An array of IDs of Commits attached to the story.
	CommitIds []int64 `json:"commit_ids"`
	// A manual override for the time/date the Story was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// The ID of the epic the story belongs to.
	EpicId *int64 `json:"epic_id"`
	// An array of labels attached to the story.
	Labels []CreateLabelParams `json:"labels"`
	MoveTo string              `json:"move_to"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids"`
	// True if the story is archived, otherwise false.
	Archived bool `json:"archived"`
	// A manual override for the time/date the Story was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	// The ID of the workflow state to put the story in.
	WorkflowStateId int64 `json:"workflow_state_id"`
	// An array of IDs of Pull/Merge Requests attached to the story.
	PullRequestIds []int64 `json:"pull_request_ids"`
}

func (m *UpdateStory) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateStory) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
