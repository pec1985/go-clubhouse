package models

import "encoding/json"

type CreateEntityTemplateExternalTicket struct {
	// The id of the ticket in the external system.
	ExternalId string `json:"external_id"`
	// The url for the ticket in the external system.
	ExternalUrl string `json:"external_url"`
}

func (m *CreateEntityTemplateExternalTicket) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateEntityTemplateExternalTicket) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
