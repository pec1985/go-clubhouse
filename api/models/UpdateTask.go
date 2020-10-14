package models

import "encoding/json"

type UpdateTask struct {
	// AfterID move task after this task ID.
	AfterID int64 `json:"after_id"`
	// BeforeID move task before this task ID.
	BeforeID int64 `json:"before_id"`
	// Complete a true/false boolean indicating whether the task is complete.
	Complete bool `json:"complete"`
	// Description the Task's description.
	Description string `json:"description"`
	// OwnerIDs an array of UUIDs of the owners of this story.
	OwnerIDs []string `json:"owner_ids"`
}

func (m *UpdateTask) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateTask) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
