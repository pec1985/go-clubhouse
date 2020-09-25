package api

import (
	"bytes"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List Repositories returns a list of all Repositories and their attributes.
func (a *api) ListRepositories() (*[]models.Repository, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out []models.Repository
	if err := a.request("GET", "/api/v3/repositories", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
