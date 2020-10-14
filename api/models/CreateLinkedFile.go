package models

import "encoding/json"

type CreateLinkedFile struct {
	// ContentType The content type of the image (e.g. txt/plain).
	ContentType string `json:"content_type"`
	// Description The description of the file.
	Description string `json:"description"`
	// Name The name of the file.
	Name string `json:"name"`
	// Size The filesize, if the integration provided it.
	Size int64 `json:"size"`
	// StoryID The ID of the linked story.
	StoryID int64 `json:"story_id"`
	// ThumbnailURL The URL of the thumbnail, if the integration provided it.
	ThumbnailURL string `json:"thumbnail_url"`
	// Type The integration type of the file (e.g. google, dropbox, box).
	Type string `json:"type"`
	// UploaderID The UUID of the member that uploaded the file.
	UploaderID string `json:"uploader_id"`
	// URL The URL of linked file.
	URL string `json:"url"`
}

func (m *CreateLinkedFile) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateLinkedFile) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
