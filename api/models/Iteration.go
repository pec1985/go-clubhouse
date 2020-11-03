package models

import (
	"encoding/json"
	"time"
)

// Iteration an Iteration is a defined, time-boxed period of development for a collection of Stories. See https://help.clubhouse.io/hc/en-us/articles/360028953452-Iterations-Overview for more information.
type Iteration struct {
	// AppURL the Clubhouse application url for the Iteration.
	AppURL string `json:"app_url,omitempty"`
	// CreatedAt the instant when this iteration was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Description the description of the iteration.
	Description string `json:"description,omitempty"`
	// EndDate the date this iteration begins.
	EndDate string `json:"end_date,omitempty"`
	// EntityType a string description of this resource
	EntityType string `json:"entity_type,omitempty"`
	// FollowerIDs an array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	// GroupIDs an array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	GroupIDs []string `json:"group_ids,omitempty"`
	// GroupMentionIDs an array of Group IDs that have been mentioned in the Story description.
	GroupMentionIDs []string `json:"group_mention_ids,omitempty"`
	// ID the ID of the iteration.
	ID int64 `json:"id,omitempty"`
	// Labels an array of labels attached to the iteration.
	Labels []Label `json:"labels,omitempty"`
	// MemberMentionIDs an array of Member IDs that have been mentioned in the Story description.
	MemberMentionIDs []string `json:"member_mention_ids,omitempty"`
	// MentionIDs deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids,omitempty"`
	// Name the name of the iteration.
	Name string `json:"name,omitempty"`
	// StartDate the date this iteration begins.
	StartDate string         `json:"start_date,omitempty"`
	Stats     IterationStats `json:"stats,omitempty"`
	// Status the status of the iteration. Values are either "unstarted", "started", or "done".
	Status string `json:"status,omitempty"`
	// UpdatedAt the instant when this iteration was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *Iteration) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Iteration) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
