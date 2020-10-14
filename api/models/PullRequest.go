package models

import (
	"encoding/json"
	"time"
)

// PullRequest Corresponds to a VCS Pull Request attached to a Clubhouse story.
type PullRequest struct {
	// BranchID The ID of the branch for the particular pull request.
	BranchID int64 `json:"branch_id"`
	// BranchName The name of the branch for the particular pull request.
	BranchName string `json:"branch_name"`
	// BuildStatus The status of the Continuous Integration workflow for the pull request.
	BuildStatus string `json:"build_status"`
	// Closed True/False boolean indicating whether the VCS pull request has been closed.
	Closed bool `json:"closed"`
	// CreatedAt The time/date the pull request was created.
	CreatedAt time.Time `json:"created_at"`
	// Draft True/False boolean indicating whether the VCS pull request is in the draft state.
	Draft bool `json:"draft"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ID The unique ID associated with the pull request in Clubhouse.
	ID int64 `json:"id"`
	// NumAdded Number of lines added in the pull request, according to VCS.
	NumAdded int64 `json:"num_added"`
	// NumCommits The number of commits on the pull request.
	NumCommits *int64 `json:"num_commits"`
	// NumModified Number of lines modified in the pull request, according to VCS.
	NumModified *int64 `json:"num_modified"`
	// NumRemoved Number of lines removed in the pull request, according to VCS.
	NumRemoved int64 `json:"num_removed"`
	// Number The pull requests unique number ID in VCS.
	Number int64 `json:"number"`
	// ReviewStatus The status of the review for the pull request.
	ReviewStatus string `json:"review_status"`
	// TargetBranchID The ID of the target branch for the particular pull request.
	TargetBranchID int64 `json:"target_branch_id"`
	// TargetBranchName The name of the target branch for the particular pull request.
	TargetBranchName string `json:"target_branch_name"`
	// Title The title of the pull request.
	Title string `json:"title"`
	// UpdatedAt The time/date the pull request was created.
	UpdatedAt time.Time `json:"updated_at"`
	// URL The URL for the pull request.
	URL string `json:"url"`
	// VcsLabels An array of PullRequestLabels attached to the PullRequest.
	VcsLabels *[]PullRequestLabel `json:"vcs_labels"`
}

func (m *PullRequest) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *PullRequest) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
