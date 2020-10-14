package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// Create Comment allows you to create a Comment on any Story.
func (a *api) CreateComment(storyID int64, comment *models.CreateComment) error {
	params := url.Values{}
	var body *bytes.Buffer
	if comment != nil {
		jsonbody, _ := json.Marshal(comment)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/stories/"+fmt.Sprint(storyID)+"/comments", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Delete a Comment from any story.
func (a *api) DeleteComment(storyID int64, commentID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/stories/"+fmt.Sprint(storyID)+"/comments/"+fmt.Sprint(commentID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Get Comment is used to get Comment information.
func (a *api) GetComment(storyID int64, commentID int64) (*models.Comment, error) {
	params := url.Values{}
	var out models.Comment
	if err := a.Request("GET", "/api/v3/stories/"+fmt.Sprint(storyID)+"/comments/"+fmt.Sprint(commentID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update Comment replaces the text of the existing Comment.
func (a *api) UpdateComment(storyID int64, commentID int64, comment *models.UpdateComment) (*models.Comment, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if comment != nil {
		jsonbody, _ := json.Marshal(comment)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Comment
	if err := a.Request("PUT", "/api/v3/stories/"+fmt.Sprint(storyID)+"/comments/"+fmt.Sprint(commentID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
