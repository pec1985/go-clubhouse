package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// API the api interface
type API interface {
	// functions
}

// New creates a new instance of the api
func New(client *http.Client, token string) API {
	return &api{
		client: client,
		token:  token,
		url:    "https://api.clubhouse.io/",
	}
}

type api struct {
	client *http.Client
	token  string
	url    string
}

func (a *api) request(method string, endpoint string, params url.Values, data io.Reader, out interface{}) error {

	ur := strings.TrimRight(a.url, "/")
	endpoint = strings.TrimPrefix(endpoint, "/")

	req, err := http.NewRequest(http.MethodGet, ur+"/"+endpoint+"?"+params.Encode(), data)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Clubhouse-Token", a.token)
	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	return fmt.Errorf("status code: %v. message: %v", resp.StatusCode, string(b))
}
