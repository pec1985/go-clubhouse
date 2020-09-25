package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

// Update Category allows you to replace a Category name with another name. If you try to name a Category something that already exists, you will receive a 422 response.
func (a *api) UpdateCategory(categoryPublicId int64, updateCategory *models.UpdateCategory) (*models.Category, error) {
	params := url.Values{}
	jsonbody, _ := json.Marshal(updateCategory)
	body := bytes.NewBuffer(jsonbody)
	var out models.Category
	if err := a.request("PUT", "/api/v3/categories/"+fmt.Sprint(categoryPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete Category can be used to delete any Category.
func (a *api) DeleteCategory(categoryPublicId int64) error {
	params := url.Values{}
	body := bytes.Buffer{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/categories/"+fmt.Sprint(categoryPublicId)+"", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Get Category returns information about the selected Category.
func (a *api) GetCategory(categoryPublicId int64) (*models.Category, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out models.Category
	if err := a.request("GET", "/api/v3/categories/"+fmt.Sprint(categoryPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create Category allows you to create a new Category in Clubhouse.
func (a *api) CreateCategory(createCategory *models.CreateCategory) error {
	params := url.Values{}
	jsonbody, _ := json.Marshal(createCategory)
	body := bytes.NewBuffer(jsonbody)
	var out interface{}
	if err := a.request("POST", "/api/v3/categories", params, body, &out); err != nil {
		return err
	}
	return nil
}
