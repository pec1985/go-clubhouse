package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListMembers returns information about members of the organization.
func (a *api) ListMembers(listMembers *models.ListMembers) (*[]models.Member, error) {
	params := url.Values{}
	if listMembers != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(listMembers)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out []models.Member
	if err := a.Request("GET", "/api/v3/members", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetMember returns information about a Member.
func (a *api) GetMember(memberID string, getMember *models.GetMember) (*models.Member, error) {
	params := url.Values{}
	if getMember != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(getMember)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out models.Member
	if err := a.Request("GET", "/api/v3/members/"+memberID+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
