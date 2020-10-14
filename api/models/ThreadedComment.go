package models

import (
	"encoding/json"
	"time"
)

// ThreadedComment comments associated with Epic Discussions.
type ThreadedComment struct {
	// AppURL the Clubhouse application url for the Comment.
	AppURL string `json:"app_url"`
	// AuthorID the unique ID of the Member that authored the Comment.
	AuthorID string `json:"author_id"`
	// Comments a nested array of threaded comments.
	Comments []ThreadedComment `json:"comments"`
	// CreatedAt the time/date the Comment was created.
	CreatedAt time.Time `json:"created_at"`
	// Deleted true/false boolean indicating whether the Comment is deleted.
	Deleted bool `json:"deleted"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID this field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// GroupMentionIDs an array of Group IDs that have been mentioned in this Comment.
	GroupMentionIDs []string `json:"group_mention_ids"`
	// ID the unique ID of the Comment.
	ID int64 `json:"id"`
	// MemberMentionIDs an array of Member IDs that have been mentioned in this Comment.
	MemberMentionIDs []string `json:"member_mention_ids"`
	// MentionIDs deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids"`
	// Text the text of the Comment.
	Text string `json:"text"`
	// UpdatedAt the time/date the Comment was updated.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ThreadedComment) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *ThreadedComment) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
