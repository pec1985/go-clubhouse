package models

import (
	"encoding/json"
	"time"
)

// EpicStats a group of calculated values for this Epic.
type EpicStats struct {
	// AverageCycleTime the average cycle time (in seconds) of completed stories in this Epic.
	AverageCycleTime int64 `json:"average_cycle_time"`
	// AverageLeadTime the average lead time (in seconds) of completed stories in this Epic.
	AverageLeadTime int64 `json:"average_lead_time"`
	// LastStoryUpdate the date of the last update of a Story in this Epic.
	LastStoryUpdate *time.Time `json:"last_story_update"`
	// NumPoints the total number of points in this Epic.
	NumPoints int64 `json:"num_points"`
	// NumPointsDone the total number of completed points in this Epic.
	NumPointsDone int64 `json:"num_points_done"`
	// NumPointsStarted the total number of started points in this Epic.
	NumPointsStarted int64 `json:"num_points_started"`
	// NumPointsUnstarted the total number of unstarted points in this Epic.
	NumPointsUnstarted int64 `json:"num_points_unstarted"`
	// NumRelatedDocuments the total number of documents associated with this Epic.
	NumRelatedDocuments int64 `json:"num_related_documents"`
	// NumStoriesDone the total number of done Stories in this Epic.
	NumStoriesDone int64 `json:"num_stories_done"`
	// NumStoriesStarted the total number of started Stories in this Epic.
	NumStoriesStarted int64 `json:"num_stories_started"`
	// NumStoriesUnestimated the total number of Stories with no point estimate.
	NumStoriesUnestimated int64 `json:"num_stories_unestimated"`
	// NumStoriesUnstarted the total number of unstarted Stories in this Epic.
	NumStoriesUnstarted int64 `json:"num_stories_unstarted"`
}

func (m *EpicStats) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *EpicStats) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
