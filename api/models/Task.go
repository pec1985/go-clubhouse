package models

import (
	"encoding/json"
	"time"
)

type Task struct {
	// Complete True/false boolean indicating whether the Task has been completed.
	Complete bool `json:"complete"`
	// CompletedAt The time/date the Task was completed.
	CompletedAt *time.Time `json:"completed_at"`
	// CreatedAt The time/date the Task was created.
	CreatedAt time.Time `json:"created_at"`
	// Description Full text of the Task.
	Description string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ExternalID This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id"`
	// GroupMentionIDs An array of UUIDs of Groups mentioned in this Task.
	GroupMentionIDs []string `json:"group_mention_ids"`
	// ID The unique ID of the Task.
	ID int64 `json:"id"`
	// MemberMentionIDs An array of UUIDs of Members mentioned in this Task.
	MemberMentionIDs []string `json:"member_mention_ids"`
	// MentionIDs Deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids"`
	// OwnerIDs An array of UUIDs of the Owners of this Task.
	OwnerIDs []string `json:"owner_ids"`
	// Position The number corresponding to the Task's position within a list of Tasks on a Story.
	Position int64 `json:"position"`
	// StoryID The unique identifier of the parent Story.
	StoryID int64 `json:"story_id"`
	// UpdatedAt The time/date the Task was updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

func (m *Task) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Task) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
