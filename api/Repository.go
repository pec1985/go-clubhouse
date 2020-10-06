package api

import (
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// List Repositories returns a list of all Repositories and their attributes.
func (a *api) ListRepositories() (*[]models.Repository, error) {
	params := url.Values{}
	var out []models.Repository
	if err := a.request("GET", "/api/v3/repositories", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get Repository returns information about the selected Repository.
func (a *api) GetRepository(repoPublicId int64) (*models.Repository, error) {
	params := url.Values{}
	var out models.Repository
	if err := a.request("GET", "/api/v3/repositories/"+fmt.Sprint(repoPublicId)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
