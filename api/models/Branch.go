package models

import (
	"encoding/json"
	"time"
)

// branch refers to a VCS branch. Branches are feature branches associated with Clubhouse Stories.
type Branch struct {
	// CreatedAt the time/date the Branch was created.
	CreatedAt *time.Time `json:"created_at"`
	// Deleted a true/false boolean indicating if the Branch has been deleted.
	Deleted bool `json:"deleted"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ID the unique ID of the Branch.
	ID *int64 `json:"id"`
	// MergedBranchIDs the IDs of the Branches the Branch has been merged into.
	MergedBranchIDs []int64 `json:"merged_branch_ids"`
	// Name the name of the Branch.
	Name string `json:"name"`
	// Persistent a true/false boolean indicating if the Branch is persistent; e.g. master.
	Persistent bool `json:"persistent"`
	// PullRequests an array of PullRequests attached to the Branch (there is usually only one).
	PullRequests []PullRequest `json:"pull_requests"`
	// RepositoryID the ID of the Repository that contains the Branch.
	RepositoryID *int64 `json:"repository_id"`
	// UpdatedAt the time/date the Branch was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// URL the URL of the Branch.
	URL string `json:"url"`
}

func (m *Branch) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Branch) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
