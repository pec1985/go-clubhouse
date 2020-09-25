package models

import (
	"encoding/json"
	"time"
)

type ThreadedComment struct {
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// Deprecated: use member_mention_ids.
	MentionIds []string `json:"mention_ids"`
	// The text of the Comment.
	Text string `json:"text"`
	// The Clubhouse application url for the Comment.
	AppUrl string `json:"app_url"`
	// The unique ID of the Member that authored the Comment.
	AuthorId string `json:"author_id"`
	// True/false boolean indicating whether the Comment is deleted.
	Deleted bool `json:"deleted"`
	// This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id"`
	// An array of Group IDs that have been mentioned in this Comment.
	GroupMentionIds []string `json:"group_mention_ids"`
	// The unique ID of the Comment.
	Id int64 `json:"id"`
	// An array of Member IDs that have been mentioned in this Comment.
	MemberMentionIds []string `json:"member_mention_ids"`
	// The time/date the Comment was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// A nested array of threaded comments.
	Comments []ThreadedComment `json:"comments"`
	// The time/date the Comment was created.
	CreatedAt time.Time `json:"created_at"`
}

func (m *ThreadedComment) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *ThreadedComment) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
