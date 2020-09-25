package models

import (
	"encoding/json"
	"time"
)

type Team struct {
	// An array of IDs of projects within the Team.
	ProjectIds []float64 `json:"project_ids"`
	// The time/date the Team was created.
	CreatedAt time.Time `json:"created_at"`
	// The description of the Team.
	Description string `json:"description"`
	// The unique identifier of the Team.
	Id int64 `json:"id"`
	// The name of the Team.
	Name string `json:"name"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// A number representing the position of the Team in relation to every other Team within the Organization.
	Position float64 `json:"position"`
	// The time/date the Team was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	Workflow  Workflow  `json:"workflow"`
}

func (m *Team) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Team) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
