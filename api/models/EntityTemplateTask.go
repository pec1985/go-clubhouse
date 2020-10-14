package models

import "encoding/json"

// EntityTemplateTask Request parameters for specifying how to pre-populate a task through a template.
type EntityTemplateTask struct {
	// Complete True/false boolean indicating whether the Task is completed. Defaults to false.
	Complete bool `json:"complete"`
	// Description The Task description.
	Description string `json:"description"`
	// ExternalID This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id"`
	// OwnerIDs An array of UUIDs for any members you want to add as Owners on this new Task.
	OwnerIDs []string `json:"owner_ids"`
}

func (m *EntityTemplateTask) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EntityTemplateTask) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
