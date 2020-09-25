package api

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// Get Repository returns information about the selected Repository.
func (a *api) GetRepository(repoPublicId int64) (*models.Repository, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out models.Repository
	if err := a.request("GET", "/api/v3/repositories/"+fmt.Sprint(repoPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
