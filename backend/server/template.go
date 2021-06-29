package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type headerData struct {
	ProcessTime  int64  `json:"process_time_ms"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type response struct {
	Header headerData  `json:"header"`
	Data   interface{} `json:"data"`
}

func RenderResponse(w http.ResponseWriter, statusCode int, data interface{}, startTime time.Time) {
	resp := response{
		Header: headerData{
			ProcessTime: time.Since(startTime).Milliseconds(),
		},
		Data: data,
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(resp)
	w.Write(d)
	return
}

func RenderError(w http.ResponseWriter, statusCode int, err error, startTime time.Time) {
	resp := response{
		Header: headerData{
			ProcessTime:  time.Since(startTime).Milliseconds(),
			ErrorMessage: err.Error(),
		},
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(resp)
	w.Write(d)
	return
}
