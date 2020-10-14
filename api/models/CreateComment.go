package models

import (
	"encoding/json"
	"time"
)

type CreateComment struct {
	// AuthorID The Member ID of the Comment's author. Defaults to the user identified by the API token.
	AuthorID string `json:"author_id"`
	// CreatedAt Defaults to the time/date the comment is created, but can be set to reflect another date.
	CreatedAt time.Time `json:"created_at"`
	// ExternalID This field can be set to another unique ID. In the case that the comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id"`
	// Text The comment text.
	Text string `json:"text"`
	// UpdatedAt Defaults to the time/date the comment is last updated, but can be set to reflect another date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *CreateComment) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateComment) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
