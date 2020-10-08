package api

import (
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// List Workflows returns a list of all Workflows in the organization.
func (a *api) ListWorkflows() (*[]models.Workflow, error) {
	params := url.Values{}
	var out []models.Workflow
	if err := a.Request("GET", "/api/v3/workflows", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
