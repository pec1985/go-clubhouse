package api

import (
	"bytes"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List Workflows returns a list of all Workflows in the organization.
func (a *api) ListWorkflows() (*[]models.Workflow, error) {
	var body *bytes.Buffer
	params := url.Values{}
	var out []models.Workflow
	if err := a.request("GET", "/api/v3/workflows", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
