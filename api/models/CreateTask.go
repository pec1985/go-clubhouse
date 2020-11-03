package models

import "time"

type CreateTask struct {
	// Complete true/false boolean indicating whether the Task is completed. Defaults to false.
	Complete bool `json:"complete,omitempty"`
	// CreatedAt defaults to the time/date the Task is created but can be set to reflect another creation time/date.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Description the Task description.
	Description string `json:"description,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`
	// OwnerIDs an array of UUIDs for any members you want to add as Owners on this new Task.
	OwnerIDs []string `json:"owner_ids,omitempty"`
	// UpdatedAt defaults to the time/date the Task is created in Clubhouse but can be set to reflect another time/date.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *CreateTask) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateTask) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
