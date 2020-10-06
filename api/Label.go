package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/v1/api/models"
)

// List Labels returns a list of all Labels and their attributes.
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
	if err := a.request("GET", "/api/v3/labels", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create Label allows you to create a new Label in Clubhouse.
func (a *api) CreateLabel(createLabelParams *models.CreateLabelParams) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createLabelParams != nil {
		jsonbody, _ := json.Marshal(createLabelParams)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/labels", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Delete Label can be used to delete any Label.
func (a *api) DeleteLabel(labelPublicId int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/labels/"+fmt.Sprint(labelPublicId)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Get Label returns information about the selected Label.
func (a *api) GetLabel(labelPublicId int64) (*models.Label, error) {
	params := url.Values{}
	var out models.Label
	if err := a.request("GET", "/api/v3/labels/"+fmt.Sprint(labelPublicId)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update Label allows you to replace a Label name with another name. If you try to name a Label something that already exists, you will receive a 422 response.
func (a *api) UpdateLabel(labelPublicId int64, updateLabel *models.UpdateLabel) (*models.Label, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if updateLabel != nil {
		jsonbody, _ := json.Marshal(updateLabel)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Label
	if err := a.request("PUT", "/api/v3/labels/"+fmt.Sprint(labelPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
