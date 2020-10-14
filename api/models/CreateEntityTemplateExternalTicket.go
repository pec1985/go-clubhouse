package models

import "encoding/json"

type CreateEntityTemplateExternalTicket struct {
	// ExternalID the id of the ticket in the external system.
	ExternalID string `json:"external_id"`
	// ExternalURL the url for the ticket in the external system.
	ExternalURL string `json:"external_url"`
}

func (m *CreateEntityTemplateExternalTicket) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateEntityTemplateExternalTicket) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
