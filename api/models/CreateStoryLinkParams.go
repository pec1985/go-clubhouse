package models

import "encoding/json"

type CreateStoryLinkParams struct {
	// The unique ID of the Story defined as object.
	ObjectId int64 `json:"object_id"`
	// The unique ID of the Story defined as subject.
	SubjectId int64 `json:"subject_id"`
	// How the subject Story acts on the object Story. This can be "blocks", "duplicates", or "relates to".
	Verb string `json:"verb"`
}

func (m *CreateStoryLinkParams) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateStoryLinkParams) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
