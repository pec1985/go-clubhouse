package api

import (
	"net/url"
)

// Enables Groups for the current workspace2
func (a *api) EnableGroups() error {
	params := url.Values{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/groups/enable", params, nil, &out); err != nil {
		return err
	}
	return nil
}
