package api

import (
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListLabelEpics list all of the Epics with the Label.
func (a *api) ListLabelEpics(labelID int64) (*[]models.EpicSlim, error) {
	params := url.Values{}
	var out []models.EpicSlim
	if err := a.Request("GET", "/api/v3/labels/"+fmt.Sprint(labelID)+"/epics", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
