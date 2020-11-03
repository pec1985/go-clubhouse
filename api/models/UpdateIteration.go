package models

type UpdateIteration struct {
	// Description the description of the Iteration.
	Description string `json:"description,omitempty"`
	// EndDate the date this Iteration ends, e.g. 2019-07-05.
	EndDate string `json:"end_date,omitempty"`
	// FollowerIDs an array of UUIDs for any Members you want to add as Followers.
	FollowerIDs []string `json:"follower_ids,omitempty"`
	// GroupIDs an array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	GroupIDs []string `json:"group_ids,omitempty"`
	// Labels an array of Labels attached to the Iteration.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// Name the name of this Iteration
	Name string `json:"name,omitempty"`
	// StartDate the date this Iteration begins, e.g. 2019-07-01
	StartDate string `json:"start_date,omitempty"`
}

func (m *UpdateIteration) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateIteration) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
