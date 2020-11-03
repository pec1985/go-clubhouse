package models

import "encoding/json"

// IterationStats a group of calculated values for this Iteration.
type IterationStats struct {
	// AverageCycleTime the average cycle time (in seconds) of completed stories in this Iteration.
	AverageCycleTime int64 `json:"average_cycle_time,omitempty"`
	// AverageLeadTime the average lead time (in seconds) of completed stories in this Iteration.
	AverageLeadTime int64 `json:"average_lead_time,omitempty"`
	// NumPoints the total number of points in this Iteration.
	NumPoints int64 `json:"num_points,omitempty"`
	// NumPointsDone the total number of completed points in this Iteration.
	NumPointsDone int64 `json:"num_points_done,omitempty"`
	// NumPointsStarted the total number of started points in this Iteration.
	NumPointsStarted int64 `json:"num_points_started,omitempty"`
	// NumPointsUnstarted the total number of unstarted points in this Iteration.
	NumPointsUnstarted int64 `json:"num_points_unstarted,omitempty"`
	// NumRelatedDocuments the total number of documents related to an Iteration
	NumRelatedDocuments int64 `json:"num_related_documents,omitempty"`
	// NumStoriesDone the total number of done Stories in this Iteration.
	NumStoriesDone int64 `json:"num_stories_done,omitempty"`
	// NumStoriesStarted the total number of started Stories in this Iteration.
	NumStoriesStarted int64 `json:"num_stories_started,omitempty"`
	// NumStoriesUnestimated the total number of Stories with no point estimate.
	NumStoriesUnestimated int64 `json:"num_stories_unestimated,omitempty"`
	// NumStoriesUnstarted the total number of unstarted Stories in this Iteration.
	NumStoriesUnstarted int64 `json:"num_stories_unstarted,omitempty"`
}

func (m *IterationStats) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *IterationStats) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
