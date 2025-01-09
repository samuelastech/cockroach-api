package common

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   bool   `json:"error"`
	Code    int    `json:"code"`
}

func SendJSON(w http.ResponseWriter, code int, body Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	body.Code = code
	if code >= 400 {
		body.Error = true
	}
	return json.NewEncoder(w).Encode(body)
}
