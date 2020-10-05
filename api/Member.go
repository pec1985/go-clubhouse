package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List Members returns information about members of the organization.
func (a *api) ListMembers(listMembers *models.ListMembers) (*[]models.Member, error) {
	var body *bytes.Buffer
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
	if err := a.request("GET", "/api/v3/members", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Returns information about a Member.
func (a *api) GetMember(memberPublicId string, getMember *models.GetMember) (*models.Member, error) {
	var body *bytes.Buffer
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
	if err := a.request("GET", "/api/v3/members/"+memberPublicId+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
