package models

import "encoding/json"

type SupportTicket struct {
	ExternalId  string    `json:"external_id"`
	ExternalUrl string    `json:"external_url"`
	Id          string    `json:"id"`
	StoryIds    []float64 `json:"story_ids"`
}

func (m *SupportTicket) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *SupportTicket) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
