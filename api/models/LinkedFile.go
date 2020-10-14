package models

import (
	"encoding/json"
	"time"
)

// LinkedFile linked files are stored on a third-party website and linked to one or more Stories. Clubhouse currently supports linking files from Google Drive, Dropbox, Box, and by URL.
type LinkedFile struct {
	// ContentType the content type of the image (e.g. txt/plain).
	ContentType *string `json:"content_type"`
	// CreatedAt the time/date the LinkedFile was created.
	CreatedAt time.Time `json:"created_at"`
	// Description the description of the file.
	Description *string `json:"description"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// GroupMentionIDs the groups that are mentioned in the description of the file.
	GroupMentionIDs []string `json:"group_mention_ids"`
	// ID the unique identifier for the file.
	ID int64 `json:"id"`
	// MemberMentionIDs the members that are mentioned in the description of the file.
	MemberMentionIDs []string `json:"member_mention_ids"`
	// MentionIDs deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids"`
	// Name the name of the linked file.
	Name string `json:"name"`
	// Size the filesize, if the integration provided it.
	Size *int64 `json:"size"`
	// StoryIDs the IDs of the stories this file is attached to.
	StoryIDs []int64 `json:"story_ids"`
	// ThumbnailURL the URL of the file thumbnail, if the integration provided it.
	ThumbnailURL *string `json:"thumbnail_url"`
	// Type the integration type (e.g. google, dropbox, box).
	Type string `json:"type"`
	// UpdatedAt the time/date the LinkedFile was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UploaderID the UUID of the member that uploaded the file.
	UploaderID string `json:"uploader_id"`
	// URL the URL of the file.
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
