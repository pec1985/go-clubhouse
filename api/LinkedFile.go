package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListLinkedFiles returns a list of all Linked-Files and their attributes.
func (a *api) ListLinkedFiles() (*[]models.LinkedFile, error) {
	params := url.Values{}
	var out []models.LinkedFile
	if err := a.Request("GET", "/api/v3/linked-files", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateLinkedFile allows you to create a new Linked File in Clubhouse.
func (a *api) CreateLinkedFile(linkedFile *models.CreateLinkedFile) error {
	params := url.Values{}
	var body *bytes.Buffer
	if linkedFile != nil {
		jsonbody, _ := json.Marshal(linkedFile)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/linked-files", params, body, &out); err != nil {
		return err
	}
	return nil
}

// DeleteLinkedFile can be used to delete any previously attached Linked-File.
func (a *api) DeleteLinkedFile(linkedFileID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/linked-files/"+fmt.Sprint(linkedFileID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// GetLinkedFile Get File returns information about the selected Linked File.
func (a *api) GetLinkedFile(linkedFileID int64) (*models.LinkedFile, error) {
	params := url.Values{}
	var out models.LinkedFile
	if err := a.Request("GET", "/api/v3/linked-files/"+fmt.Sprint(linkedFileID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateLinkedFile Updated Linked File allows you to update properties of a previously attached Linked-File.
func (a *api) UpdateLinkedFile(linkedFileID int64, linkedFile *models.UpdateLinkedFile) (*models.LinkedFile, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if linkedFile != nil {
		jsonbody, _ := json.Marshal(linkedFile)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.LinkedFile
	if err := a.Request("PUT", "/api/v3/linked-files/"+fmt.Sprint(linkedFileID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
