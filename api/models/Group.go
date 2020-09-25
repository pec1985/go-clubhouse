package models

import "encoding/json"

type Group struct {
	// The Clubhouse application url for the Group.
	AppUrl string `json:"app_url"`
	// Whether or not the Group is archived.
	Archived bool `json:"archived"`
	// The description of the Group.
	Description string `json:"description"`
	DisplayIcon *Icon  `json:"display_icon"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The id of the Group.
	Id string `json:"id"`
	// The Member IDs contain within the Group.
	MemberIds []string `json:"member_ids"`
	// The mention name of the Group.
	MentionName string `json:"mention_name"`
	// The name of the Group.
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
