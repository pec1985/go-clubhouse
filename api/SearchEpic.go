package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// Search Epics lets you search Epics based on desired parameters. Since ordering of stories can change over time (due to search ranking decay, new Epics being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.
func (a *api) SearchEpics(search *models.Search) (*models.EpicSearchResults, error) {
	var body *bytes.Buffer
	params := url.Values{}
	if search != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(search)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out models.EpicSearchResults
	if err := a.request("GET", "/api/v3/search/epics", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
