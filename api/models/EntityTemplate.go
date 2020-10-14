package models

import (
	"encoding/json"
	"time"
)

// EntityTemplate An entity template can be used to prefill various fields when creating new stories.
type EntityTemplate struct {
	// AuthorID The unique ID of the member who created the template.
	AuthorID string `json:"author_id"`
	// CreatedAt The time/date when the entity template was created.
	CreatedAt time.Time `json:"created_at"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ID The unique identifier for the entity template.
	ID string `json:"id"`
	// LastUsedAt The last time that someone created an entity using this template.
	LastUsedAt time.Time `json:"last_used_at"`
	// Name The template's name.
	Name          string        `json:"name"`
	StoryContents StoryContents `json:"story_contents"`
	// UpdatedAt The time/date when the entity template was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *EntityTemplate) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EntityTemplate) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
