package models

// CreateStoryLinkParams request parameters for creating a Story Link within a Story.
type CreateStoryLinkParams struct {
	// ObjectID the unique ID of the Story defined as object.
	ObjectID int64 `json:"object_id,omitempty"`
	// SubjectID the unique ID of the Story defined as subject.
	SubjectID int64 `json:"subject_id,omitempty"`
	// Verb how the subject Story acts on the object Story. This can be "blocks", "duplicates", or "relates to".
	Verb string `json:"verb,omitempty"`
}

func (m *CreateStoryLinkParams) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateStoryLinkParams) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
