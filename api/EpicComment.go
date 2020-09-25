package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// Get a list of all Comments on an Epic.
func (a *api) ListEpicComments(epicPublicId int64) (*[]models.ThreadedComment, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out []models.ThreadedComment
	if err := a.request("GET", "/api/v3/epics/"+fmt.Sprint(epicPublicId)+"/comments", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// This endpoint allows you to create a threaded Comment on an Epic.
func (a *api) CreateEpicComment(epicPublicId int64, createEpicComment *models.CreateEpicComment) error {
	params := url.Values{}
	jsonbody, _ := json.Marshal(createEpicComment)
	body := bytes.NewBuffer(jsonbody)
	var out interface{}
	if err := a.request("POST", "/api/v3/epics/"+fmt.Sprint(epicPublicId)+"/comments", params, body, &out); err != nil {
		return err
	}
	return nil
}

// This endpoint allows you to delete a Comment from an Epic.
func (a *api) DeleteEpicComment(epicPublicId int64, commentPublicId int64) error {
	params := url.Values{}
	body := bytes.Buffer{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/epics/"+fmt.Sprint(epicPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"", params, body, &out); err != nil {
		return err
	}
	return nil
}

// This endpoint returns information about the selected Epic Comment.
func (a *api) GetEpicComment(epicPublicId int64, commentPublicId int64) (*models.ThreadedComment, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out models.ThreadedComment
	if err := a.request("GET", "/api/v3/epics/"+fmt.Sprint(epicPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// This endpoint allows you to update a threaded Comment on an Epic.
func (a *api) UpdateEpicComment(epicPublicId int64, commentPublicId int64, updateComment *models.UpdateComment) (*models.ThreadedComment, error) {
	params := url.Values{}
	jsonbody, _ := json.Marshal(updateComment)
	body := bytes.NewBuffer(jsonbody)
	var out models.ThreadedComment
	if err := a.request("PUT", "/api/v3/epics/"+fmt.Sprint(epicPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
