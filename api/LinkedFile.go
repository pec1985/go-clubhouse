package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// List Linked Files returns a list of all Linked-Files and their attributes.
func (a *api) ListLinkedFiles() (*[]models.LinkedFile, error) {
	var body *bytes.Buffer
	params := url.Values{}
	var out []models.LinkedFile
	if err := a.request("GET", "/api/v3/linked-files", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create Linked File allows you to create a new Linked File in Clubhouse.
func (a *api) CreateLinkedFile(createLinkedFile *models.CreateLinkedFile) error {
	var body *bytes.Buffer
	params := url.Values{}
	if createLinkedFile != nil {
		jsonbody, _ := json.Marshal(createLinkedFile)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/linked-files", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Delete Linked File can be used to delete any previously attached Linked-File.
func (a *api) DeleteLinkedFile(linkedFilePublicId int64) error {
	var body *bytes.Buffer
	params := url.Values{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/linked-files/"+fmt.Sprint(linkedFilePublicId)+"", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Get File returns information about the selected Linked File.
func (a *api) GetLinkedFile(linkedFilePublicId int64) (*models.LinkedFile, error) {
	var body *bytes.Buffer
	params := url.Values{}
	var out models.LinkedFile
	if err := a.request("GET", "/api/v3/linked-files/"+fmt.Sprint(linkedFilePublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Updated Linked File allows you to update properties of a previously attached Linked-File.
func (a *api) UpdateLinkedFile(linkedFilePublicId int64, updateLinkedFile *models.UpdateLinkedFile) (*models.LinkedFile, error) {
	var body *bytes.Buffer
	params := url.Values{}
	if updateLinkedFile != nil {
		jsonbody, _ := json.Marshal(updateLinkedFile)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.LinkedFile
	if err := a.request("PUT", "/api/v3/linked-files/"+fmt.Sprint(linkedFilePublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
