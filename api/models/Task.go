package models

import (
	"encoding/json"
	"time"
)

type Task struct {
	// Complete true/false boolean indicating whether the Task has been completed.
	Complete bool `json:"complete,omitempty"`
	// CompletedAt the time/date the Task was completed.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt the time/date the Task was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Description full text of the Task.
	Description string `json:"description,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ExternalID this field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID *string `json:"external_id,omitempty"`
	// GroupMentionIDs an array of UUIDs of Groups mentioned in this Task.
	GroupMentionIDs []string `json:"group_mention_ids,omitempty"`
	// ID the unique ID of the Task.
	ID int64 `json:"id,omitempty"`
	// MemberMentionIDs an array of UUIDs of Members mentioned in this Task.
	MemberMentionIDs []string `json:"member_mention_ids,omitempty"`
	// MentionIDs deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids,omitempty"`
	// OwnerIDs an array of UUIDs of the Owners of this Task.
	OwnerIDs []string `json:"owner_ids,omitempty"`
	// Position the number corresponding to the Task's position within a list of Tasks on a Story.
	Position int64 `json:"position,omitempty"`
	// StoryID the unique identifier of the parent Story.
	StoryID int64 `json:"story_id,omitempty"`
	// UpdatedAt the time/date the Task was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (m *Task) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Task) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
