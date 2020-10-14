package models

import (
	"encoding/json"
	"time"
)

// Commit refers to a VCS commit and all associated details.
type Commit struct {
	// AuthorEmail The email address of the VCS user that authored the Commit.
	AuthorEmail string `json:"author_email"`
	// AuthorID The ID of the Member that authored the Commit, if known.
	AuthorID       *string  `json:"author_id"`
	AuthorIdentity Identity `json:"author_identity"`
	// CreatedAt The time/date the Commit was created.
	CreatedAt time.Time `json:"created_at"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// Hash The Commit hash.
	Hash string `json:"hash"`
	// ID The unique ID of the Commit.
	ID *int64 `json:"id"`
	// MergedBranchIDs The IDs of the Branches the Commit has been merged into.
	MergedBranchIDs []int64 `json:"merged_branch_ids"`
	// Message The Commit message.
	Message string `json:"message"`
	// RepositoryID The ID of the Repository that contains the Commit.
	RepositoryID *int64 `json:"repository_id"`
	// Timestamp The time/date the Commit was pushed.
	Timestamp time.Time `json:"timestamp"`
	// UpdatedAt The time/date the Commit was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// URL The URL of the Commit.
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
