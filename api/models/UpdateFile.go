package models

import "time"

type UpdateFile struct {
	// CreatedAt the time/date that the file was uploaded.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Description the description of the file.
	Description string `json:"description,omitempty"`
	// ExternalID an additional ID that you may wish to assign to the file.
	ExternalID string `json:"external_id,omitempty"`
	// Name the name of the file.
	Name string `json:"name,omitempty"`
	// UpdatedAt the time/date that the file was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UploaderID the unique ID assigned to the Member who uploaded the file to Clubhouse.
	UploaderID string `json:"uploader_id,omitempty"`
}

func (m *UpdateFile) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateFile) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
