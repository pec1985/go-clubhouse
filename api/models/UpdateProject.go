package models

import "encoding/json"

type UpdateProject struct {
	// Abbreviation the Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation string `json:"abbreviation"`
	// Archived a true/false boolean indicating whether the Story is in archived state.
	Archived bool `json:"archived"`
	// Color the color that represents the Project in the UI.
	Color string `json:"color"`
	// DaysToThermometer the number of days before the thermometer appears in the Story summary.
	DaysToThermometer int64 `json:"days_to_thermometer"`
	// Description the Project's description.
	Description string `json:"description"`
	// FollowerIDs an array of UUIDs for any Members you want to add as Followers.
	FollowerIDs []string `json:"follower_ids"`
	// Name the Project's name.
	Name string `json:"name"`
	// ShowThermometer configuration to enable or disable thermometers in the Story summary.
	ShowThermometer bool `json:"show_thermometer"`
	// TeamID the ID of the team the project belongs to.
	TeamID int64 `json:"team_id"`
}

func (m *UpdateProject) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateProject) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
