package models

import "time"

type UpdateStory struct {
	// AfterID the ID of the story we want to move this story after.
	AfterID int64 `json:"after_id,omitempty"`
	// Archived true if the story is archived, otherwise false.
	Archived bool `json:"archived,omitempty"`
	// BeforeID the ID of the story we want to move this story before.
	BeforeID int64 `json:"before_id,omitempty"`
	// BranchIDs an array of IDs of Branches attached to the story.
	BranchIDs []int64 `json:"branch_ids,omitempty"`
	// CommitIDs an array of IDs of Commits attached to the story.
	CommitIDs []int64 `json:"commit_ids,omitempty"`
	// CompletedAtOverride a manual override for the time/date the Story was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override,omitempty"`
	// Deadline the due date of the story.
	Deadline *time.Time `json:"deadline,omitempty"`
	// Description the description of the story.
	Description string `json:"description,omitempty"`
	// EpicID the ID of the epic the story belongs to.
	EpicID *int64 `json:"epic_id,omitempty"`
	// Estimate the numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate,omitempty"`
	// ExternalLinks an array of External Links associated with this story.
	ExternalLinks []string `json:"external_links,omitempty"`
	// FileIDs an array of IDs of files attached to the story.
	FileIDs []int64 `json:"file_ids,omitempty"`
	// FollowerIDs an array of UUIDs of the followers of this story.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	// GroupID the ID of the group to associate with this story
	GroupID *string `json:"group_id,omitempty"`
	// IterationID the ID of the iteration the story belongs to.
	IterationID *int64 `json:"iteration_id,omitempty"`
	// Labels an array of labels attached to the story.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// LinkedFileIDs an array of IDs of linked files attached to the story.
	LinkedFileIDs []int64 `json:"linked_file_ids,omitempty"`
	MoveTo        string  `json:"move_to,omitempty"`
	// Name the title of the story.
	Name string `json:"name,omitempty"`
	// OwnerIDs an array of UUIDs of the owners of this story.
	OwnerIDs []string `json:"owner_ids,omitempty"`
	// ProjectID the ID of the project the story belongs to.
	ProjectID int64 `json:"project_id,omitempty"`
	// PullRequestIDs an array of IDs of Pull/Merge Requests attached to the story.
	PullRequestIDs []int64 `json:"pull_request_ids,omitempty"`
	// RequestedByID the ID of the member that requested the story.
	RequestedByID string `json:"requested_by_id,omitempty"`
	// StartedAtOverride a manual override for the time/date the Story was started.
	StartedAtOverride *time.Time `json:"started_at_override,omitempty"`
	// StoryType the type of story (feature, bug, chore).
	StoryType string `json:"story_type,omitempty"`
	// WorkflowStateID the ID of the workflow state to put the story in.
	WorkflowStateID int64 `json:"workflow_state_id,omitempty"`
}

func (m *UpdateStory) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateStory) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
