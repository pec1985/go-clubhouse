package models

import "time"

// branch refers to a VCS branch. Branches are feature branches associated with Clubhouse Stories.
type Branch struct {
	// CreatedAt the time/date the Branch was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Deleted a true/false boolean indicating if the Branch has been deleted.
	Deleted bool `json:"deleted,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ID the unique ID of the Branch.
	ID *int64 `json:"id,omitempty"`
	// MergedBranchIDs the IDs of the Branches the Branch has been merged into.
	MergedBranchIDs []int64 `json:"merged_branch_ids,omitempty"`
	// Name the name of the Branch.
	Name string `json:"name,omitempty"`
	// Persistent a true/false boolean indicating if the Branch is persistent; e.g. master.
	Persistent bool `json:"persistent,omitempty"`
	// PullRequests an array of PullRequests attached to the Branch (there is usually only one).
	PullRequests []PullRequest `json:"pull_requests,omitempty"`
	// RepositoryID the ID of the Repository that contains the Branch.
	RepositoryID *int64 `json:"repository_id,omitempty"`
	// UpdatedAt the time/date the Branch was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// URL the URL of the Branch.
	URL string `json:"url,omitempty"`
}

func (m *Branch) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Branch) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
