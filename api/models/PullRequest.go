package models

import "time"

// PullRequest corresponds to a VCS Pull Request attached to a Clubhouse story.
type PullRequest struct {
	// BranchID the ID of the branch for the particular pull request.
	BranchID int64 `json:"branch_id,omitempty"`
	// BranchName the name of the branch for the particular pull request.
	BranchName string `json:"branch_name,omitempty"`
	// BuildStatus the status of the Continuous Integration workflow for the pull request.
	BuildStatus string `json:"build_status,omitempty"`
	// Closed true/false boolean indicating whether the VCS pull request has been closed.
	Closed bool `json:"closed,omitempty"`
	// CreatedAt the time/date the pull request was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Draft true/false boolean indicating whether the VCS pull request is in the draft state.
	Draft bool `json:"draft,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ID the unique ID associated with the pull request in Clubhouse.
	ID int64 `json:"id,omitempty"`
	// NumAdded number of lines added in the pull request, according to VCS.
	NumAdded int64 `json:"num_added,omitempty"`
	// NumCommits the number of commits on the pull request.
	NumCommits *int64 `json:"num_commits,omitempty"`
	// NumModified number of lines modified in the pull request, according to VCS.
	NumModified *int64 `json:"num_modified,omitempty"`
	// NumRemoved number of lines removed in the pull request, according to VCS.
	NumRemoved int64 `json:"num_removed,omitempty"`
	// Number the pull requests unique number ID in VCS.
	Number int64 `json:"number,omitempty"`
	// ReviewStatus the status of the review for the pull request.
	ReviewStatus string `json:"review_status,omitempty"`
	// TargetBranchID the ID of the target branch for the particular pull request.
	TargetBranchID int64 `json:"target_branch_id,omitempty"`
	// TargetBranchName the name of the target branch for the particular pull request.
	TargetBranchName string `json:"target_branch_name,omitempty"`
	// Title the title of the pull request.
	Title string `json:"title,omitempty"`
	// UpdatedAt the time/date the pull request was created.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// URL the URL for the pull request.
	URL string `json:"url,omitempty"`
	// VcsLabels an array of PullRequestLabels attached to the PullRequest.
	VcsLabels *[]PullRequestLabel `json:"vcs_labels,omitempty"`
}

func (m *PullRequest) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *PullRequest) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
