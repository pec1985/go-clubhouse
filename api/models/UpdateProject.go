package models

import "encoding/json"

type UpdateProject struct {
	// The Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation string `json:"abbreviation"`
	// A true/false boolean indicating whether the Story is in archived state.
	Archived bool `json:"archived"`
	// The color that represents the Project in the UI.
	Color string `json:"color"`
	// The number of days before the thermometer appears in the Story summary.
	DaysToThermometer int64 `json:"days_to_thermometer"`
	// The Project's description.
	Description string `json:"description"`
	// An array of UUIDs for any Members you want to add as Followers.
	FollowerIds []string `json:"follower_ids"`
	// The Project's name.
	Name string `json:"name"`
	// Configuration to enable or disable thermometers in the Story summary.
	ShowThermometer bool `json:"show_thermometer"`
	// The ID of the team the project belongs to.
	TeamId int64 `json:"team_id"`
}

func (m *UpdateProject) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateProject) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
