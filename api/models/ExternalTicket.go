package models

import "encoding/json"

// ExternalTicket A External Ticket allows you to create a link between an external system, like Zendesk, and a Clubhouse story.
type ExternalTicket struct {
	// ExternalID The ID used in the external system.
	ExternalID string `json:"external_id"`
	// ExternalURL The full qualified url of the external ticket.
	ExternalURL string `json:"external_url"`
	// ID A unique ID internal to Clubhouse.
	ID string `json:"id"`
	// StoryIDs The Clubhouse Story ids associated with this External Ticket.
	StoryIDs []float64 `json:"story_ids"`
}

func (m *ExternalTicket) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *ExternalTicket) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
