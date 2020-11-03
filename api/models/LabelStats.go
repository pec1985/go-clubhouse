package models

// LabelStats a group of calculated values for this Label.
type LabelStats struct {
	// NumEpics the total number of Epics with this Label.
	NumEpics int64 `json:"num_epics,omitempty"`
	// NumPointsCompleted the total number of completed points with this Label.
	NumPointsCompleted int64 `json:"num_points_completed,omitempty"`
	// NumPointsInProgress the total number of in-progress points with this Label.
	NumPointsInProgress int64 `json:"num_points_in_progress,omitempty"`
	// NumPointsTotal the total number of points with this Label.
	NumPointsTotal int64 `json:"num_points_total,omitempty"`
	// NumRelatedDocuments the total number of Documents associated this Label.
	NumRelatedDocuments int64 `json:"num_related_documents,omitempty"`
	// NumStoriesCompleted the total number of completed Stories with this Label.
	NumStoriesCompleted int64 `json:"num_stories_completed,omitempty"`
	// NumStoriesInProgress the total number of in-progress Stories with this Label.
	NumStoriesInProgress int64 `json:"num_stories_in_progress,omitempty"`
	// NumStoriesTotal the total number of Stories with this Label.
	NumStoriesTotal int64 `json:"num_stories_total,omitempty"`
	// NumStoriesUnestimated the total number of Stories with no point estimate with this Label.
	NumStoriesUnestimated int64 `json:"num_stories_unestimated,omitempty"`
}

func (m *LabelStats) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *LabelStats) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
