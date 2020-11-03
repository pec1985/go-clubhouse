package models

type CreateStoryLink struct {
	// ObjectID the ID of the object Story.
	ObjectID int64 `json:"object_id,omitempty"`
	// SubjectID the ID of the subject Story.
	SubjectID int64 `json:"subject_id,omitempty"`
	// Verb the type of link.
	Verb string `json:"verb,omitempty"`
}

func (m *CreateStoryLink) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateStoryLink) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
