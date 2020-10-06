## Unofficial Cloubhouse Go SDK

This API is generated from the swagger file located in https://clubhouse.io/api/rest/v3/clubhouse.swagger.json

To just use it, `go get` it

```
go get github.com/pec1985/go-clubhouse.io/v1/api v1.0.0
```

Sample api call

```
package main

import (
	"fmt"
	"net/http"

	"github.com/pec1985/go-clubhouse.io/v1/api"
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
```

To play with the code, clone the repo

```
git clone https://github.com/pec1985/go-clubhouse.io.git 
```

To generate new code, simply cd to the `generator` folder and run the main.go

```
cd $(GOPATH)/src/github.com/pec1985/go-clubhouse.io
cd generator
go run .
```

Open the repo in VSCode and start playing with it.

Pedro Enrique
Sept 25th, 2020
MIT License