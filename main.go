package main

import (
	"fmt"
	"net/http"

	"github.com/pec1985/go-clubhouse/v1/api"
)

func main() {

	// sample api call
	token := "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	api := api.New(http.DefaultClient, token)
	projects, err := api.ListProjects()
	if err != nil {
		panic(err)
	}
	for _, p := range *projects {
		fmt.Println(p.StringifyPretty())
	}
}
