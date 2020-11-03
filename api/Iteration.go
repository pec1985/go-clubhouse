package api

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

func (a *api) ListIterations() (*[]models.IterationSlim, error) {
	params := url.Values{}
	var out []models.IterationSlim
	if err := a.Request("GET", "/api/v3/iterations", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (a *api) CreateIteration(iteration *models.CreateIteration) error {
	params := url.Values{}
	var body *bytes.Buffer
	if iteration != nil {
		jsonbody, _ := toPayload(iteration, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/iterations", params, body, &out); err != nil {
		return err
	}
	return nil
}
func (a *api) DeleteIteration(iterationID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/iterations/"+fmt.Sprint(iterationID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}
func (a *api) GetIteration(iterationID int64) (*models.Iteration, error) {
	params := url.Values{}
	var out models.Iteration
	if err := a.Request("GET", "/api/v3/iterations/"+fmt.Sprint(iterationID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (a *api) UpdateIteration(iterationID int64, iteration *models.UpdateIteration) (*models.Iteration, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if iteration != nil {
		jsonbody, _ := toPayload(iteration, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Iteration
	if err := a.Request("PUT", "/api/v3/iterations/"+fmt.Sprint(iterationID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
