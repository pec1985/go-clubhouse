package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// DeleteReaction delete a Reaction from any comment.
func (a *api) DeleteReaction(storyID int64, commentID int64, orDeleteReaction *models.CreateOrDeleteReaction) error {
	params := url.Values{}
	if orDeleteReaction != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(orDeleteReaction)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/stories/"+fmt.Sprint(storyID)+"/comments/"+fmt.Sprint(commentID)+"/reactions", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// CreateReaction create a reaction to a comment.
func (a *api) CreateReaction(storyID int64, commentID int64, orDeleteReaction *models.CreateOrDeleteReaction) error {
	params := url.Values{}
	var body *bytes.Buffer
	if orDeleteReaction != nil {
		jsonbody, _ := toPayload(orDeleteReaction, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/stories/"+fmt.Sprint(storyID)+"/comments/"+fmt.Sprint(commentID)+"/reactions", params, body, &out); err != nil {
		return err
	}
	return nil
}
