package api

import (
	"bytes"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// Get Epic Workflow returns the Epic Workflow for the organization.
func (a *api) GetEpicWorkflow() (*models.EpicWorkflow, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out models.EpicWorkflow
	if err := a.request("GET", "/api/v3/epic-workflow", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
