package models

import "time"

// commit refers to a VCS commit and all associated details.
type Commit struct {
	// AuthorEmail the email address of the VCS user that authored the Commit.
	AuthorEmail string `json:"author_email,omitempty"`
	// AuthorID the ID of the Member that authored the Commit, if known.
	AuthorID       *string  `json:"author_id,omitempty"`
	AuthorIdentity Identity `json:"author_identity,omitempty"`
	// CreatedAt the time/date the Commit was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// Hash the Commit hash.
	Hash string `json:"hash,omitempty"`
	// ID the unique ID of the Commit.
	ID *int64 `json:"id,omitempty"`
	// MergedBranchIDs the IDs of the Branches the Commit has been merged into.
	MergedBranchIDs []int64 `json:"merged_branch_ids,omitempty"`
	// Message the Commit message.
	Message string `json:"message,omitempty"`
	// RepositoryID the ID of the Repository that contains the Commit.
	RepositoryID *int64 `json:"repository_id,omitempty"`
	// Timestamp the time/date the Commit was pushed.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// UpdatedAt the time/date the Commit was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// URL the URL of the Commit.
	URL string `json:"url,omitempty"`
}

func (m *Commit) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Commit) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
