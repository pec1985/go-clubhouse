package models

import (
	"encoding/json"
	"time"
)

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
