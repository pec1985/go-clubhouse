package models

import (
	"encoding/json"
	"time"
)

// CreateTaskParams request parameters for creating a Task on a Story.
type CreateTaskParams struct {
	// Complete true/false boolean indicating whether the Task is completed. Defaults to false.
	Complete bool `json:"complete"`
	// CreatedAt defaults to the time/date the Task is created but can be set to reflect another creation time/date.
	CreatedAt time.Time `json:"created_at"`
	// Description the Task description.
	Description string `json:"description"`
	// ExternalID this field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id"`
	// OwnerIDs an array of UUIDs for any members you want to add as Owners on this new Task.
	OwnerIDs []string `json:"owner_ids"`
	// UpdatedAt defaults to the time/date the Task is created in Clubhouse but can be set to reflect another time/date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *CreateTaskParams) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateTaskParams) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
