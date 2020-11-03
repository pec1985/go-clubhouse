package models

import "time"

// Workflow details of the workflow associated with the Team.
type Workflow struct {
	// AutoAssignOwner indicates if an owner is automatically assigned when an unowned story is started.
	AutoAssignOwner bool `json:"auto_assign_owner,omitempty"`
	// CreatedAt the date the Workflow was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DefaultStateID the unique ID of the default state that new Stories are entered into.
	DefaultStateID int64 `json:"default_state_id,omitempty"`
	// Description a description of the workflow.
	Description string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ID the unique ID of the Workflow.
	ID int64 `json:"id,omitempty"`
	// Name the name of the workflow.
	Name string `json:"name,omitempty"`
	// ProjectIDs an array of IDs of projects within the Workflow.
	ProjectIDs []float64 `json:"project_ids,omitempty"`
	// States a map of the states in this Workflow.
	States []WorkflowState `json:"states,omitempty"`
	// TeamID the ID of the team the workflow belongs to.
	TeamID int64 `json:"team_id,omitempty"`
	// UpdatedAt the date the Workflow was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *Workflow) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Workflow) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
