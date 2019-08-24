package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response map[string]string

type ErrorInfo struct {
	Message string `json:"message"`
	Slug    string `json:"slug,omitempty"`
}

type ErrorResponse struct {
	ErrorInfo *ErrorInfo `json:"error,omitempty"`
}

type Message struct {
	Status bool       `json:"status"`
	Step   string     `json:"step"`
	Error  *ErrorInfo `json:"error"`
}

func SendResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if body == nil {
		return
	}

	data, err := json.Marshal(body)
	if err != nil {
		//log.ErrorErr(err, "failed marshaling response")

		http.Error(w, "{\"error\": \"Internal Server Error.\"}", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, bytes.NewReader(data))
	if err != nil {
		//log.ErrorErr(err, "failed sending response")
		return
	}
}

func InternalServerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, fmt.Sprintf("{\"error\":\"%s\"}", http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
}

func BadRequestError(w http.ResponseWriter, err error) {
	errorInfo := &ErrorInfo{
		Message: err.Error(),
	}

	SendResponse(w, http.StatusBadRequest, ErrorResponse{errorInfo})
}

func NotFoundError(w http.ResponseWriter, err error) {
	errorInfo := &ErrorInfo{
		Message: err.Error(),
	}

	SendResponse(w, http.StatusNotFound, ErrorResponse{errorInfo})
}
