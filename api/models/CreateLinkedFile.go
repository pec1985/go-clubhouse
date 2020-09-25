package models

import "encoding/json"

type CreateLinkedFile struct {
	// The filesize, if the integration provided it.
	Size int64 `json:"size"`
	// The ID of the linked story.
	StoryId int64 `json:"story_id"`
	// The URL of the thumbnail, if the integration provided it.
	ThumbnailUrl string `json:"thumbnail_url"`
	// The integration type of the file (e.g. google, dropbox, box).
	Type string `json:"type"`
	// The content type of the image (e.g. txt/plain).
	ContentType string `json:"content_type"`
	// The description of the file.
	Description string `json:"description"`
	// The name of the file.
	Name string `json:"name"`
	// The UUID of the member that uploaded the file.
	UploaderId string `json:"uploader_id"`
	// The URL of linked file.
	Url string `json:"url"`
}

func (m *CreateLinkedFile) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateLinkedFile) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
