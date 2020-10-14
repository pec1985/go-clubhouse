package models

import (
	"encoding/json"
	"time"
)

type UpdateFile struct {
	// CreatedAt the time/date that the file was uploaded.
	CreatedAt time.Time `json:"created_at"`
	// Description the description of the file.
	Description string `json:"description"`
	// ExternalID an additional ID that you may wish to assign to the file.
	ExternalID string `json:"external_id"`
	// Name the name of the file.
	Name string `json:"name"`
	// UpdatedAt the time/date that the file was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UploaderID the unique ID assigned to the Member who uploaded the file to Clubhouse.
	UploaderID string `json:"uploader_id"`
}

func (m *UpdateFile) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateFile) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
