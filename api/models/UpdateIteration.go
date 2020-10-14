package models

import "encoding/json"

type UpdateIteration struct {
	// Description the description of the Iteration.
	Description string `json:"description"`
	// EndDate the date this Iteration ends, e.g. 2019-07-05.
	EndDate string `json:"end_date"`
	// FollowerIDs an array of UUIDs for any Members you want to add as Followers.
	FollowerIDs []string `json:"follower_ids"`
	// GroupIDs an array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	GroupIDs []string `json:"group_ids"`
	// Labels an array of Labels attached to the Iteration.
	Labels []CreateLabelParams `json:"labels"`
	// Name the name of this Iteration
	Name string `json:"name"`
	// StartDate the date this Iteration begins, e.g. 2019-07-01
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
