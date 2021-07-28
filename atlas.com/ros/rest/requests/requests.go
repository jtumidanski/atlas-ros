package requests

import (
	json2 "atlas-ros/json"
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	BaseRequest string = "http://atlas-nginx:80"
)

func Get(url string, resp interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	err = ProcessResponse(r, resp)
	return err
}

func ProcessResponse(r *http.Response, rb interface{}) error {
	err := json2.FromJSON(rb, r.Body)
	if err != nil {
		return err
	}

	return nil
}

func Post(url string, input interface{}) (*http.Response, error) {
	jsonReq, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	r, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	return r, nil
}
