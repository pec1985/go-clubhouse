package models

import (
	"encoding/json"
	"time"
)

// PullRequest corresponds to a VCS Pull Request attached to a Clubhouse story.
type PullRequest struct {
	// BranchID the ID of the branch for the particular pull request.
	BranchID int64 `json:"branch_id"`
	// BranchName the name of the branch for the particular pull request.
	BranchName string `json:"branch_name"`
	// BuildStatus the status of the Continuous Integration workflow for the pull request.
	BuildStatus string `json:"build_status"`
	// Closed true/false boolean indicating whether the VCS pull request has been closed.
	Closed bool `json:"closed"`
	// CreatedAt the time/date the pull request was created.
	CreatedAt time.Time `json:"created_at"`
	// Draft true/false boolean indicating whether the VCS pull request is in the draft state.
	Draft bool `json:"draft"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ID the unique ID associated with the pull request in Clubhouse.
	ID int64 `json:"id"`
	// NumAdded number of lines added in the pull request, according to VCS.
	NumAdded int64 `json:"num_added"`
	// NumCommits the number of commits on the pull request.
	NumCommits *int64 `json:"num_commits"`
	// NumModified number of lines modified in the pull request, according to VCS.
	NumModified *int64 `json:"num_modified"`
	// NumRemoved number of lines removed in the pull request, according to VCS.
	NumRemoved int64 `json:"num_removed"`
	// Number the pull requests unique number ID in VCS.
	Number int64 `json:"number"`
	// ReviewStatus the status of the review for the pull request.
	ReviewStatus string `json:"review_status"`
	// TargetBranchID the ID of the target branch for the particular pull request.
	TargetBranchID int64 `json:"target_branch_id"`
	// TargetBranchName the name of the target branch for the particular pull request.
	TargetBranchName string `json:"target_branch_name"`
	// Title the title of the pull request.
	Title string `json:"title"`
	// UpdatedAt the time/date the pull request was created.
	UpdatedAt time.Time `json:"updated_at"`
	// URL the URL for the pull request.
	URL string `json:"url"`
	// VcsLabels an array of PullRequestLabels attached to the PullRequest.
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
