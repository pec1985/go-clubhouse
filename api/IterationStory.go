package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// Get a list of all Stories in an Iteration.
func (a *api) ListIterationStories(iterationID int64, getIterationStories *models.GetIterationStories) (*[]models.StorySlim, error) {
	params := url.Values{}
	if getIterationStories != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(getIterationStories)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out []models.StorySlim
	if err := a.Request("GET", "/api/v3/iterations/"+fmt.Sprint(iterationID)+"/stories", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
