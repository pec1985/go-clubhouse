package models

import (
	"encoding/json"
	"time"
)

type UpdateFile struct {
	// The time/date that the file was uploaded.
	CreatedAt time.Time `json:"created_at"`
	// The description of the file.
	Description string `json:"description"`
	// An additional ID that you may wish to assign to the file.
	ExternalId string `json:"external_id"`
	// The name of the file.
	Name string `json:"name"`
	// The time/date that the file was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The unique ID assigned to the Member who uploaded the file to Clubhouse.
	UploaderId string `json:"uploader_id"`
}

func (m *UpdateFile) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateFile) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
