package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/v1/api/models"
)

func (a *api) ListIterations() (*[]models.IterationSlim, error) {
	params := url.Values{}
	var out []models.IterationSlim
	if err := a.request("GET", "/api/v3/iterations", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (a *api) CreateIteration(createIteration *models.CreateIteration) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createIteration != nil {
		jsonbody, _ := json.Marshal(createIteration)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/iterations", params, body, &out); err != nil {
		return err
	}
	return nil
}
func (a *api) DeleteIteration(iterationPublicId int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/iterations/"+fmt.Sprint(iterationPublicId)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}
func (a *api) GetIteration(iterationPublicId int64) (*models.Iteration, error) {
	params := url.Values{}
	var out models.Iteration
	if err := a.request("GET", "/api/v3/iterations/"+fmt.Sprint(iterationPublicId)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (a *api) UpdateIteration(iterationPublicId int64, updateIteration *models.UpdateIteration) (*models.Iteration, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if updateIteration != nil {
		jsonbody, _ := json.Marshal(updateIteration)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Iteration
	if err := a.request("PUT", "/api/v3/iterations/"+fmt.Sprint(iterationPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
