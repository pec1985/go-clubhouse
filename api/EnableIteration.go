package api

import (
	"bytes"
	"net/url"
)

// Enables Iterations for the current workspace
func (a *api) EnableIterations() error {
	params := url.Values{}
	body := bytes.Buffer{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/iterations/enable", params, body, &out); err != nil {
		return err
	}
	return nil
}
