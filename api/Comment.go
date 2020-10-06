package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/v1/api/models"
)

// Create Comment allows you to create a Comment on any Story.
func (a *api) CreateComment(storyPublicId int64, createComment *models.CreateComment) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createComment != nil {
		jsonbody, _ := json.Marshal(createComment)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/comments", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Delete a Comment from any story.
func (a *api) DeleteComment(storyPublicId int64, commentPublicId int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Get Comment is used to get Comment information.
func (a *api) GetComment(storyPublicId int64, commentPublicId int64) (*models.Comment, error) {
	params := url.Values{}
	var out models.Comment
	if err := a.request("GET", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update Comment replaces the text of the existing Comment.
func (a *api) UpdateComment(storyPublicId int64, commentPublicId int64, updateComment *models.UpdateComment) (*models.Comment, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if updateComment != nil {
		jsonbody, _ := json.Marshal(updateComment)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Comment
	if err := a.request("PUT", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
