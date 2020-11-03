package models

// Group a Group.
type Group struct {
	// AppURL the Clubhouse application url for the Group.
	AppURL string `json:"app_url,omitempty"`
	// Archived whether or not the Group is archived.
	Archived bool `json:"archived,omitempty"`
	// Description the description of the Group.
	Description string `json:"description,omitempty"`
	DisplayIcon *Icon  `json:"display_icon,omitempty"`
	// EntityType a string description of this resource.
	EntityType string `json:"entity_type,omitempty"`
	// ID the id of the Group.
	ID string `json:"id,omitempty"`
	// MemberIDs the Member IDs contain within the Group.
	MemberIDs []string `json:"member_ids,omitempty"`
	// MentionName the mention name of the Group.
	MentionName string `json:"mention_name,omitempty"`
	// Name the name of the Group.
	Name string `json:"name,omitempty"`
}

func (m *Group) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *Group) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
