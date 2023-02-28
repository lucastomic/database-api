package parser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ReqBodyToMap takes a request and parses its body to a map[string]any
// If there is any error it will return it as second parameter
func ReqBodyToMap(req *http.Request) (map[string]any, error) {
	var response map[string]any
	bodyJson, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyJson, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
