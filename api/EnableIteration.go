package api

import (
	"net/url"
)

// Enables Iterations for the current workspace
func (a *api) EnableIterations() error {
	params := url.Values{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/iterations/enable", params, nil, &out); err != nil {
		return err
	}
	return nil
}
