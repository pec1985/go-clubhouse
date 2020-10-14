package models

import (
	"encoding/json"
	"time"
)

// StoryLink story links allow you create semantic relationships between two stories. Relationship types are relates to, blocks / blocked by, and duplicates / is duplicated by. The format is `subject -> link -> object`, or for example "story 5 blocks story 6".
type StoryLink struct {
	// CreatedAt the time/date when the Story Link was created.
	CreatedAt time.Time `json:"created_at"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ID the unique identifier of the Story Link.
	ID int64 `json:"id"`
	// ObjectID the ID of the object Story.
	ObjectID int64 `json:"object_id"`
	// SubjectID the ID of the subject Story.
	SubjectID int64 `json:"subject_id"`
	// UpdatedAt the time/date when the Story Link was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// Verb how the subject Story acts on the object Story. This can be "blocks", "duplicates", or "relates to".
	Verb string `json:"verb"`
}

func (m *StoryLink) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *StoryLink) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
