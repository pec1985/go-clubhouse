package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/v1/api/models"
)

// Delete Multiple Stories allows you to delete multiple archived stories at once.
func (a *api) DeleteMultipleStories(deleteStories *models.DeleteStories) error {
	params := url.Values{}
	if deleteStories != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(deleteStories)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/stories/bulk", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Create Multiple Stories allows you to create multiple stories in a single request using the same syntax as [Create Story](https://clubhouse.io/api/#create-story).
func (a *api) CreateMultipleStories(createStories *models.CreateStories) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createStories != nil {
		jsonbody, _ := json.Marshal(createStories)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/stories/bulk", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Update Multiple Stories allows you to make changes to numerous stories at once.
func (a *api) UpdateMultipleStories(updateStories *models.UpdateStories) (*[]models.StorySlim, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if updateStories != nil {
		jsonbody, _ := json.Marshal(updateStories)
		body = bytes.NewBuffer(jsonbody)
	}
	var out []models.StorySlim
	if err := a.request("PUT", "/api/v3/stories/bulk", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
