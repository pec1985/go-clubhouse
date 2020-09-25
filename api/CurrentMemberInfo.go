package api

import (
	"bytes"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// Returns information about the authenticated member.
func (a *api) GetCurrentMemberInfo() (*models.MemberInfo, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out models.MemberInfo
	if err := a.request("GET", "/api/v3/member", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
