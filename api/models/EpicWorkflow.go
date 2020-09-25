package models

import (
	"encoding/json"
	"time"
)

type EpicWorkflow struct {
	// The date the Epic Workflow was created.
	CreatedAt time.Time `json:"created_at"`
	// The unique ID of the default Epic State that new Epics are assigned by default.
	DefaultEpicStateId int64 `json:"default_epic_state_id"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// A map of the Epic States in this Epic Workflow.
	EpicStates []EpicState `json:"epic_states"`
	// The unique ID of the Epic Workflow.
	Id int64 `json:"id"`
	// The date the Epic Workflow was updated.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *EpicWorkflow) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EpicWorkflow) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
