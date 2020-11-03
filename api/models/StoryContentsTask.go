package models

import "encoding/json"

type StoryContentsTask struct {
	// Complete true/false boolean indicating whether the Task has been completed.
	Complete bool `json:"complete,omitempty"`
	// Description full text of the Task.
	Description string `json:"description,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id,omitempty"`
	// OwnerIDs an array of UUIDs of the Owners of this Task.
	OwnerIDs []string `json:"owner_ids,omitempty"`
	// Position the number corresponding to the Task's position within a list of Tasks on a Story.
	Position int64 `json:"position,omitempty"`
}

func (m *StoryContentsTask) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *StoryContentsTask) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
