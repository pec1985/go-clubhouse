package models

type UpdateTask struct {
	// AfterID move task after this task ID.
	AfterID int64 `json:"after_id,omitempty"`
	// BeforeID move task before this task ID.
	BeforeID int64 `json:"before_id,omitempty"`
	// Complete a true/false boolean indicating whether the task is complete.
	Complete bool `json:"complete,omitempty"`
	// Description the Task's description.
	Description string `json:"description,omitempty"`
	// OwnerIDs an array of UUIDs of the owners of this story.
	OwnerIDs []string `json:"owner_ids,omitempty"`
}

func (m *UpdateTask) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateTask) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
