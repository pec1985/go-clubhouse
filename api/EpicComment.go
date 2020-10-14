package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// Get a list of all Comments on an Epic.
func (a *api) ListEpicComments(epicID int64) (*[]models.ThreadedComment, error) {
	params := url.Values{}
	var out []models.ThreadedComment
	if err := a.Request("GET", "/api/v3/epics/"+fmt.Sprint(epicID)+"/comments", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// This endpoint allows you to create a threaded Comment on an Epic.
func (a *api) CreateEpicComment(epicID int64, epicComment *models.CreateEpicComment) error {
	params := url.Values{}
	var body *bytes.Buffer
	if epicComment != nil {
		jsonbody, _ := json.Marshal(epicComment)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/epics/"+fmt.Sprint(epicID)+"/comments", params, body, &out); err != nil {
		return err
	}
	return nil
}

// This endpoint allows you to delete a Comment from an Epic.
func (a *api) DeleteEpicComment(epicID int64, commentID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/epics/"+fmt.Sprint(epicID)+"/comments/"+fmt.Sprint(commentID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// This endpoint returns information about the selected Epic Comment.
func (a *api) GetEpicComment(epicID int64, commentID int64) (*models.ThreadedComment, error) {
	params := url.Values{}
	var out models.ThreadedComment
	if err := a.Request("GET", "/api/v3/epics/"+fmt.Sprint(epicID)+"/comments/"+fmt.Sprint(commentID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// This endpoint allows you to update a threaded Comment on an Epic.
func (a *api) UpdateEpicComment(epicID int64, commentID int64, comment *models.UpdateComment) (*models.ThreadedComment, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if comment != nil {
		jsonbody, _ := json.Marshal(comment)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.ThreadedComment
	if err := a.Request("PUT", "/api/v3/epics/"+fmt.Sprint(epicID)+"/comments/"+fmt.Sprint(commentID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
