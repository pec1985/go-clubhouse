package models

import "encoding/json"

// Group A Group.
type Group struct {
	// AppURL The Clubhouse application url for the Group.
	AppURL string `json:"app_url"`
	// Archived Whether or not the Group is archived.
	Archived bool `json:"archived"`
	// Description The description of the Group.
	Description string `json:"description"`
	DisplayIcon *Icon  `json:"display_icon"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ID The id of the Group.
	ID string `json:"id"`
	// MemberIDs The Member IDs contain within the Group.
	MemberIDs []string `json:"member_ids"`
	// MentionName The mention name of the Group.
	MentionName string `json:"mention_name"`
	// Name The name of the Group.
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
