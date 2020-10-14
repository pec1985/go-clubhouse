package models

import "encoding/json"

// IterationStats a group of calculated values for this Iteration.
type IterationStats struct {
	// AverageCycleTime the average cycle time (in seconds) of completed stories in this Iteration.
	AverageCycleTime int64 `json:"average_cycle_time"`
	// AverageLeadTime the average lead time (in seconds) of completed stories in this Iteration.
	AverageLeadTime int64 `json:"average_lead_time"`
	// NumPoints the total number of points in this Iteration.
	NumPoints int64 `json:"num_points"`
	// NumPointsDone the total number of completed points in this Iteration.
	NumPointsDone int64 `json:"num_points_done"`
	// NumPointsStarted the total number of started points in this Iteration.
	NumPointsStarted int64 `json:"num_points_started"`
	// NumPointsUnstarted the total number of unstarted points in this Iteration.
	NumPointsUnstarted int64 `json:"num_points_unstarted"`
	// NumRelatedDocuments the total number of documents related to an Iteration
	NumRelatedDocuments int64 `json:"num_related_documents"`
	// NumStoriesDone the total number of done Stories in this Iteration.
	NumStoriesDone int64 `json:"num_stories_done"`
	// NumStoriesStarted the total number of started Stories in this Iteration.
	NumStoriesStarted int64 `json:"num_stories_started"`
	// NumStoriesUnestimated the total number of Stories with no point estimate.
	NumStoriesUnestimated int64 `json:"num_stories_unestimated"`
	// NumStoriesUnstarted the total number of unstarted Stories in this Iteration.
	NumStoriesUnstarted int64 `json:"num_stories_unstarted"`
}

func (m *IterationStats) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *IterationStats) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
