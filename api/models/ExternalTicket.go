package models

import "encoding/json"

// ExternalTicket a External Ticket allows you to create a link between an external system, like Zendesk, and a Clubhouse story.
type ExternalTicket struct {
	// ExternalID the ID used in the external system.
	ExternalID string `json:"external_id,omitempty"`
	// ExternalURL the full qualified url of the external ticket.
	ExternalURL string `json:"external_url,omitempty"`
	// ID a unique ID internal to Clubhouse.
	ID string `json:"id,omitempty"`
	// StoryIDs the Clubhouse Story ids associated with this External Ticket.
	StoryIDs []float64 `json:"story_ids,omitempty"`
}

func (m *ExternalTicket) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *ExternalTicket) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
