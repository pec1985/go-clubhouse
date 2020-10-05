package api

import (
	"bytes"
	"net/url"
)

// Enables the Story Template feature for the given Organization.
func (a *api) EnableStoryTemplates() error {
	var body *bytes.Buffer
	params := url.Values{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/entity-templates/enable", params, body, &out); err != nil {
		return err
	}
	return nil
}
