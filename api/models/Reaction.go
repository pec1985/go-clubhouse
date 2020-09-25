package models

import "encoding/json"

type Reaction struct {
	// Emoji text of the reaction.
	Emoji string `json:"emoji"`
	// Permissions who have reacted with this.
	PermissionIds []string `json:"permission_ids"`
}

func (m *Reaction) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Reaction) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
