package api

import (
	"bytes"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// Returns information about the authenticated member.
func (a *api) GetCurrentMemberInfo() (*models.MemberInfo, error) {
	var body *bytes.Buffer
	params := url.Values{}
	var out models.MemberInfo
	if err := a.request("GET", "/api/v3/member", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
