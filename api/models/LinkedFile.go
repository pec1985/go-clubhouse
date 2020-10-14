package models

import (
	"encoding/json"
	"time"
)

// LinkedFile Linked files are stored on a third-party website and linked to one or more Stories. Clubhouse currently supports linking files from Google Drive, Dropbox, Box, and by URL.
type LinkedFile struct {
	// ContentType The content type of the image (e.g. txt/plain).
	ContentType *string `json:"content_type"`
	// CreatedAt The time/date the LinkedFile was created.
	CreatedAt time.Time `json:"created_at"`
	// Description The description of the file.
	Description *string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// GroupMentionIDs The groups that are mentioned in the description of the file.
	GroupMentionIDs []string `json:"group_mention_ids"`
	// ID The unique identifier for the file.
	ID int64 `json:"id"`
	// MemberMentionIDs The members that are mentioned in the description of the file.
	MemberMentionIDs []string `json:"member_mention_ids"`
	// MentionIDs Deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids"`
	// Name The name of the linked file.
	Name string `json:"name"`
	// Size The filesize, if the integration provided it.
	Size *int64 `json:"size"`
	// StoryIDs The IDs of the stories this file is attached to.
	StoryIDs []int64 `json:"story_ids"`
	// ThumbnailURL The URL of the file thumbnail, if the integration provided it.
	ThumbnailURL *string `json:"thumbnail_url"`
	// Type The integration type (e.g. google, dropbox, box).
	Type string `json:"type"`
	// UpdatedAt The time/date the LinkedFile was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UploaderID The UUID of the member that uploaded the file.
	UploaderID string `json:"uploader_id"`
	// URL The URL of the file.
	URL string `json:"url"`
}

func (m *LinkedFile) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *LinkedFile) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
