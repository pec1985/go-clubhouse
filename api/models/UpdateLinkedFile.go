package models

type UpdateLinkedFile struct {
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

func (m *UpdateLinkedFile) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateLinkedFile) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
