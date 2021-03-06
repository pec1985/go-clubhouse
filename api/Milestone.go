package api

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListMilestones returns a list of all Milestones and their attributes.
func (a *api) ListMilestones() (*[]models.Milestone, error) {
	params := url.Values{}
	var out []models.Milestone
	if err := a.Request("GET", "/api/v3/milestones", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateMilestone allows you to create a new Milestone in Clubhouse.
func (a *api) CreateMilestone(milestone *models.CreateMilestone) error {
	params := url.Values{}
	var body *bytes.Buffer
	if milestone != nil {
		jsonbody, _ := toPayload(milestone, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/milestones", params, body, &out); err != nil {
		return err
	}
	return nil
}

// DeleteMilestone can be used to delete any Milestone.
func (a *api) DeleteMilestone(milestoneID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/milestones/"+fmt.Sprint(milestoneID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// GetMilestone returns information about a chosen Milestone.
func (a *api) GetMilestone(milestoneID int64) (*models.Milestone, error) {
	params := url.Values{}
	var out models.Milestone
	if err := a.Request("GET", "/api/v3/milestones/"+fmt.Sprint(milestoneID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateMilestone can be used to update Milestone properties.
func (a *api) UpdateMilestone(milestoneID int64, milestone *models.UpdateMilestone) (*models.Milestone, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if milestone != nil {
		jsonbody, _ := toPayload(milestone, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Milestone
	if err := a.Request("PUT", "/api/v3/milestones/"+fmt.Sprint(milestoneID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
