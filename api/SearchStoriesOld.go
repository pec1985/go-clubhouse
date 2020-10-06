package api

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/v1/models"
)

// Search Stories lets you search Stories based on desired parameters.
func (a *api) SearchStoriesOld(searchStories *models.SearchStories) error {
	params := url.Values{}
	var body *bytes.Buffer
	if searchStories != nil {
		jsonbody, _ := json.Marshal(searchStories)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/stories/search", params, body, &out); err != nil {
		return err
	}
	return nil
}
