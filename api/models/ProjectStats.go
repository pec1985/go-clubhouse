package models

import "encoding/json"

type ProjectStats struct {
	// The total number of points in this Project.
	NumPoints int64 `json:"num_points"`
	// The total number of documents related to this Project
	NumRelatedDocuments int64 `json:"num_related_documents"`
	// The total number of stories in this Project.
	NumStories int64 `json:"num_stories"`
}

func (m *ProjectStats) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *ProjectStats) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
