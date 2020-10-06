package api

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// List all the entity templates for an organization.
func (a *api) ListEntityTemplates() (*[]models.EntityTemplate, error) {
	params := url.Values{}
	var out []models.EntityTemplate
	if err := a.request("GET", "/api/v3/entity-templates", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create a new entity template for your organization.
func (a *api) CreateEntityTemplate(createEntityTemplate *models.CreateEntityTemplate) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createEntityTemplate != nil {
		jsonbody, _ := json.Marshal(createEntityTemplate)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/entity-templates", params, body, &out); err != nil {
		return err
	}
	return nil
}
func (a *api) DeleteEntityTemplate(entityTemplatePublicId string) error {
	params := url.Values{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/entity-templates/"+entityTemplatePublicId+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Get Entity Template returns information about a given entity template.
func (a *api) GetEntityTemplate(entityTemplatePublicId string) (*models.EntityTemplate, error) {
	params := url.Values{}
	var out models.EntityTemplate
	if err := a.request("GET", "/api/v3/entity-templates/"+entityTemplatePublicId+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update an entity template's name or its contents.
func (a *api) UpdateEntityTemplate(entityTemplatePublicId string, updateEntityTemplate *models.UpdateEntityTemplate) (*models.EntityTemplate, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if updateEntityTemplate != nil {
		jsonbody, _ := json.Marshal(updateEntityTemplate)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.EntityTemplate
	if err := a.request("PUT", "/api/v3/entity-templates/"+entityTemplatePublicId+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
