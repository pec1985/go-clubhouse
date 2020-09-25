package models

import "encoding/json"

type MemberInfo struct {
	Id          string             `json:"id"`
	MentionName string             `json:"mention_name"`
	Name        string             `json:"name"`
	Workspace2  BasicWorkspaceInfo `json:"workspace2"`
}

func (m *MemberInfo) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *MemberInfo) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
