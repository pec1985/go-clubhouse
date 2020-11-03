package models

import (
	"encoding/json"
	"time"
)

// LinkedFile linked files are stored on a third-party website and linked to one or more Stories. Clubhouse currently supports linking files from Google Drive, Dropbox, Box, and by URL.
type LinkedFile struct {
	// ContentType the content type of the image (e.g. txt/plain).
	ContentType *string `json:"content_type,omitempty"`
	// CreatedAt the time/date the LinkedFile was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Description the description of the file.
	Description *string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// GroupMentionIDs the groups that are mentioned in the description of the file.
	GroupMentionIDs []string `json:"group_mention_ids,omitempty"`
	// ID the unique identifier for the file.
	ID int64 `json:"id,omitempty"`
	// MemberMentionIDs the members that are mentioned in the description of the file.
	MemberMentionIDs []string `json:"member_mention_ids,omitempty"`
	// MentionIDs deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids,omitempty"`
	// Name the name of the linked file.
	Name string `json:"name,omitempty"`
	// Size the filesize, if the integration provided it.
	Size *int64 `json:"size,omitempty"`
	// StoryIDs the IDs of the stories this file is attached to.
	StoryIDs []int64 `json:"story_ids,omitempty"`
	// ThumbnailURL the URL of the file thumbnail, if the integration provided it.
	ThumbnailURL *string `json:"thumbnail_url,omitempty"`
	// Type the integration type (e.g. google, dropbox, box).
	Type string `json:"type,omitempty"`
	// UpdatedAt the time/date the LinkedFile was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UploaderID the UUID of the member that uploaded the file.
	UploaderID string `json:"uploader_id,omitempty"`
	// URL the URL of the file.
	URL string `json:"url,omitempty"`
}

func (m *LinkedFile) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *LinkedFile) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
