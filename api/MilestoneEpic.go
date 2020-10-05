package api

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List all of the Epics within the Milestone.
func (a *api) ListMilestoneEpics(milestonePublicId int64) (*[]models.EpicSlim, error) {
	var body *bytes.Buffer
	params := url.Values{}
	var out []models.EpicSlim
	if err := a.request("GET", "/api/v3/milestones/"+fmt.Sprint(milestonePublicId)+"/epics", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
