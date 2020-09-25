package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List Stories returns a list of all Stories in a selected Project and their attributes.
func (a *api) ListStories(projectPublicId int64, getProjectStories *models.GetProjectStories) (*[]models.StorySlim, error) {
	params := url.Values{}
	{
		kv := map[string]interface{}{}
		b, _ := json.Marshal(getProjectStories)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	body := bytes.Buffer{}
	var out []models.StorySlim
	if err := a.request("GET", "/api/v3/projects/"+fmt.Sprint(projectPublicId)+"/stories", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
