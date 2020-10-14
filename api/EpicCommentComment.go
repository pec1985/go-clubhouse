package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// CreateEpicCommentComment this endpoint allows you to create a nested Comment reply to an existing Epic Comment.
func (a *api) CreateEpicCommentComment(epicID int64, commentID int64, commentComment *models.CreateCommentComment) error {
	params := url.Values{}
	var body *bytes.Buffer
	if commentComment != nil {
		jsonbody, _ := json.Marshal(commentComment)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/epics/"+fmt.Sprint(epicID)+"/comments/"+fmt.Sprint(commentID)+"", params, body, &out); err != nil {
		return err
	}
	return nil
}
