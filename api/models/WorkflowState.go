package models

import (
	"encoding/json"
	"time"
)

// WorkflowState workflow State is any of the at least 3 columns. Workflow States correspond to one of 3 types: Unstarted, Started, or Done.
type WorkflowState struct {
	// Color the hex color for this Workflow State.
	Color string `json:"color,omitempty"`
	// CreatedAt the time/date the Workflow State was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Description the description of what sort of Stories belong in that Workflow state.
	Description string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ID the unique ID of the Workflow State.
	ID int64 `json:"id,omitempty"`
	// Name the Workflow State's name.
	Name string `json:"name,omitempty"`
	// NumStories the number of Stories currently in that Workflow State.
	NumStories int64 `json:"num_stories,omitempty"`
	// NumStoryTemplates the number of Story Templates associated with that Workflow State.
	NumStoryTemplates int64 `json:"num_story_templates,omitempty"`
	// Position the position that the Workflow State is in, starting with 0 at the left.
	Position int64 `json:"position,omitempty"`
	// Type the type of Workflow State (Unstarted, Started, or Finished)
	Type string `json:"type,omitempty"`
	// UpdatedAt when the Workflow State was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Verb the verb that triggers a move to that Workflow State when making VCS commits.
	Verb *string `json:"verb,omitempty"`
}

func (m *WorkflowState) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *WorkflowState) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
