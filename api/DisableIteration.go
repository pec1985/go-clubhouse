package api

import (
	"net/url"
)

// Disables Iterations for the current workspace
func (a *api) DisableIterations() error {
	params := url.Values{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/iterations/disable", params, nil, &out); err != nil {
		return err
	}
	return nil
}
