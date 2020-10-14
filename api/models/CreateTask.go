package models

import (
	"encoding/json"
	"time"
)

type CreateTask struct {
	// Complete True/false boolean indicating whether the Task is completed. Defaults to false.
	Complete bool `json:"complete"`
	// CreatedAt Defaults to the time/date the Task is created but can be set to reflect another creation time/date.
	CreatedAt time.Time `json:"created_at"`
	// Description The Task description.
	Description string `json:"description"`
	// ExternalID This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id"`
	// OwnerIDs An array of UUIDs for any members you want to add as Owners on this new Task.
	OwnerIDs []string `json:"owner_ids"`
	// UpdatedAt Defaults to the time/date the Task is created in Clubhouse but can be set to reflect another time/date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *CreateTask) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateTask) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
