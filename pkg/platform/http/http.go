package http

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/karchx/nrs/pkg/issue"
)

func query(method, url, token string, jsonBody map[string]interface{}) (map[string]interface{}, error) {
  bodyBuffer := bytes.Buffer{}
  err := json.NewEncoder(&bodyBuffer).Encode(jsonBody)

  req, err :=  http.NewRequest(method, url, &bodyBuffer)
  if err != nil {
    return nil, err
  }

  req.Header.Add("Authorization", "token "+token)
  req.Header.Add("Content-Type", "application/json")

  return issue.QueryHTTP(req)
}
