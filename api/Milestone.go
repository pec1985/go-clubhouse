package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// List Milestones returns a list of all Milestones and their attributes.
func (a *api) ListMilestones() (*[]models.Milestone, error) {
	params := url.Values{}
	var out []models.Milestone
	if err := a.Request("GET", "/api/v3/milestones", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create Milestone allows you to create a new Milestone in Clubhouse.
func (a *api) CreateMilestone(createMilestone *models.CreateMilestone) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createMilestone != nil {
		jsonbody, _ := json.Marshal(createMilestone)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/milestones", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Delete Milestone can be used to delete any Milestone.
func (a *api) DeleteMilestone(milestonePublicId int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/milestones/"+fmt.Sprint(milestonePublicId)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Get Milestone returns information about a chosen Milestone.
func (a *api) GetMilestone(milestonePublicId int64) (*models.Milestone, error) {
	params := url.Values{}
	var out models.Milestone
	if err := a.Request("GET", "/api/v3/milestones/"+fmt.Sprint(milestonePublicId)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update Milestone can be used to update Milestone properties.
func (a *api) UpdateMilestone(milestonePublicId int64, updateMilestone *models.UpdateMilestone) (*models.Milestone, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if updateMilestone != nil {
		jsonbody, _ := json.Marshal(updateMilestone)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Milestone
	if err := a.Request("PUT", "/api/v3/milestones/"+fmt.Sprint(milestonePublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
