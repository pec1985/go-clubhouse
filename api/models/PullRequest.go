package models

import (
	"encoding/json"
	"time"
)

type PullRequest struct {
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique ID associated with the pull request in Clubhouse.
	Id int64 `json:"id"`
	// The number of commits on the pull request.
	NumCommits *int64 `json:"num_commits"`
	// The ID of the target branch for the particular pull request.
	TargetBranchId int64 `json:"target_branch_id"`
	// The name of the branch for the particular pull request.
	BranchName string `json:"branch_name"`
	// The status of the Continuous Integration workflow for the pull request.
	BuildStatus string `json:"build_status"`
	// True/False boolean indicating whether the VCS pull request has been closed.
	Closed bool `json:"closed"`
	// Number of lines added in the pull request, according to VCS.
	NumAdded int64 `json:"num_added"`
	// The name of the target branch for the particular pull request.
	TargetBranchName string `json:"target_branch_name"`
	// The title of the pull request.
	Title string `json:"title"`
	// The status of the review for the pull request.
	ReviewStatus string `json:"review_status"`
	// An array of PullRequestLabels attached to the PullRequest.
	VcsLabels *[]PullRequestLabel `json:"vcs_labels"`
	// Number of lines modified in the pull request, according to VCS.
	NumModified *int64 `json:"num_modified"`
	// Number of lines removed in the pull request, according to VCS.
	NumRemoved int64 `json:"num_removed"`
	// The pull requests unique number ID in VCS.
	Number int64 `json:"number"`
	// The time/date the pull request was created.
	UpdatedAt time.Time `json:"updated_at"`
	// The URL for the pull request.
	Url string `json:"url"`
	// The ID of the branch for the particular pull request.
	BranchId int64 `json:"branch_id"`
	// The time/date the pull request was created.
	CreatedAt time.Time `json:"created_at"`
	// True/False boolean indicating whether the VCS pull request is in the draft state.
	Draft bool `json:"draft"`
}

func (m *PullRequest) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *PullRequest) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
