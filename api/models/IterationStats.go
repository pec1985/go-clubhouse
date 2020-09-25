package models

import "encoding/json"

type IterationStats struct {
	// The total number of unstarted Stories in this Iteration.
	NumStoriesUnstarted int64 `json:"num_stories_unstarted"`
	// The average cycle time (in seconds) of completed stories in this Iteration.
	AverageCycleTime int64 `json:"average_cycle_time"`
	// The average lead time (in seconds) of completed stories in this Iteration.
	AverageLeadTime int64 `json:"average_lead_time"`
	// The total number of points in this Iteration.
	NumPoints int64 `json:"num_points"`
	// The total number of unstarted points in this Iteration.
	NumPointsUnstarted int64 `json:"num_points_unstarted"`
	// The total number of documents related to an Iteration
	NumRelatedDocuments int64 `json:"num_related_documents"`
	// The total number of started Stories in this Iteration.
	NumStoriesStarted int64 `json:"num_stories_started"`
	// The total number of Stories with no point estimate.
	NumStoriesUnestimated int64 `json:"num_stories_unestimated"`
	// The total number of completed points in this Iteration.
	NumPointsDone int64 `json:"num_points_done"`
	// The total number of started points in this Iteration.
	NumPointsStarted int64 `json:"num_points_started"`
	// The total number of done Stories in this Iteration.
	NumStoriesDone int64 `json:"num_stories_done"`
}

func (m *IterationStats) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *IterationStats) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
