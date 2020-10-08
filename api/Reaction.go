package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// Delete a Reaction from any comment.
func (a *api) DeleteReaction(storyPublicId int64, commentPublicId int64, createOrDeleteReaction *models.CreateOrDeleteReaction) error {
	params := url.Values{}
	if createOrDeleteReaction != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(createOrDeleteReaction)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"/reactions", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Create a reaction to a comment.
func (a *api) CreateReaction(storyPublicId int64, commentPublicId int64, createOrDeleteReaction *models.CreateOrDeleteReaction) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createOrDeleteReaction != nil {
		jsonbody, _ := json.Marshal(createOrDeleteReaction)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"/reactions", params, body, &out); err != nil {
		return err
	}
	return nil
}
