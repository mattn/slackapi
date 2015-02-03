package api

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func get(url string, params url.Values) (*http.Response, error) {
	return http.Get(url + "?" + params.Encode())
}
func getJSON(url string, params url.Values, v interface{}) error {
	resp, err := get(url, params)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
