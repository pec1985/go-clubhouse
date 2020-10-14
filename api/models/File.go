package models

import (
	"encoding/json"
	"time"
)

// File a File is any document uploaded to your Clubhouse. Files attached from a third-party service can be accessed using the Linked Files endpoint.
type File struct {
	// ContentType free form string corresponding to a text or image file.
	ContentType string `json:"content_type"`
	// CreatedAt the time/date that the file was created.
	CreatedAt time.Time `json:"created_at"`
	// Description the description of the file.
	Description *string `json:"description"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID this field can be set to another unique ID. In the case that the File has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// Filename the name assigned to the file in Clubhouse upon upload.
	Filename string `json:"filename"`
	// GroupMentionIDs the unique IDs of the Groups who are mentioned in the file description.
	GroupMentionIDs []string `json:"group_mention_ids"`
	// ID the unique ID for the file.
	ID int64 `json:"id"`
	// MemberMentionIDs the unique IDs of the Members who are mentioned in the file description.
	MemberMentionIDs []string `json:"member_mention_ids"`
	// MentionIDs deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids"`
	// Name the optional User-specified name of the file.
	Name string `json:"name"`
	// Size the size of the file.
	Size int64 `json:"size"`
	// StoryIDs the unique IDs of the Stories associated with this file.
	StoryIDs []int64 `json:"story_ids"`
	// ThumbnailURL the url where the thumbnail of the file can be found in Clubhouse.
	ThumbnailURL *string `json:"thumbnail_url"`
	// UpdatedAt the time/date that the file was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// UploaderID the unique ID of the Member who uploaded the file.
	UploaderID string `json:"uploader_id"`
	// URL the URL for the file.
	URL *string `json:"url"`
}

func (m *File) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *File) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
