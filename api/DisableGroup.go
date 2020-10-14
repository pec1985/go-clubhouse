package api

import (
	"net/url"
)

// DisableGroups disables Groups for the current workspace2
func (a *api) DisableGroups() error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("PUT", "/api/v3/groups/disable", params, nil, &out); err != nil {
		return err
	}
	return nil
}
