package api

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List Repositories returns a list of all Repositories and their attributes.
func (a *api) ListRepositories() (*[]models.Repository, error) {
	var body *bytes.Buffer
	params := url.Values{}
	var out []models.Repository
	if err := a.request("GET", "/api/v3/repositories", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get Repository returns information about the selected Repository.
func (a *api) GetRepository(repoPublicId int64) (*models.Repository, error) {
	var body *bytes.Buffer
	params := url.Values{}
	var out models.Repository
	if err := a.request("GET", "/api/v3/repositories/"+fmt.Sprint(repoPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
