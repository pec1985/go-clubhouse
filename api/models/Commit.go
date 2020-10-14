package models

import (
	"encoding/json"
	"time"
)

// commit refers to a VCS commit and all associated details.
type Commit struct {
	// AuthorEmail the email address of the VCS user that authored the Commit.
	AuthorEmail string `json:"author_email"`
	// AuthorID the ID of the Member that authored the Commit, if known.
	AuthorID       *string  `json:"author_id"`
	AuthorIdentity Identity `json:"author_identity"`
	// CreatedAt the time/date the Commit was created.
	CreatedAt time.Time `json:"created_at"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// Hash the Commit hash.
	Hash string `json:"hash"`
	// ID the unique ID of the Commit.
	ID *int64 `json:"id"`
	// MergedBranchIDs the IDs of the Branches the Commit has been merged into.
	MergedBranchIDs []int64 `json:"merged_branch_ids"`
	// Message the Commit message.
	Message string `json:"message"`
	// RepositoryID the ID of the Repository that contains the Commit.
	RepositoryID *int64 `json:"repository_id"`
	// Timestamp the time/date the Commit was pushed.
	Timestamp time.Time `json:"timestamp"`
	// UpdatedAt the time/date the Commit was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// URL the URL of the Commit.
	URL string `json:"url"`
}

func (m *Commit) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Commit) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
