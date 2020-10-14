package models

import "encoding/json"

// LabelStats A group of calculated values for this Label.
type LabelStats struct {
	// NumEpics The total number of Epics with this Label.
	NumEpics int64 `json:"num_epics"`
	// NumPointsCompleted The total number of completed points with this Label.
	NumPointsCompleted int64 `json:"num_points_completed"`
	// NumPointsInProgress The total number of in-progress points with this Label.
	NumPointsInProgress int64 `json:"num_points_in_progress"`
	// NumPointsTotal The total number of points with this Label.
	NumPointsTotal int64 `json:"num_points_total"`
	// NumRelatedDocuments The total number of Documents associated this Label.
	NumRelatedDocuments int64 `json:"num_related_documents"`
	// NumStoriesCompleted The total number of completed Stories with this Label.
	NumStoriesCompleted int64 `json:"num_stories_completed"`
	// NumStoriesInProgress The total number of in-progress Stories with this Label.
	NumStoriesInProgress int64 `json:"num_stories_in_progress"`
	// NumStoriesTotal The total number of Stories with this Label.
	NumStoriesTotal int64 `json:"num_stories_total"`
	// NumStoriesUnestimated The total number of Stories with no point estimate with this Label.
	NumStoriesUnestimated int64 `json:"num_stories_unestimated"`
}

func (m *LabelStats) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *LabelStats) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
