package api

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List Category Milestones returns a list of all Milestones with the Category.
func (a *api) ListCategoryMilestones(categoryPublicId int64) (*[]models.Milestone, error) {
	var body *bytes.Buffer
	params := url.Values{}
	var out []models.Milestone
	if err := a.request("GET", "/api/v3/categories/"+fmt.Sprint(categoryPublicId)+"/milestones", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
