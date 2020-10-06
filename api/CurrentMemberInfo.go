package api

import (
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/v1/models"
)

// Returns information about the authenticated member.
func (a *api) GetCurrentMemberInfo() (*models.MemberInfo, error) {
	params := url.Values{}
	var out models.MemberInfo
	if err := a.request("GET", "/api/v3/member", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
