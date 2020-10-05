package api

import (
	"bytes"
	"net/url"
)

// Enables Iterations for the current workspace
func (a *api) EnableIterations() error {
	var body *bytes.Buffer
	params := url.Values{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/iterations/enable", params, body, &out); err != nil {
		return err
	}
	return nil
}
