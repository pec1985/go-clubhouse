package models

import "time"

// Comment a Comment is any note added within the Comment field of a Story.
type Comment struct {
	// AppURL the Clubhouse application url for the Comment.
	AppURL string `json:"app_url,omitempty"`
	// AuthorID the unique ID of the Member who is the Comment's author.
	AuthorID *string `json:"author_id,omitempty"`
	// CreatedAt the time/date when the Comment was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id,omitempty"`
	// GroupMentionIDs the unique IDs of the Group who are mentioned in the Comment.
	GroupMentionIDs []string `json:"group_mention_ids,omitempty"`
	// ID the unique ID of the Comment.
	ID int64 `json:"id,omitempty"`
	// MemberMentionIDs the unique IDs of the Member who are mentioned in the Comment.
	MemberMentionIDs []string `json:"member_mention_ids,omitempty"`
	// MentionIDs deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids,omitempty"`
	// Position the Comments numerical position in the list from oldest to newest.
	Position int64 `json:"position,omitempty"`
	// Reactions a set of Reactions to this Comment.
	Reactions []Reaction `json:"reactions,omitempty"`
	// StoryID the ID of the Story on which the Comment appears.
	StoryID int64 `json:"story_id,omitempty"`
	// Text the text of the Comment.
	Text string `json:"text,omitempty"`
	// UpdatedAt the time/date when the Comment was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (m *Comment) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Comment) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
