package models

import "encoding/json"

type StoryContentsTask struct {
	// This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id"`
	// An array of UUIDs of the Owners of this Task.
	OwnerIds []string `json:"owner_ids"`
	// The number corresponding to the Task's position within a list of Tasks on a Story.
	Position int64 `json:"position"`
	// True/false boolean indicating whether the Task has been completed.
	Complete bool `json:"complete"`
	// Full text of the Task.
	Description string `json:"description"`
}

func (m *StoryContentsTask) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *StoryContentsTask) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
