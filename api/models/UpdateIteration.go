package models

import "encoding/json"

type UpdateIteration struct {
	// The description of the Iteration.
	Description string `json:"description"`
	// The date this Iteration ends, e.g. 2019-07-05.
	EndDate string `json:"end_date"`
	// An array of UUIDs for any Members you want to add as Followers.
	FollowerIds []string `json:"follower_ids"`
	// An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	GroupIds []string `json:"group_ids"`
	// An array of Labels attached to the Iteration.
	Labels []CreateLabelParams `json:"labels"`
	// The name of this Iteration
	Name string `json:"name"`
	// The date this Iteration begins, e.g. 2019-07-01
	StartDate string `json:"start_date"`
}

func (m *UpdateIteration) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateIteration) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
