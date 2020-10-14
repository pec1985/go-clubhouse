package models

import (
	"encoding/json"
	"time"
)

// EpicStats A group of calculated values for this Epic.
type EpicStats struct {
	// AverageCycleTime The average cycle time (in seconds) of completed stories in this Epic.
	AverageCycleTime int64 `json:"average_cycle_time"`
	// AverageLeadTime The average lead time (in seconds) of completed stories in this Epic.
	AverageLeadTime int64 `json:"average_lead_time"`
	// LastStoryUpdate The date of the last update of a Story in this Epic.
	LastStoryUpdate *time.Time `json:"last_story_update"`
	// NumPoints The total number of points in this Epic.
	NumPoints int64 `json:"num_points"`
	// NumPointsDone The total number of completed points in this Epic.
	NumPointsDone int64 `json:"num_points_done"`
	// NumPointsStarted The total number of started points in this Epic.
	NumPointsStarted int64 `json:"num_points_started"`
	// NumPointsUnstarted The total number of unstarted points in this Epic.
	NumPointsUnstarted int64 `json:"num_points_unstarted"`
	// NumRelatedDocuments The total number of documents associated with this Epic.
	NumRelatedDocuments int64 `json:"num_related_documents"`
	// NumStoriesDone The total number of done Stories in this Epic.
	NumStoriesDone int64 `json:"num_stories_done"`
	// NumStoriesStarted The total number of started Stories in this Epic.
	NumStoriesStarted int64 `json:"num_stories_started"`
	// NumStoriesUnestimated The total number of Stories with no point estimate.
	NumStoriesUnestimated int64 `json:"num_stories_unestimated"`
	// NumStoriesUnstarted The total number of unstarted Stories in this Epic.
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
