package models

type CreateEntityTemplateExternalTicket struct {
	// ExternalID the id of the ticket in the external system.
	ExternalID string `json:"external_id,omitempty"`
	// ExternalURL the url for the ticket in the external system.
	ExternalURL string `json:"external_url,omitempty"`
}

func (m *CreateEntityTemplateExternalTicket) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateEntityTemplateExternalTicket) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
