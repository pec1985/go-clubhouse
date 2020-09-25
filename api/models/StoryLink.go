package models

import (
	"encoding/json"
	"time"
)

type StoryLink struct {
	// The time/date when the Story Link was created.
	CreatedAt time.Time `json:"created_at"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique identifier of the Story Link.
	Id int64 `json:"id"`
	// The ID of the object Story.
	ObjectId int64 `json:"object_id"`
	// The ID of the subject Story.
	SubjectId int64 `json:"subject_id"`
	// The time/date when the Story Link was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// How the subject Story acts on the object Story. This can be "blocks", "duplicates", or "relates to".
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
