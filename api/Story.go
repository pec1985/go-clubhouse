package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// Delete Story can be used to delete any Story.
func (a *api) DeleteStory(storyPublicId int64) error {
	params := url.Values{}
	body := bytes.Buffer{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Get Story returns information about a chosen Story.
func (a *api) GetStory(storyPublicId int64) (*models.Story, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out models.Story
	if err := a.request("GET", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update Story can be used to update Story properties.
func (a *api) UpdateStory(storyPublicId int64, updateStory *models.UpdateStory) (*models.Story, error) {
	params := url.Values{}
	jsonbody, _ := json.Marshal(updateStory)
	body := bytes.NewBuffer(jsonbody)
	var out models.Story
	if err := a.request("PUT", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create Story is used to add a new story to your Clubhouse.
func (a *api) CreateStory(createStoryParams *models.CreateStoryParams) error {
	params := url.Values{}
	jsonbody, _ := json.Marshal(createStoryParams)
	body := bytes.NewBuffer(jsonbody)
	var out interface{}
	if err := a.request("POST", "/api/v3/stories", params, body, &out); err != nil {
		return err
	}
	return nil
}
