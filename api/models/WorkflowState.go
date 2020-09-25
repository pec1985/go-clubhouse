package models

import (
	"encoding/json"
	"time"
)

type WorkflowState struct {
	// The hex color for this Workflow State.
	Color string `json:"color"`
	// The time/date the Workflow State was created.
	CreatedAt time.Time `json:"created_at"`
	// The description of what sort of Stories belong in that Workflow state.
	Description string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique ID of the Workflow State.
	Id int64 `json:"id"`
	// The Workflow State's name.
	Name string `json:"name"`
	// The number of Stories currently in that Workflow State.
	NumStories int64 `json:"num_stories"`
	// The number of Story Templates associated with that Workflow State.
	NumStoryTemplates int64 `json:"num_story_templates"`
	// The position that the Workflow State is in, starting with 0 at the left.
	Position int64 `json:"position"`
	// The type of Workflow State (Unstarted, Started, or Finished)
	Type string `json:"type"`
	// When the Workflow State was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The verb that triggers a move to that Workflow State when making VCS commits.
	Verb *string `json:"verb"`
}

func (m *WorkflowState) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *WorkflowState) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
