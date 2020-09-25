package models

import (
	"encoding/json"
	"time"
)

type Workflow struct {
	// The ID of the team the workflow belongs to.
	TeamId int64 `json:"team_id"`
	// Indicates if an owner is automatically assigned when an unowned story is started.
	AutoAssignOwner bool `json:"auto_assign_owner"`
	// The unique ID of the default state that new Stories are entered into.
	DefaultStateId int64 `json:"default_state_id"`
	// A description of the workflow.
	Description string `json:"description"`
	// The name of the workflow.
	Name string `json:"name"`
	// An array of IDs of projects within the Workflow.
	ProjectIds []float64 `json:"project_ids"`
	// A map of the states in this Workflow.
	States []WorkflowState `json:"states"`
	// The date the Workflow was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The date the Workflow was created.
	CreatedAt time.Time `json:"created_at"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique ID of the Workflow.
	Id int64 `json:"id"`
}

func (m *Workflow) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Workflow) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
