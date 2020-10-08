package api

import (
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// Get Epic Workflow returns the Epic Workflow for the organization.
func (a *api) GetEpicWorkflow() (*models.EpicWorkflow, error) {
	params := url.Values{}
	var out models.EpicWorkflow
	if err := a.Request("GET", "/api/v3/epic-workflow", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
