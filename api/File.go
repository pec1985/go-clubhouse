package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/v1/api/models"
)

// List Files returns a list of all Files and related attributes in your Clubhouse.
func (a *api) ListFiles() (*[]models.File, error) {
	params := url.Values{}
	var out []models.File
	if err := a.request("GET", "/api/v3/files", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (a *api) CreateFiles(createFiles *models.CreateFiles) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createFiles != nil {
		jsonbody, _ := json.Marshal(createFiles)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/files", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Delete File can be used to delete any previously attached File.
func (a *api) DeleteFile(filePublicId int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/files/"+fmt.Sprint(filePublicId)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Get File returns information about the selected File.
func (a *api) GetFile(filePublicId int64) (*models.File, error) {
	params := url.Values{}
	var out models.File
	if err := a.request("GET", "/api/v3/files/"+fmt.Sprint(filePublicId)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update File can used to update the properties of a file uploaded to Clubhouse.
func (a *api) UpdateFile(filePublicId int64, updateFile *models.UpdateFile) (*models.File, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if updateFile != nil {
		jsonbody, _ := json.Marshal(updateFile)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.File
	if err := a.request("PUT", "/api/v3/files/"+fmt.Sprint(filePublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
