package util

import (
	"encoding/json"
	"io"
	"modernday.com/stevejob/streaming/model"
	"net/http"
)

func FlushResponse(res http.ResponseWriter, code int, message string) {
	responseBody, _ := json.Marshal(model.MDResponse{Code: code, Message: message})
	_, _ = io.WriteString(res, string(responseBody))
}