package models

// MilestoneStats a group of calculated values for this Milestone.
type MilestoneStats struct {
	// AverageCycleTime the average cycle time (in seconds) of completed stories in this Milestone.
	AverageCycleTime int64 `json:"average_cycle_time,omitempty"`
	// AverageLeadTime the average lead time (in seconds) of completed stories in this Milestone.
	AverageLeadTime int64 `json:"average_lead_time,omitempty"`
	// NumRelatedDocuments the number of related documents tp this Milestone.
	NumRelatedDocuments int64 `json:"num_related_documents,omitempty"`
}

func (m *MilestoneStats) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *MilestoneStats) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
