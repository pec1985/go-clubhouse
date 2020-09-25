package api

import (
	"bytes"
	"net/url"
)

// Disables the Story Template feature for the given Organization.
func (a *api) DisableStoryTemplates() error {
	params := url.Values{}
	body := bytes.Buffer{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/entity-templates/disable", params, body, &out); err != nil {
		return err
	}
	return nil
}
