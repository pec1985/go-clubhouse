package models

import "encoding/json"

type MemberInfo struct {
	ID          string             `json:"id,omitempty"`
	MentionName string             `json:"mention_name,omitempty"`
	Name        string             `json:"name,omitempty"`
	Workspace2  BasicWorkspaceInfo `json:"workspace2,omitempty"`
}

func (m *MemberInfo) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *MemberInfo) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
