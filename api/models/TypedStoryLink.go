package models

import (
	"encoding/json"
	"time"
)

// TypedStoryLink the type of Story Link. The string can be subject or object.
type TypedStoryLink struct {
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
	// Type this indicates whether the Story is the subject or object in the Story Link.
	Type string `json:"type"`
	// UpdatedAt the time/date when the Story Link was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// Verb how the subject Story acts on the object Story. This can be "blocks", "duplicates", or "relates to".
	Verb string `json:"verb"`
}

func (m *TypedStoryLink) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *TypedStoryLink) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
