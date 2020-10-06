package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/v1/api/models"
)

// Search lets you search Epics and Stories based on desired parameters. Since ordering of the results can change over time (due to search ranking decay, new Epics and Stories being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.
func (a *api) Search(search *models.Search) (*models.SearchResults, error) {
	params := url.Values{}
	if search != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(search)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out models.SearchResults
	if err := a.request("GET", "/api/v3/search", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
