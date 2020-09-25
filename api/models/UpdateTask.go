package models

import "encoding/json"

type UpdateTask struct {
	// The Task's description.
	Description string `json:"description"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids"`
	// Move task after this task ID.
	AfterId int64 `json:"after_id"`
	// Move task before this task ID.
	BeforeId int64 `json:"before_id"`
	// A true/false boolean indicating whether the task is complete.
	Complete bool `json:"complete"`
}

func (m *UpdateTask) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateTask) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
