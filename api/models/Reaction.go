package models

import "encoding/json"

// Reaction emoji reaction on a comment.
type Reaction struct {
	Emoji string `json:"emoji"`
	// PermissionIDs permissions who have reacted with this.
	PermissionIDs []string `json:"permission_ids"`
}

func (m *Reaction) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Reaction) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
