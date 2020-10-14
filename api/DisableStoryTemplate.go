package api

import (
	"net/url"
)

// DisableStoryTemplates disables the Story Template feature for the given Organization.
func (a *api) DisableStoryTemplates() error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("PUT", "/api/v3/entity-templates/disable", params, nil, &out); err != nil {
		return err
	}
	return nil
}
