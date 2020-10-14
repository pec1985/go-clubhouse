package models

import "encoding/json"

type StoryContentsTask struct {
	// Complete True/false boolean indicating whether the Task has been completed.
	Complete bool `json:"complete"`
	// Description Full text of the Task.
	Description string `json:"description"`
	// ExternalID This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// OwnerIDs An array of UUIDs of the Owners of this Task.
	OwnerIDs []string `json:"owner_ids"`
	// Position The number corresponding to the Task's position within a list of Tasks on a Story.
	Position int64 `json:"position"`
}

func (m *StoryContentsTask) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *StoryContentsTask) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
