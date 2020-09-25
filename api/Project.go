package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List Projects returns a list of all Projects and their attributes.
func (a *api) ListProjects() (*[]models.Project, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out []models.Project
	if err := a.request("GET", "/api/v3/projects", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create Project is used to create a new Clubhouse Project.
func (a *api) CreateProject(createProject *models.CreateProject) error {
	params := url.Values{}
	jsonbody, _ := json.Marshal(createProject)
	body := bytes.NewBuffer(jsonbody)
	var out interface{}
	if err := a.request("POST", "/api/v3/projects", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Delete Project can be used to delete a Project. Projects can only be deleted if all associated Stories are moved or deleted. In the case that the Project cannot be deleted, you will receive a 422 response.
func (a *api) DeleteProject(projectPublicId int64) error {
	params := url.Values{}
	body := bytes.Buffer{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/projects/"+fmt.Sprint(projectPublicId)+"", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Get Project returns information about the selected Project.
func (a *api) GetProject(projectPublicId int64) (*models.Project, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out models.Project
	if err := a.request("GET", "/api/v3/projects/"+fmt.Sprint(projectPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update Project can be used to change properties of a Project.
func (a *api) UpdateProject(projectPublicId int64, updateProject *models.UpdateProject) (*models.Project, error) {
	params := url.Values{}
	jsonbody, _ := json.Marshal(updateProject)
	body := bytes.NewBuffer(jsonbody)
	var out models.Project
	if err := a.request("PUT", "/api/v3/projects/"+fmt.Sprint(projectPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
