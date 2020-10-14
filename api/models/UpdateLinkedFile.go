package models

import "encoding/json"

type UpdateLinkedFile struct {
	// Description the description of the file.
	Description string `json:"description"`
	// Name the name of the file.
	Name string `json:"name"`
	// Size the filesize, if the integration provided it.
	Size int64 `json:"size"`
	// StoryID the ID of the linked story.
	StoryID int64 `json:"story_id"`
	// ThumbnailURL the URL of the thumbnail, if the integration provided it.
	ThumbnailURL string `json:"thumbnail_url"`
	// Type the integration type of the file (e.g. google, dropbox, box).
	Type string `json:"type"`
	// UploaderID the UUID of the member that uploaded the file.
	UploaderID string `json:"uploader_id"`
	// URL the URL of linked file.
	URL string `json:"url"`
}

func (m *UpdateLinkedFile) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateLinkedFile) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
