package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// CreateTask is used to create a new task in a Story.
func (a *api) CreateTask(storyID int64, task *models.CreateTask) error {
	params := url.Values{}
	var body *bytes.Buffer
	if task != nil {
		jsonbody, _ := json.Marshal(task)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/stories/"+fmt.Sprint(storyID)+"/tasks", params, body, &out); err != nil {
		return err
	}
	return nil
}

// DeleteTask can be used to delete any previously created Task on a Story.
func (a *api) DeleteTask(storyID int64, taskID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/stories/"+fmt.Sprint(storyID)+"/tasks/"+fmt.Sprint(taskID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// GetTask returns information about a chosen Task.
func (a *api) GetTask(storyID int64, taskID int64) (*models.Task, error) {
	params := url.Values{}
	var out models.Task
	if err := a.Request("GET", "/api/v3/stories/"+fmt.Sprint(storyID)+"/tasks/"+fmt.Sprint(taskID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateTask can be used to update Task properties.
func (a *api) UpdateTask(storyID int64, taskID int64, task *models.UpdateTask) (*models.Task, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if task != nil {
		jsonbody, _ := json.Marshal(task)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Task
	if err := a.Request("PUT", "/api/v3/stories/"+fmt.Sprint(storyID)+"/tasks/"+fmt.Sprint(taskID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
