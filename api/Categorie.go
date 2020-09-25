package api

import (
	"bytes"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List Categories returns a list of all Categories and their attributes.
func (a *api) ListCategories() (*[]models.Category, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out []models.Category
	if err := a.request("GET", "/api/v3/categories", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
