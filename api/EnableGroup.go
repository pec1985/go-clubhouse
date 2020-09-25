package api

import (
	"bytes"
	"net/url"
)

// Enables Groups for the current workspace2
func (a *api) EnableGroups() error {
	params := url.Values{}
	body := bytes.Buffer{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/groups/enable", params, body, &out); err != nil {
		return err
	}
	return nil
}
