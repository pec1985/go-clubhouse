package models

import (
	"encoding/json"
	"time"
)

// EpicWorkflow epic Workflow is the array of defined Epic States. Epic Workflow can be queried using the API but must be updated in the Clubhouse UI.
type EpicWorkflow struct {
	// CreatedAt the date the Epic Workflow was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DefaultEpicStateID the unique ID of the default Epic State that new Epics are assigned by default.
	DefaultEpicStateID int64 `json:"default_epic_state_id,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// EpicStates a map of the Epic States in this Epic Workflow.
	EpicStates []EpicState `json:"epic_states,omitempty"`
	// ID the unique ID of the Epic Workflow.
	ID int64 `json:"id,omitempty"`
	// UpdatedAt the date the Epic Workflow was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *EpicWorkflow) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EpicWorkflow) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
