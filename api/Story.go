package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListStories returns a list of all Stories in a selected Project and their attributes.
func (a *api) ListStories(projectID int64, getProjectStories *models.GetProjectStories) (*[]models.StorySlim, error) {
	params := url.Values{}
	if getProjectStories != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(getProjectStories)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out []models.StorySlim
	if err := a.Request("GET", "/api/v3/projects/"+fmt.Sprint(projectID)+"/stories", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateStory is used to add a new story to your Clubhouse.
func (a *api) CreateStory(storyParams *models.CreateStoryParams) error {
	params := url.Values{}
	var body *bytes.Buffer
	if storyParams != nil {
		jsonbody, _ := toPayload(storyParams, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/stories", params, body, &out); err != nil {
		return err
	}
	return nil
}

// DeleteStory can be used to delete any Story.
func (a *api) DeleteStory(storyID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/stories/"+fmt.Sprint(storyID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// GetStory returns information about a chosen Story.
func (a *api) GetStory(storyID int64) (*models.Story, error) {
	params := url.Values{}
	var out models.Story
	if err := a.Request("GET", "/api/v3/stories/"+fmt.Sprint(storyID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateStory can be used to update Story properties.
func (a *api) UpdateStory(storyID int64, story *models.UpdateStory) (*models.Story, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if story != nil {
		jsonbody, _ := toPayload(story, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Story
	if err := a.Request("PUT", "/api/v3/stories/"+fmt.Sprint(storyID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
