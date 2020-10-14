package models

import "encoding/json"

// Group a Group.
type Group struct {
	// AppURL the Clubhouse application url for the Group.
	AppURL string `json:"app_url"`
	// Archived whether or not the Group is archived.
	Archived bool `json:"archived"`
	// Description the description of the Group.
	Description string `json:"description"`
	DisplayIcon *Icon  `json:"display_icon"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type"`
	// ID the id of the Group.
	ID string `json:"id"`
	// MemberIDs the Member IDs contain within the Group.
	MemberIDs []string `json:"member_ids"`
	// MentionName the mention name of the Group.
	MentionName string `json:"mention_name"`
	// Name the name of the Group.
	Name string `json:"name"`
}

func (m *Group) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Group) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
