package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/v1/api/models"
)

// List Stories returns a list of all Stories in a selected Project and their attributes.
func (a *api) ListStories(projectPublicId int64, getProjectStories *models.GetProjectStories) (*[]models.StorySlim, error) {
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
	if err := a.request("GET", "/api/v3/projects/"+fmt.Sprint(projectPublicId)+"/stories", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create Story is used to add a new story to your Clubhouse.
func (a *api) CreateStory(createStoryParams *models.CreateStoryParams) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createStoryParams != nil {
		jsonbody, _ := json.Marshal(createStoryParams)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/stories", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Delete Story can be used to delete any Story.
func (a *api) DeleteStory(storyPublicId int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Get Story returns information about a chosen Story.
func (a *api) GetStory(storyPublicId int64) (*models.Story, error) {
	params := url.Values{}
	var out models.Story
	if err := a.request("GET", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update Story can be used to update Story properties.
func (a *api) UpdateStory(storyPublicId int64, updateStory *models.UpdateStory) (*models.Story, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if updateStory != nil {
		jsonbody, _ := json.Marshal(updateStory)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Story
	if err := a.request("PUT", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
