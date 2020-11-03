package models

import (
	"encoding/json"
	"time"
)

// EntityTemplate an entity template can be used to prefill various fields when creating new stories.
type EntityTemplate struct {
	// AuthorID the unique ID of the member who created the template.
	AuthorID string `json:"author_id,omitempty"`
	// CreatedAt the time/date when the entity template was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ID the unique identifier for the entity template.
	ID string `json:"id,omitempty"`
	// LastUsedAt the last time that someone created an entity using this template.
	LastUsedAt time.Time `json:"last_used_at,omitempty"`
	// Name the template's name.
	Name          string        `json:"name,omitempty"`
	StoryContents StoryContents `json:"story_contents,omitempty"`
	// UpdatedAt the time/date when the entity template was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *EntityTemplate) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EntityTemplate) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
