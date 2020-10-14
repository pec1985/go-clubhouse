package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// DeleteMultipleStories allows you to delete multiple archived stories at once.
func (a *api) DeleteMultipleStories(stories *models.DeleteStories) error {
	params := url.Values{}
	if stories != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(stories)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/stories/bulk", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// CreateMultipleStories allows you to create multiple stories in a single request using the same syntax as [Create Story](https://clubhouse.io/api/#create-story).
func (a *api) CreateMultipleStories(stories *models.CreateStories) error {
	params := url.Values{}
	var body *bytes.Buffer
	if stories != nil {
		jsonbody, _ := json.Marshal(stories)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/stories/bulk", params, body, &out); err != nil {
		return err
	}
	return nil
}

// UpdateMultipleStories allows you to make changes to numerous stories at once.
func (a *api) UpdateMultipleStories(stories *models.UpdateStories) (*[]models.StorySlim, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if stories != nil {
		jsonbody, _ := json.Marshal(stories)
		body = bytes.NewBuffer(jsonbody)
	}
	var out []models.StorySlim
	if err := a.Request("PUT", "/api/v3/stories/bulk", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
