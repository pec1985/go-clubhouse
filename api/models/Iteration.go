package models

import (
	"encoding/json"
	"time"
)

type Iteration struct {
	// The date this iteration begins.
	EndDate time.Time `json:"end_date"`
	// A string description of this resource
	EntityType string `json:"entity_type"`
	// An array of Group IDs that have been mentioned in the Story description.
	GroupMentionIds []string `json:"group_mention_ids"`
	// Deprecated: use member_mention_ids.
	MentionIds []string       `json:"mention_ids"`
	Stats      IterationStats `json:"stats"`
	// The Clubhouse application url for the Iteration.
	AppUrl string `json:"app_url"`
	// An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	GroupIds []string `json:"group_ids"`
	// An array of Member IDs that have been mentioned in the Story description.
	MemberMentionIds []string `json:"member_mention_ids"`
	// The instant when this iteration was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The instant when this iteration was created.
	CreatedAt time.Time `json:"created_at"`
	// The name of the iteration.
	Name string `json:"name"`
	// The date this iteration begins.
	StartDate time.Time `json:"start_date"`
	// The description of the iteration.
	Description string `json:"description"`
	// The ID of the iteration.
	Id int64 `json:"id"`
	// An array of labels attached to the iteration.
	Labels []Label `json:"labels"`
	// The status of the iteration. Values are either "unstarted", "started", or "done".
	Status string `json:"status"`
	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []string `json:"follower_ids"`
}

func (m *Iteration) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Iteration) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
