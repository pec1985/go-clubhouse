package models

import (
	"encoding/json"
	"time"
)

type UpdateFile struct {
	// CreatedAt The time/date that the file was uploaded.
	CreatedAt time.Time `json:"created_at"`
	// Description The description of the file.
	Description string `json:"description"`
	// ExternalID An additional ID that you may wish to assign to the file.
	ExternalID string `json:"external_id"`
	// Name The name of the file.
	Name string `json:"name"`
	// UpdatedAt The time/date that the file was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UploaderID The unique ID assigned to the Member who uploaded the file to Clubhouse.
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
