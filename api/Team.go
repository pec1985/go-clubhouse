package api

import (
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListTeams returns a list of all Teams in the organization.
func (a *api) ListTeams() (*[]models.Team, error) {
	params := url.Values{}
	var out []models.Team
	if err := a.Request("GET", "/api/v3/teams", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTeam is used to get Team information.
func (a *api) GetTeam(teamID int64) (*models.Team, error) {
	params := url.Values{}
	var out models.Team
	if err := a.Request("GET", "/api/v3/teams/"+fmt.Sprint(teamID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
