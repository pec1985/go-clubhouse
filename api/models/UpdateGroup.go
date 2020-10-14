package models

import "encoding/json"

type UpdateGroup struct {
	// Archived whether or not this Group is archived.
	Archived *bool `json:"archived"`
	// Description the description of this Group.
	Description string `json:"description"`
	// DisplayIconID the Icon id for the avatar of this Group.
	DisplayIconID *string `json:"display_icon_id"`
	// MemberIDs the Member ids to add to this Group.
	MemberIDs []string `json:"member_ids"`
	// MentionName the mention name of this Group.
	MentionName string `json:"mention_name"`
	// Name the name of this Group.
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
