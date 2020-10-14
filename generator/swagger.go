package main

import (
	"encoding/json"
	"net/http"
)

func downloadSwaggerFile() (*swaggerPayload, error) {

	resp, err := http.Get("https://clubhouse.io/api/rest/v3/clubhouse.swagger.json")
	if err != nil {
		return nil, err
	}
	var data *swaggerPayload
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

type swaggerPayload struct {
	Consumes    *[]string `json:"consumes"`
	Definitions *map[string]swaggerPayloadDefinition
	Info        *struct {
		Description *string `json:"description"`
		Title       *string `json:"title"`
		Version     *string `json:"version"`
	} `json:"info"`
	Paths *map[string]map[string]swaggerPayloadPath `json:"paths"`
}
type swaggerPayloadDefinition struct {
	AdditionalProperties *json.RawMessage `json:"additionalProperties"`
	Description          *string          `json:"description"`
	Properties           *map[string]swaggerPayloadDefinitionProperty
	Required             *[]string `json:"required"`
	Type                 *string   `json:"type"`
}

type swaggerPayloadPath struct {
	Description *string `json:"description"`
	Parameters  *[]struct {
		Description *string `json:"description"`
		In          *string `json:"in"`
		Name        *string `json:"name"`
		Required    *bool   `json:"required"`
		Format      *string `json:"format"`
		Type        *string `json:"type"`
		Schema      *struct {
			Ref *string `json:"$ref"`
		} `json:"schema"`
	} `json:"parameters"`
	Responses *map[string]swaggerPayloadPathDescription `json:"responses"`
	Summary   *string                                   `json:"summary"`
}

type swaggerPayloadPathDescription struct {
	Description *string `json:"description"`
	Schema      *struct {
		Ref   *string `json:"$ref"`
		Items *struct {
			Ref *string `json:"$ref"`
		} `json:"items"`
		Type *string `json:"type"`
	} `json:"schema"`
}

type swaggerPayloadDefinitionProperty struct {
	Items *struct {
		Format *string `json:"format"`
		Type   *string `json:"type"`
		Ref    *string `json:"$ref"`
	} `json:"items"`
	Description *string `json:"description"`
	Format      *string `json:"format"`
	Type        *string `json:"type"`
	Ref         *string `json:"$ref"`
	XNullable   *bool   `json:"x-nullable"`
}
