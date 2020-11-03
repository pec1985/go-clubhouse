package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListEpics returns a list of all Epics and their attributes.
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
	if err := a.Request("GET", "/api/v3/epics", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateEpic allows you to create a new Epic in Clubhouse.
func (a *api) CreateEpic(epic *models.CreateEpic) error {
	params := url.Values{}
	var body *bytes.Buffer
	if epic != nil {
		jsonbody, _ := toPayload(epic, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/epics", params, body, &out); err != nil {
		return err
	}
	return nil
}

// DeleteEpic can be used to delete the Epic. The only required parameter is Epic ID.
func (a *api) DeleteEpic(epicID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/epics/"+fmt.Sprint(epicID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// GetEpic returns information about the selected Epic.
func (a *api) GetEpic(epicID int64) (*models.Epic, error) {
	params := url.Values{}
	var out models.Epic
	if err := a.Request("GET", "/api/v3/epics/"+fmt.Sprint(epicID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateEpic can be used to update numerous fields in the Epic. The only required parameter is Epic ID, which can be found in the Clubhouse UI.
func (a *api) UpdateEpic(epicID int64, epic *models.UpdateEpic) (*models.Epic, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if epic != nil {
		jsonbody, _ := toPayload(epic, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Epic
	if err := a.Request("PUT", "/api/v3/epics/"+fmt.Sprint(epicID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
