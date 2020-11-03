package models

// Reaction emoji reaction on a comment.
type Reaction struct {
	Emoji string `json:"emoji,omitempty"`
	// PermissionIDs permissions who have reacted with this.
	PermissionIDs []string `json:"permission_ids,omitempty"`
}

func (m *Reaction) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Reaction) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
