package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/v1/models"
)

// List all of the Stories with the Label.
func (a *api) ListLabelStories(labelPublicId int64, getLabelStories *models.GetLabelStories) (*[]models.StorySlim, error) {
	params := url.Values{}
	if getLabelStories != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(getLabelStories)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out []models.StorySlim
	if err := a.request("GET", "/api/v3/labels/"+fmt.Sprint(labelPublicId)+"/stories", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
