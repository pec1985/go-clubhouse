package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/v1/models"
)

// List Epics returns a list of all Epics and their attributes.
func (a *api) ListEpics(listEpics *models.ListEpics) (*[]models.EpicSlim, error) {
	params := url.Values{}
	if listEpics != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(listEpics)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out []models.EpicSlim
	if err := a.request("GET", "/api/v3/epics", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create Epic allows you to create a new Epic in Clubhouse.
func (a *api) CreateEpic(createEpic *models.CreateEpic) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createEpic != nil {
		jsonbody, _ := json.Marshal(createEpic)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/epics", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Delete Epic can be used to delete the Epic. The only required parameter is Epic ID.
func (a *api) DeleteEpic(epicPublicId int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/epics/"+fmt.Sprint(epicPublicId)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Get Epic returns information about the selected Epic.
func (a *api) GetEpic(epicPublicId int64) (*models.Epic, error) {
	params := url.Values{}
	var out models.Epic
	if err := a.request("GET", "/api/v3/epics/"+fmt.Sprint(epicPublicId)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update Epic can be used to update numerous fields in the Epic. The only required parameter is Epic ID, which can be found in the Clubhouse UI.
func (a *api) UpdateEpic(epicPublicId int64, updateEpic *models.UpdateEpic) (*models.Epic, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if updateEpic != nil {
		jsonbody, _ := json.Marshal(updateEpic)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Epic
	if err := a.request("PUT", "/api/v3/epics/"+fmt.Sprint(epicPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
