package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// ParseJSON parses the request body and saves it into a target
func ParseJSON(reqBody io.ReadCloser, target interface{}) (interface{}, error) {
	body, err := ioutil.ReadAll(reqBody)
	defer reqBody.Close()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, target)
	if err != nil {
		return nil, err
	}
	return target, nil
}
