package api

import (
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListCategoryMilestones returns a list of all Milestones with the Category.
func (a *api) ListCategoryMilestones(categoryID int64) (*[]models.Milestone, error) {
	params := url.Values{}
	var out []models.Milestone
	if err := a.Request("GET", "/api/v3/categories/"+fmt.Sprint(categoryID)+"/milestones", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
