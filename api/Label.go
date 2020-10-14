package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListLabels returns a list of all Labels and their attributes.
func (a *api) ListLabels(listLabels *models.ListLabels) (*[]models.Label, error) {
	params := url.Values{}
	if listLabels != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(listLabels)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out []models.Label
	if err := a.Request("GET", "/api/v3/labels", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateLabel allows you to create a new Label in Clubhouse.
func (a *api) CreateLabel(labelParams *models.CreateLabelParams) error {
	params := url.Values{}
	var body *bytes.Buffer
	if labelParams != nil {
		jsonbody, _ := json.Marshal(labelParams)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/labels", params, body, &out); err != nil {
		return err
	}
	return nil
}

// DeleteLabel can be used to delete any Label.
func (a *api) DeleteLabel(labelID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/labels/"+fmt.Sprint(labelID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// GetLabel returns information about the selected Label.
func (a *api) GetLabel(labelID int64) (*models.Label, error) {
	params := url.Values{}
	var out models.Label
	if err := a.Request("GET", "/api/v3/labels/"+fmt.Sprint(labelID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateLabel allows you to replace a Label name with another name. If you try to name a Label something that already exists, you will receive a 422 response.
func (a *api) UpdateLabel(labelID int64, label *models.UpdateLabel) (*models.Label, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if label != nil {
		jsonbody, _ := json.Marshal(label)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Label
	if err := a.Request("PUT", "/api/v3/labels/"+fmt.Sprint(labelID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
