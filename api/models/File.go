package models

import (
	"encoding/json"
	"time"
)

type File struct {
	// The time/date that the file was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// The unique IDs of the Members who are mentioned in the file description.
	MemberMentionIds []string `json:"member_mention_ids"`
	// Deprecated: use member_mention_ids.
	MentionIds []string `json:"mention_ids"`
	// The optional User-specified name of the file.
	Name string `json:"name"`
	// The size of the file.
	Size int64 `json:"size"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique IDs of the Groups who are mentioned in the file description.
	GroupMentionIds []string `json:"group_mention_ids"`
	// This field can be set to another unique ID. In the case that the File has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id"`
	// The unique ID for the file.
	Id int64 `json:"id"`
	// The unique IDs of the Stories associated with this file.
	StoryIds []int64 `json:"story_ids"`
	// The url where the thumbnail of the file can be found in Clubhouse.
	ThumbnailUrl *string `json:"thumbnail_url"`
	// The unique ID of the Member who uploaded the file.
	UploaderId string `json:"uploader_id"`
	// The URL for the file.
	Url *string `json:"url"`
	// Free form string corresponding to a text or image file.
	ContentType string `json:"content_type"`
	// The time/date that the file was created.
	CreatedAt time.Time `json:"created_at"`
	// The description of the file.
	Description *string `json:"description"`
	// The name assigned to the file in Clubhouse upon upload.
	Filename string `json:"filename"`
}

func (m *File) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *File) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
