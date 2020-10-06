package api

import (
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/v1/api/models"
)

// List Category Milestones returns a list of all Milestones with the Category.
func (a *api) ListCategoryMilestones(categoryPublicId int64) (*[]models.Milestone, error) {
	params := url.Values{}
	var out []models.Milestone
	if err := a.request("GET", "/api/v3/categories/"+fmt.Sprint(categoryPublicId)+"/milestones", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
