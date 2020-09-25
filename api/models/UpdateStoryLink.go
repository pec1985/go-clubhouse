package models

import "encoding/json"

type UpdateStoryLink struct {
	// The ID of the object Story.
	ObjectId int64 `json:"object_id"`
	// The ID of the subject Story.
	SubjectId int64 `json:"subject_id"`
	// The type of link.
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
