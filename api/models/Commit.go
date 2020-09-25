package models

import (
	"encoding/json"
	"time"
)

type Commit struct {
	// The email address of the VCS user that authored the Commit.
	AuthorEmail string `json:"author_email"`
	// The ID of the Member that authored the Commit, if known.
	AuthorId       *string  `json:"author_id"`
	AuthorIdentity Identity `json:"author_identity"`
	// The time/date the Commit was created.
	CreatedAt time.Time `json:"created_at"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The Commit hash.
	Hash string `json:"hash"`
	// The unique ID of the Commit.
	Id *int64 `json:"id"`
	// The IDs of the Branches the Commit has been merged into.
	MergedBranchIds []int64 `json:"merged_branch_ids"`
	// The Commit message.
	Message string `json:"message"`
	// The ID of the Repository that contains the Commit.
	RepositoryId *int64 `json:"repository_id"`
	// The time/date the Commit was pushed.
	Timestamp time.Time `json:"timestamp"`
	// The time/date the Commit was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// The URL of the Commit.
	Url string `json:"url"`
}

func (m *Commit) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Commit) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
