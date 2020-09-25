package models

import "encoding/json"

type MilestoneStats struct {
	// The average cycle time (in seconds) of completed stories in this Milestone.
	AverageCycleTime int64 `json:"average_cycle_time"`
	// The average lead time (in seconds) of completed stories in this Milestone.
	AverageLeadTime int64 `json:"average_lead_time"`
	// The number of related documents tp this Milestone.
	NumRelatedDocuments int64 `json:"num_related_documents"`
}

func (m *MilestoneStats) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *MilestoneStats) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
