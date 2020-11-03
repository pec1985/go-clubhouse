package models

type BasicWorkspaceInfo struct {
	EstimateScale []int64 `json:"estimate_scale,omitempty"`
	UrlSlug       string  `json:"url_slug,omitempty"`
}

func (m *BasicWorkspaceInfo) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *BasicWorkspaceInfo) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
