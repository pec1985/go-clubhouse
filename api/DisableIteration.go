package api

import (
	"bytes"
	"net/url"
)

// Disables Iterations for the current workspace
func (a *api) DisableIterations() error {
	var body *bytes.Buffer
	params := url.Values{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/iterations/disable", params, body, &out); err != nil {
		return err
	}
	return nil
}
