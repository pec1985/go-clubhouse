package models

import (
	"encoding/json"
	"time"
)

// Comment A Comment is any note added within the Comment field of a Story.
type Comment struct {
	// AppURL The Clubhouse application url for the Comment.
	AppURL string `json:"app_url"`
	// AuthorID The unique ID of the Member who is the Comment's author.
	AuthorID *string `json:"author_id"`
	// CreatedAt The time/date when the Comment was created.
	CreatedAt time.Time `json:"created_at"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// GroupMentionIDs The unique IDs of the Group who are mentioned in the Comment.
	GroupMentionIDs []string `json:"group_mention_ids"`
	// ID The unique ID of the Comment.
	ID int64 `json:"id"`
	// MemberMentionIDs The unique IDs of the Member who are mentioned in the Comment.
	MemberMentionIDs []string `json:"member_mention_ids"`
	// MentionIDs Deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids"`
	// Position The Comments numerical position in the list from oldest to newest.
	Position int64 `json:"position"`
	// Reactions A set of Reactions to this Comment.
	Reactions []Reaction `json:"reactions"`
	// StoryID The ID of the Story on which the Comment appears.
	StoryID int64 `json:"story_id"`
	// Text The text of the Comment.
	Text string `json:"text"`
	// UpdatedAt The time/date when the Comment was updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

func (m *Comment) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Comment) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
