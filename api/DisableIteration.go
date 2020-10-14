package api

import (
	"net/url"
)

// DisableIterations disables Iterations for the current workspace
func (a *api) DisableIterations() error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("PUT", "/api/v3/iterations/disable", params, nil, &out); err != nil {
		return err
	}
	return nil
}
