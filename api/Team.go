package api

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List Teams returns a list of all Teams in the organization.
func (a *api) ListTeams() (*[]models.Team, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out []models.Team
	if err := a.request("GET", "/api/v3/teams", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get Team is used to get Team information.
func (a *api) GetTeam(teamPublicId int64) (*models.Team, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out models.Team
	if err := a.request("GET", "/api/v3/teams/"+fmt.Sprint(teamPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
