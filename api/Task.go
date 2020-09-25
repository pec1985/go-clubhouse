package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// Update Task can be used to update Task properties.
func (a *api) UpdateTask(storyPublicId int64, taskPublicId int64, updateTask *models.UpdateTask) (*models.Task, error) {
	params := url.Values{}
	jsonbody, _ := json.Marshal(updateTask)
	body := bytes.NewBuffer(jsonbody)
	var out models.Task
	if err := a.request("PUT", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/tasks/"+fmt.Sprint(taskPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete Task can be used to delete any previously created Task on a Story.
func (a *api) DeleteTask(storyPublicId int64, taskPublicId int64) error {
	params := url.Values{}
	body := bytes.Buffer{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/tasks/"+fmt.Sprint(taskPublicId)+"", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Returns information about a chosen Task.
func (a *api) GetTask(storyPublicId int64, taskPublicId int64) (*models.Task, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out models.Task
	if err := a.request("GET", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/tasks/"+fmt.Sprint(taskPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create Task is used to create a new task in a Story.
func (a *api) CreateTask(storyPublicId int64, createTask *models.CreateTask) error {
	params := url.Values{}
	jsonbody, _ := json.Marshal(createTask)
	body := bytes.NewBuffer(jsonbody)
	var out interface{}
	if err := a.request("POST", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/tasks", params, body, &out); err != nil {
		return err
	}
	return nil
}
