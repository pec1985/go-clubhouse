package models

import (
	"encoding/json"
	"time"
)

// Workflow Details of the workflow associated with the Team.
type Workflow struct {
	// AutoAssignOwner Indicates if an owner is automatically assigned when an unowned story is started.
	AutoAssignOwner bool `json:"auto_assign_owner"`
	// CreatedAt The date the Workflow was created.
	CreatedAt time.Time `json:"created_at"`
	// DefaultStateID The unique ID of the default state that new Stories are entered into.
	DefaultStateID int64 `json:"default_state_id"`
	// Description A description of the workflow.
	Description string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ID The unique ID of the Workflow.
	ID int64 `json:"id"`
	// Name The name of the workflow.
	Name string `json:"name"`
	// ProjectIDs An array of IDs of projects within the Workflow.
	ProjectIDs []float64 `json:"project_ids"`
	// States A map of the states in this Workflow.
	States []WorkflowState `json:"states"`
	// TeamID The ID of the team the workflow belongs to.
	TeamID int64 `json:"team_id"`
	// UpdatedAt The date the Workflow was updated.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Workflow) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Workflow) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
