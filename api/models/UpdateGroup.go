package models

import "encoding/json"

type UpdateGroup struct {
	// Whether or not this Group is archived.
	Archived *bool `json:"archived"`
	// The description of this Group.
	Description string `json:"description"`
	// The Icon id for the avatar of this Group.
	DisplayIconId *string `json:"display_icon_id"`
	// The Member ids to add to this Group.
	MemberIds []string `json:"member_ids"`
	// The mention name of this Group.
	MentionName string `json:"mention_name"`
	// The name of this Group.
	Name string `json:"name"`
}

func (m *UpdateGroup) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateGroup) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
