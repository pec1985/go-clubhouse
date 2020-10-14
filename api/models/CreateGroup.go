package models

import "encoding/json"

type CreateGroup struct {
	// Description The description of the Group.
	Description string `json:"description"`
	// DisplayIconID The Icon id for the avatar of this Group.
	DisplayIconID string `json:"display_icon_id"`
	// MemberIDs The Member ids to add to this Group.
	MemberIDs []string `json:"member_ids"`
	// MentionName The mention name of this Group.
	MentionName string `json:"mention_name"`
	// Name The name of this Group.
	Name string `json:"name"`
}

func (m *CreateGroup) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateGroup) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
