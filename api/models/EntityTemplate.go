package models

import (
	"encoding/json"
	"time"
)

type EntityTemplate struct {
	// The last time that someone created an entity using this template.
	LastUsedAt time.Time `json:"last_used_at"`
	// The template's name.
	Name          string        `json:"name"`
	StoryContents StoryContents `json:"story_contents"`
	// The time/date when the entity template was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The unique ID of the member who created the template.
	AuthorId string `json:"author_id"`
	// The time/date when the entity template was created.
	CreatedAt time.Time `json:"created_at"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique identifier for the entity template.
	Id string `json:"id"`
}

func (m *EntityTemplate) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EntityTemplate) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
