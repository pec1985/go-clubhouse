package api

import (
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListRepositories returns a list of all Repositories and their attributes.
func (a *api) ListRepositories() (*[]models.Repository, error) {
	params := url.Values{}
	var out []models.Repository
	if err := a.Request("GET", "/api/v3/repositories", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRepository returns information about the selected Repository.
func (a *api) GetRepository(repoID int64) (*models.Repository, error) {
	params := url.Values{}
	var out models.Repository
	if err := a.Request("GET", "/api/v3/repositories/"+fmt.Sprint(repoID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
