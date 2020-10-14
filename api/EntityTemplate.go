package api

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListEntityTemplates list all the entity templates for an organization.
func (a *api) ListEntityTemplates() (*[]models.EntityTemplate, error) {
	params := url.Values{}
	var out []models.EntityTemplate
	if err := a.Request("GET", "/api/v3/entity-templates", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateEntityTemplate create a new entity template for your organization.
func (a *api) CreateEntityTemplate(entityTemplate *models.CreateEntityTemplate) error {
	params := url.Values{}
	var body *bytes.Buffer
	if entityTemplate != nil {
		jsonbody, _ := json.Marshal(entityTemplate)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/entity-templates", params, body, &out); err != nil {
		return err
	}
	return nil
}
func (a *api) DeleteEntityTemplate(entityTemplateID string) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/entity-templates/"+entityTemplateID+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// GetEntityTemplate returns information about a given entity template.
func (a *api) GetEntityTemplate(entityTemplateID string) (*models.EntityTemplate, error) {
	params := url.Values{}
	var out models.EntityTemplate
	if err := a.Request("GET", "/api/v3/entity-templates/"+entityTemplateID+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateEntityTemplate update an entity template's name or its contents.
func (a *api) UpdateEntityTemplate(entityTemplateID string, entityTemplate *models.UpdateEntityTemplate) (*models.EntityTemplate, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if entityTemplate != nil {
		jsonbody, _ := json.Marshal(entityTemplate)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.EntityTemplate
	if err := a.Request("PUT", "/api/v3/entity-templates/"+entityTemplateID+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
