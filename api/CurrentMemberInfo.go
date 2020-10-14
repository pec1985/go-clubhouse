package api

import (
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// GetCurrentMemberInfo Returns information about the authenticated member.
func (a *api) GetCurrentMemberInfo() (*models.MemberInfo, error) {
	params := url.Values{}
	var out models.MemberInfo
	if err := a.Request("GET", "/api/v3/member", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
