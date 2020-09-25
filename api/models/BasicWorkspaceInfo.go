package models

import "encoding/json"

type BasicWorkspaceInfo struct {
	EstimateScale []int64 `json:"estimate_scale"`
	UrlSlug       string  `json:"url_slug"`
}

func (m *BasicWorkspaceInfo) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *BasicWorkspaceInfo) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
