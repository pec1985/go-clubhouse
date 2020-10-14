package models

import "encoding/json"

type UpdateStoryLink struct {
	// ObjectID The ID of the object Story.
	ObjectID int64 `json:"object_id"`
	// SubjectID The ID of the subject Story.
	SubjectID int64 `json:"subject_id"`
	// Verb The type of link.
	Verb string `json:"verb"`
}

func (m *UpdateStoryLink) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateStoryLink) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
