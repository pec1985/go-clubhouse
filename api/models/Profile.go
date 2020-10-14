package models

import "encoding/json"

// Profile A group of Member profile details.
type Profile struct {
	// Deactivated A true/false boolean indicating whether the Member has been deactivated within Clubhouse.
	Deactivated bool  `json:"deactivated"`
	DisplayIcon *Icon `json:"display_icon"`
	// EmailAddress The primary email address of the Member with the Organization.
	EmailAddress *string `json:"email_address"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// GravatarHash This is the gravatar hash associated with email_address.
	GravatarHash *string `json:"gravatar_hash"`
	// ID The unique identifier of the profile.
	ID string `json:"id"`
	// MentionName The Member's username within the Organization.
	MentionName string `json:"mention_name"`
	// Name The Member's name within the Organization.
	Name *string `json:"name"`
	// TwoFactorAuthActivated If Two Factor Authentication is activated for this User.
	TwoFactorAuthActivated bool `json:"two_factor_auth_activated"`
}

func (m *Profile) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Profile) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
