package api

import (
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListMilestoneEpics List all of the Epics within the Milestone.
func (a *api) ListMilestoneEpics(milestoneID int64) (*[]models.EpicSlim, error) {
	params := url.Values{}
	var out []models.EpicSlim
	if err := a.Request("GET", "/api/v3/milestones/"+fmt.Sprint(milestoneID)+"/epics", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
