package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// API the api interface
type API interface {
	// Request you can use this method to call any api that's missing from the sdk
	Request(method string, endpoint string, params url.Values, data io.Reader, out interface{}) error
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

// Request you can use this method to call any api that's missing from the sdk
func (a *api) Request(method string, endpoint string, params url.Values, data io.Reader, out interface{}) error {

	if params == nil {
		params = url.Values{}
	}
	ur := strings.TrimRight(a.url, "/")
	endpoint = strings.TrimPrefix(endpoint, "/")
	var req *http.Request
	var err error
	if data == nil {
		req, err = http.NewRequest(method, ur+"/"+endpoint+"?"+params.Encode(), nil)
	} else {
		req, err = http.NewRequest(method, ur+"/"+endpoint+"?"+params.Encode(), data)
	}
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

// this removes empty dates
func toPayload(in interface{}, pretty bool) ([]byte, error) {
	var b []byte
	var err error
	if b, err = json.Marshal(in); err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	var marshaller func(m map[string]interface{})
	marshaller = func(m map[string]interface{}) {
		for k, v := range m {
			if child, ok := v.(map[string]interface{}); ok {
				marshaller(child)
				continue
			}
			if s, ok := v.(string); ok {
				if t, e := time.Parse(time.RFC3339Nano, s); e == nil {
					if t.IsZero() {
						delete(m, k)
					}
				}
			}
		}
	}
	marshaller(m)
	if pretty {
		json.MarshalIndent(m, "", "  ")
	}
	return json.Marshal(m)
}
