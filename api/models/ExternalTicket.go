package models

import "encoding/json"

type ExternalTicket struct {
	// The ID used in the external system.
	ExternalId string `json:"external_id"`
	// The full qualified url of the external ticket.
	ExternalUrl string `json:"external_url"`
	// A unique ID internal to Clubhouse.
	Id string `json:"id"`
	// The Clubhouse Story ids associated with this External Ticket.
	StoryIds []float64 `json:"story_ids"`
}

func (m *ExternalTicket) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *ExternalTicket) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
