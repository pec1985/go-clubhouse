package api

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List all of the Epics with the Label.
func (a *api) ListLabelEpics(labelPublicId int64) (*[]models.EpicSlim, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out []models.EpicSlim
	if err := a.request("GET", "/api/v3/labels/"+fmt.Sprint(labelPublicId)+"/epics", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
