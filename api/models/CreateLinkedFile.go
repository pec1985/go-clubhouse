package models

import "encoding/json"

type CreateLinkedFile struct {
	// ContentType the content type of the image (e.g. txt/plain).
	ContentType string `json:"content_type,omitempty"`
	// Description the description of the file.
	Description string `json:"description,omitempty"`
	// Name the name of the file.
	Name string `json:"name,omitempty"`
	// Size the filesize, if the integration provided it.
	Size int64 `json:"size,omitempty"`
	// StoryID the ID of the linked story.
	StoryID int64 `json:"story_id,omitempty"`
	// ThumbnailURL the URL of the thumbnail, if the integration provided it.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// Type the integration type of the file (e.g. google, dropbox, box).
	Type string `json:"type,omitempty"`
	// UploaderID the UUID of the member that uploaded the file.
	UploaderID string `json:"uploader_id,omitempty"`
	// URL the URL of linked file.
	URL string `json:"url,omitempty"`
}

func (m *CreateLinkedFile) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateLinkedFile) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
