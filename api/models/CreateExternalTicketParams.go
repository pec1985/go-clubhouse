package models

import "encoding/json"

type CreateExternalTicketParams struct {
	// The id of the ticket in the external system.
	ExternalId string `json:"external_id"`
	// The url for the ticket in the external system.
	ExternalUrl string `json:"external_url"`
}

func (m *CreateExternalTicketParams) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateExternalTicketParams) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
