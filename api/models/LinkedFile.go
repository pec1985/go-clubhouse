package models

import (
	"encoding/json"
	"time"
)

type LinkedFile struct {
	// The groups that are mentioned in the description of the file.
	GroupMentionIds []string `json:"group_mention_ids"`
	// The members that are mentioned in the description of the file.
	MemberMentionIds []string `json:"member_mention_ids"`
	// Deprecated: use member_mention_ids.
	MentionIds []string `json:"mention_ids"`
	// The time/date the LinkedFile was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The unique identifier for the file.
	Id int64 `json:"id"`
	// The content type of the image (e.g. txt/plain).
	ContentType *string `json:"content_type"`
	// The time/date the LinkedFile was created.
	CreatedAt time.Time `json:"created_at"`
	// The name of the linked file.
	Name string `json:"name"`
	// The IDs of the stories this file is attached to.
	StoryIds []int64 `json:"story_ids"`
	// The URL of the file thumbnail, if the integration provided it.
	ThumbnailUrl *string `json:"thumbnail_url"`
	// The integration type (e.g. google, dropbox, box).
	Type string `json:"type"`
	// The description of the file.
	Description *string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The filesize, if the integration provided it.
	Size *int64 `json:"size"`
	// The UUID of the member that uploaded the file.
	UploaderId string `json:"uploader_id"`
	// The URL of the file.
	Url string `json:"url"`
}

func (m *LinkedFile) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *LinkedFile) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
