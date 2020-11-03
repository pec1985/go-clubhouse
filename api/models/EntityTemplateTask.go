package models

import "encoding/json"

// EntityTemplateTask request parameters for specifying how to pre-populate a task through a template.
type EntityTemplateTask struct {
	// Complete true/false boolean indicating whether the Task is completed. Defaults to false.
	Complete bool `json:"complete,omitempty"`
	// Description the Task description.
	Description string `json:"description,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`
	// OwnerIDs an array of UUIDs for any members you want to add as Owners on this new Task.
	OwnerIDs []string `json:"owner_ids,omitempty"`
}

func (m *EntityTemplateTask) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EntityTemplateTask) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
