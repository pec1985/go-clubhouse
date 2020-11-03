package models

import "time"

// Team group of Projects with the same Workflow.
type Team struct {
	// CreatedAt the time/date the Team was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Description the description of the Team.
	Description string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ID the unique identifier of the Team.
	ID int64 `json:"id,omitempty"`
	// Name the name of the Team.
	Name string `json:"name,omitempty"`
	// Position a number representing the position of the Team in relation to every other Team within the Organization.
	Position float64 `json:"position,omitempty"`
	// ProjectIDs an array of IDs of projects within the Team.
	ProjectIDs []float64 `json:"project_ids,omitempty"`
	// UpdatedAt the time/date the Team was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Workflow  Workflow  `json:"workflow,omitempty"`
}

func (m *Team) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Team) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
