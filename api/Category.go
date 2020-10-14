package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListCategories returns a list of all Categories and their attributes.
func (a *api) ListCategories() (*[]models.Category, error) {
	params := url.Values{}
	var out []models.Category
	if err := a.Request("GET", "/api/v3/categories", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateCategory allows you to create a new Category in Clubhouse.
func (a *api) CreateCategory(category *models.CreateCategory) error {
	params := url.Values{}
	var body *bytes.Buffer
	if category != nil {
		jsonbody, _ := json.Marshal(category)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/categories", params, body, &out); err != nil {
		return err
	}
	return nil
}

// DeleteCategory can be used to delete any Category.
func (a *api) DeleteCategory(categoryID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/categories/"+fmt.Sprint(categoryID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// GetCategory returns information about the selected Category.
func (a *api) GetCategory(categoryID int64) (*models.Category, error) {
	params := url.Values{}
	var out models.Category
	if err := a.Request("GET", "/api/v3/categories/"+fmt.Sprint(categoryID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateCategory allows you to replace a Category name with another name. If you try to name a Category something that already exists, you will receive a 422 response.
func (a *api) UpdateCategory(categoryID int64, category *models.UpdateCategory) (*models.Category, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if category != nil {
		jsonbody, _ := json.Marshal(category)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Category
	if err := a.Request("PUT", "/api/v3/categories/"+fmt.Sprint(categoryID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
