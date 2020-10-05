package api

import (
	"bytes"
	"net/url"
)

// Disables Groups for the current workspace2
func (a *api) DisableGroups() error {
	var body *bytes.Buffer
	params := url.Values{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/groups/disable", params, body, &out); err != nil {
		return err
	}
	return nil
}
