package api

import (
	"net/url"
)

// Disables Groups for the current workspace2
func (a *api) DisableGroups() error {
	params := url.Values{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/groups/disable", params, nil, &out); err != nil {
		return err
	}
	return nil
}
