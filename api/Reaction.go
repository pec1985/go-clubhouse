package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// Delete a Reaction from any comment.
func (a *api) DeleteReaction(storyPublicId int64, commentPublicId int64, createOrDeleteReaction *models.CreateOrDeleteReaction) error {
	params := url.Values{}
	{
		kv := map[string]interface{}{}
		b, _ := json.Marshal(createOrDeleteReaction)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	body := bytes.Buffer{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"/reactions", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Create a reaction to a comment.
func (a *api) CreateReaction(storyPublicId int64, commentPublicId int64, createOrDeleteReaction *models.CreateOrDeleteReaction) error {
	params := url.Values{}
	jsonbody, _ := json.Marshal(createOrDeleteReaction)
	body := bytes.NewBuffer(jsonbody)
	var out interface{}
	if err := a.request("POST", "/api/v3/stories/"+fmt.Sprint(storyPublicId)+"/comments/"+fmt.Sprint(commentPublicId)+"/reactions", params, body, &out); err != nil {
		return err
	}
	return nil
}
