package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// This endpoint allows you to create a nested Comment reply to an existing Epic Comment.
func (a *api) CreateEpicCommentComment(epicPublicId int64, commentPublicId int64, createCommentComment *models.CreateCommentComment) error {
	params := url.Values{}
	jsonbody, _ := json.Marshal(createCommentComment)
	body := bytes.NewBuffer(jsonbody)
	var out interface{}
	if err := a.request("POST", "/api/v3/epics/"+fmt.Sprint(epicPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"", params, body, &out); err != nil {
		return err
	}
	return nil
}
