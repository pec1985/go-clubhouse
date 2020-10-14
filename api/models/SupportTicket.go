package models

import "encoding/json"

type SupportTicket struct {
	ExternalID  string    `json:"external_id"`
	ExternalURL string    `json:"external_url"`
	ID          string    `json:"id"`
	StoryIDs    []float64 `json:"story_ids"`
}

func (m *SupportTicket) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *SupportTicket) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
