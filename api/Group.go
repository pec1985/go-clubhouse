package api

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

func (a *api) ListGroups() (*[]models.Group, error) {
	params := url.Values{}
	var out []models.Group
	if err := a.Request("GET", "/api/v3/groups", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (a *api) CreateGroup(group *models.CreateGroup) error {
	params := url.Values{}
	var body *bytes.Buffer
	if group != nil {
		jsonbody, _ := json.Marshal(group)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/groups", params, body, &out); err != nil {
		return err
	}
	return nil
}
func (a *api) GetGroup(groupID string) (*models.Group, error) {
	params := url.Values{}
	var out models.Group
	if err := a.Request("GET", "/api/v3/groups/"+groupID+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (a *api) UpdateGroup(groupID string, group *models.UpdateGroup) (*models.Group, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if group != nil {
		jsonbody, _ := json.Marshal(group)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Group
	if err := a.Request("PUT", "/api/v3/groups/"+groupID+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
