package api

import (
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/v1/api/models"
)

// List Teams returns a list of all Teams in the organization.
func (a *api) ListTeams() (*[]models.Team, error) {
	params := url.Values{}
	var out []models.Team
	if err := a.request("GET", "/api/v3/teams", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get Team is used to get Team information.
func (a *api) GetTeam(teamPublicId int64) (*models.Team, error) {
	params := url.Values{}
	var out models.Team
	if err := a.request("GET", "/api/v3/teams/"+fmt.Sprint(teamPublicId)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
