package models

import (
	"encoding/json"
	"time"
)

// CreateStoryCommentParams request parameters for creating a Comment on a Clubhouse Story.
type CreateStoryCommentParams struct {
	// AuthorID the Member ID of the Comment's author. Defaults to the user identified by the API token.
	AuthorID string `json:"author_id,omitempty"`
	// CreatedAt defaults to the time/date the comment is created, but can be set to reflect another date.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`
	// Text the comment text.
	Text string `json:"text,omitempty"`
	// UpdatedAt defaults to the time/date the comment is last updated, but can be set to reflect another date.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *CreateStoryCommentParams) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateStoryCommentParams) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
