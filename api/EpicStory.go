package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/v1/api/models"
)

// Get a list of all Stories in an Epic.
func (a *api) ListEpicStories(epicPublicId int64, getEpicStories *models.GetEpicStories) (*[]models.StorySlim, error) {
	params := url.Values{}
	if getEpicStories != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(getEpicStories)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out []models.StorySlim
	if err := a.request("GET", "/api/v3/epics/"+fmt.Sprint(epicPublicId)+"/stories", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
