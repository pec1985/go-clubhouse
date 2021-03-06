package api

import (
	"bytes"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// SearchStoriesOld search Stories lets you search Stories based on desired parameters.
func (a *api) SearchStoriesOld(searchStories *models.SearchStories) error {
	params := url.Values{}
	var body *bytes.Buffer
	if searchStories != nil {
		jsonbody, _ := toPayload(searchStories, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/stories/search", params, body, &out); err != nil {
		return err
	}
	return nil
}
