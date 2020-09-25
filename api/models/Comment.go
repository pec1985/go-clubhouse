package models

import (
	"encoding/json"
	"time"
)

type Comment struct {
	// This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id"`
	// A set of Reactions to this Comment.
	Reactions []Reaction `json:"reactions"`
	// The time/date when the Comment was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// The unique IDs of the Group who are mentioned in the Comment.
	GroupMentionIds []string `json:"group_mention_ids"`
	// The time/date when the Comment was created.
	CreatedAt time.Time `json:"created_at"`
	// The unique ID of the Comment.
	Id int64 `json:"id"`
	// The unique IDs of the Member who are mentioned in the Comment.
	MemberMentionIds []string `json:"member_mention_ids"`
	// Deprecated: use member_mention_ids.
	MentionIds []string `json:"mention_ids"`
	// The Comments numerical position in the list from oldest to newest.
	Position int64 `json:"position"`
	// The text of the Comment.
	Text string `json:"text"`
	// The Clubhouse application url for the Comment.
	AppUrl string `json:"app_url"`
	// The unique ID of the Member who is the Comment's author.
	AuthorId *string `json:"author_id"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The ID of the Story on which the Comment appears.
	StoryId int64 `json:"story_id"`
}

func (m *Comment) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Comment) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
