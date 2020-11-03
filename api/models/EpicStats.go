package models

import "time"

// EpicStats a group of calculated values for this Epic.
type EpicStats struct {
	// AverageCycleTime the average cycle time (in seconds) of completed stories in this Epic.
	AverageCycleTime int64 `json:"average_cycle_time,omitempty"`
	// AverageLeadTime the average lead time (in seconds) of completed stories in this Epic.
	AverageLeadTime int64 `json:"average_lead_time,omitempty"`
	// LastStoryUpdate the date of the last update of a Story in this Epic.
	LastStoryUpdate *time.Time `json:"last_story_update,omitempty"`
	// NumPoints the total number of points in this Epic.
	NumPoints int64 `json:"num_points,omitempty"`
	// NumPointsDone the total number of completed points in this Epic.
	NumPointsDone int64 `json:"num_points_done,omitempty"`
	// NumPointsStarted the total number of started points in this Epic.
	NumPointsStarted int64 `json:"num_points_started,omitempty"`
	// NumPointsUnstarted the total number of unstarted points in this Epic.
	NumPointsUnstarted int64 `json:"num_points_unstarted,omitempty"`
	// NumRelatedDocuments the total number of documents associated with this Epic.
	NumRelatedDocuments int64 `json:"num_related_documents,omitempty"`
	// NumStoriesDone the total number of done Stories in this Epic.
	NumStoriesDone int64 `json:"num_stories_done,omitempty"`
	// NumStoriesStarted the total number of started Stories in this Epic.
	NumStoriesStarted int64 `json:"num_stories_started,omitempty"`
	// NumStoriesUnestimated the total number of Stories with no point estimate.
	NumStoriesUnestimated int64 `json:"num_stories_unestimated,omitempty"`
	// NumStoriesUnstarted the total number of unstarted Stories in this Epic.
	NumStoriesUnstarted int64 `json:"num_stories_unstarted,omitempty"`
}

func (m *EpicStats) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *EpicStats) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
